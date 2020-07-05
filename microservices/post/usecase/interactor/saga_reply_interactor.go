package interactor

import (
	"context"
	"log"

	"github.com/ezio1119/fishapp-post/pb"
	"github.com/ezio1119/fishapp-post/usecase/interactor/saga"
	"github.com/ezio1119/fishapp-post/usecase/repo"
	"google.golang.org/protobuf/encoding/protojson"
)

type sagaReplyInteractor struct {
	createPostSagaManager *saga.CreatePostSagaManager
	sagaInstanceRepo      repo.SagaInstanceRepo
}

func NewSagaReplyInteractor(m *saga.CreatePostSagaManager, sr repo.SagaInstanceRepo) SagaReplyInteractor {
	return &sagaReplyInteractor{m, sr}
}

type SagaReplyInteractor interface {
	RoomCreated(ctx context.Context, sagaID string) error
	CreateRoomFailed(ctx context.Context, sagaID string, errMsg string) error
}

func (i *sagaReplyInteractor) RoomCreated(ctx context.Context, sagaID string) error {
	sagaIn, err := i.sagaInstanceRepo.GetSagaInstance(ctx, sagaID)
	if err != nil {
		return err
	}
	p := &pb.Post{}
	if err := protojson.Unmarshal(sagaIn.SagaData, p); err != nil {
		return err
	}

	state, err := saga.NewCreatePostSagaState(p, sagaIn.CurrentState, sagaID)
	if err != nil {
		return err
	}

	s := i.createPostSagaManager.NewCreatePostSagaManager(state)

	if err := s.FSM.Event("ApprovePost", ctx); err != nil {
		if err := s.FSM.Event("RejectPost", ctx); err != nil {
			return err
		}
		return err
	}
	return nil

}

func (i *sagaReplyInteractor) CreateRoomFailed(ctx context.Context, sagaID string, errMsg string) error {
	log.Printf("error: %s\n", errMsg)

	sagaIn, err := i.sagaInstanceRepo.GetSagaInstance(ctx, sagaID)
	if err != nil {
		return err
	}

	p := &pb.Post{}
	if err := protojson.Unmarshal(sagaIn.SagaData, p); err != nil {
		return err
	}

	state, err := saga.NewCreatePostSagaState(p, sagaIn.CurrentState, sagaID)
	if err != nil {
		return err
	}

	s := i.createPostSagaManager.NewCreatePostSagaManager(state)
	if err := s.FSM.Event("RejectPost", ctx, errMsg); err != nil {
		return err
	}

	return nil
}
