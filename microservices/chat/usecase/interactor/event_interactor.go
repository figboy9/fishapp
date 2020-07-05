package interactor

import (
	"context"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb"
	"github.com/ezio1119/fishapp-chat/usecase/repo"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/encoding/protojson"
)

type eventInteractor struct {
	db         *gorm.DB
	imageRepo  repo.ImageRepo
	ctxTimeout time.Duration
}

func NewEventInteractor(db *gorm.DB, i repo.ImageRepo, t time.Duration) *eventInteractor {
	return &eventInteractor{db, i, t}
}

type EventInteractor interface {
	CreateRoom(ctx context.Context, r *domain.Room, sagaID string) error
	PostDeleted(ctx context.Context, pID int64) error
	ApplyPostCreated(ctx context.Context, postID int64, userID int64) error
	ApplyPostDeleted(ctx context.Context, postID int64, userID int64) error
}

func (i *eventInteractor) CreateRoom(ctx context.Context, r *domain.Room, sagaID string) error {
	err := i.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(r).Error; err != nil {
			return err
		}

		rProto, err := convRoomProto(r)
		if err != nil {
			return err
		}

		eventDataJSON, err := protojson.Marshal(&pb.RoomCreated{
			SagaId: sagaID,
			Room:   rProto,
		})
		if err != nil {
			return err
		}

		event := &domain.Outbox{
			ID:            uuid.New().String(),
			EventType:     "room.created",
			EventData:     eventDataJSON,
			AggregateID:   strconv.FormatInt(r.ID, 10),
			AggregateType: "room",
			Channel:       "create.post.saga.reply",
		}

		if err := tx.Create(event).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {

		eventDataJSON, err := protojson.Marshal(&pb.CreateRoomFailed{
			SagaId:  sagaID,
			Message: err.Error(),
		})
		if err != nil {
			return err
		}

		event := &domain.Outbox{
			ID:        uuid.New().String(),
			EventType: "create.room.failed",
			EventData: eventDataJSON,
			Channel:   "create.post.saga.reply",
		}

		if err := i.db.Create(event).Error; err != nil {
			return err
		}
	}

	return nil
}

func (i *eventInteractor) PostDeleted(ctx context.Context, pID int64) error {
	r := &domain.Room{}
	msgs := []*domain.Message{}

	if err := i.db.Where("post_id = ?", pID).Find(r).Error; err != nil {
		return err
	}

	if err := i.db.Where("room_id = ?", r.ID).Find(msgs).Error; err != nil {
		return err
	}

	msgIDs := make([]int64, len(msgs))
	for i, m := range msgs {
		msgIDs[i] = m.ID
	}

	if err := i.imageRepo.BatchDeleteImagesByMessageIDs(ctx, msgIDs); err != nil {
		return err
	}

	return i.db.Where("post_id = ?", pID).Delete(domain.Room{}).Error
}

func (i *eventInteractor) ApplyPostCreated(ctx context.Context, pID int64, uID int64) error {
	r := &domain.Room{}
	if err := i.db.Where("post_id = ?", pID).Find(r).Error; err != nil {
		return err
	}

	return i.db.Create(&domain.Member{RoomID: r.ID, UserID: uID}).Error
}
func (i *eventInteractor) ApplyPostDeleted(ctx context.Context, pID int64, uID int64) error {
	r := &domain.Room{}
	if err := i.db.Where("post_id = ?", pID).Find(r).Error; err != nil {
		return err
	}

	return i.db.Where("room_id = ? AND user_id = ?", r.ID, uID).Delete(&domain.Member{}).Error
}
