package interactor

import (
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

func newApplyPostCreatedEvent(a *models.ApplyPost) (*models.Outbox, error) {
	now := time.Now()
	aP, err := convApplyPostProto(a)
	if err != nil {
		return nil, err
	}

	applyPostCreated, err := protojson.Marshal(&pb.ApplyPostCreated{ApplyPost: aP})
	if err != nil {
		return nil, err
	}

	event := &models.Outbox{
		ID:            uuid.New().String(),
		EventType:     "apply.post.created",
		EventData:     applyPostCreated,
		AggregateID:   strconv.FormatInt(a.ID, 10),
		AggregateType: "apply.post",
		Channel:       "apply.post.created",
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return event, nil
}

func newApplyPostDeletedEvent(a *models.ApplyPost) (*models.Outbox, error) {
	now := time.Now()
	aP, err := convApplyPostProto(a)
	if err != nil {
		return nil, err
	}

	applyPostDeleted, err := protojson.Marshal(&pb.ApplyPostDeleted{ApplyPost: aP})
	if err != nil {
		return nil, err
	}

	event := &models.Outbox{
		ID:            uuid.New().String(),
		EventType:     "apply.post.deleted",
		EventData:     applyPostDeleted,
		AggregateID:   strconv.FormatInt(a.ID, 10),
		AggregateType: "apply.post",
		Channel:       "apply.post.deleted",
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return event, nil
}

func newPostDeletedEvent(p *models.Post) (*models.Outbox, error) {
	now := time.Now()
	pPost, err := convPostProto(p)
	if err != nil {
		return nil, err
	}

	eventData, err := protojson.Marshal(&pb.PostDeleted{Post: pPost})
	if err != nil {
		return nil, err
	}

	event := &models.Outbox{
		ID:            uuid.New().String(),
		EventType:     "post.deleted",
		EventData:     eventData,
		AggregateID:   strconv.FormatInt(pPost.Id, 10),
		AggregateType: "post",
		Channel:       "post.deleted",
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	return event, nil
}
