package saga

import (
	"time"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

func newCreateRoomEvent(c *pb.CreateRoom) (*models.Outbox, error) {
	eventData, err := protojson.Marshal(c)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	return &models.Outbox{
		ID:        uuid.New().String(),
		EventType: "create.room",
		EventData: eventData,
		Channel:   "create.room",
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func newPostApprovedEvent(p *pb.Post, sagaID string) (*models.Outbox, error) {
	postApproved := &pb.PostApproved{
		SagaId: sagaID,
		Post:   p,
	}

	jsonEvent, err := protojson.Marshal(postApproved)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &models.Outbox{
		ID:        uuid.New().String(),
		EventType: "post.approved",
		EventData: jsonEvent,
		Channel:   "create.post.result",
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func newPostRejectedEvent(p *pb.Post, sagaID string, errMsg string) (*models.Outbox, error) {
	postRejected := &pb.PostRejected{
		SagaId:       sagaID,
		Post:         p,
		ErrorMessage: errMsg,
	}

	jsonEvent, err := protojson.Marshal(postRejected)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &models.Outbox{
		ID:        uuid.New().String(),
		EventType: "post.rejected",
		EventData: jsonEvent,
		Channel:   "create.post.result",
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
