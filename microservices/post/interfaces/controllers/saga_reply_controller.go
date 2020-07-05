package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-post/pb"
	"github.com/ezio1119/fishapp-post/usecase/interactor"
)

type sagaReplyController struct {
	sagaReplyInteractor interactor.SagaReplyInteractor
}

type SagaReplyController interface {
	RoomCreated(ctx context.Context, e *pb.RoomCreated) error
	CreateRoomFailed(ctx context.Context, e *pb.CreateRoomFailed) error
}

func NewSagaReplyController(i interactor.SagaReplyInteractor) SagaReplyController {
	return &sagaReplyController{i}
}

func (c *sagaReplyController) RoomCreated(ctx context.Context, e *pb.RoomCreated) error {
	return c.sagaReplyInteractor.RoomCreated(ctx, e.SagaId)
}

func (c *sagaReplyController) CreateRoomFailed(ctx context.Context, e *pb.CreateRoomFailed) error {
	return c.sagaReplyInteractor.CreateRoomFailed(ctx, e.SagaId, e.Message)
	// for _, detail := range e.ErrorStatus.Details {
	// 	switch t := detail.(type) {
	// 	case *errdetails.BadRequest:
	// 		fmt.Println("Oops! Your request was rejected by the server.")
	// 		for _, violation := range t.GetFieldViolations() {
	// 			fmt.Printf("The %q field was wrong:\n", violation.GetField())
	// 			fmt.Printf("\t%s\n", violation.GetDescription())
	// 		}
	// 	}
	// }
}
