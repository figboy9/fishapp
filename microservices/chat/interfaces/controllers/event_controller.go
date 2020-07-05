package controllers

import (
	"context"
	"log"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/encoding/protojson"
)

type eventController struct {
	eventInteractor interactor.EventInteractor
}

func NewEventController(ei interactor.EventInteractor) *eventController {
	return &eventController{ei}
}

func (c *eventController) CreateRoom(m *stan.Msg) {
	ctx := context.Background()

	e := &pb.Event{}
	if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
		log.Printf("error wrong subject data type : %s", err)
		return
	}
	log.Printf("recieved create.room event: %#v\n", e)

	eventData := &pb.CreateRoom{}
	if err := protojson.Unmarshal(e.EventData, eventData); err != nil {
		log.Printf("error wrong eventdata type: %s", err)
		return
	}

	if err := c.eventInteractor.CreateRoom(ctx, &domain.Room{
		PostID: eventData.PostId,
		Members: []*domain.Member{
			{UserID: eventData.UserId},
		},
	}, eventData.SagaId); err != nil {
		log.Println(err)
		return
	}

}

func (c *eventController) PostDeleted(m *stan.Msg) {
	ctx := context.Background()

	e := &pb.Event{}
	if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
		log.Printf("error wrong subject data type : %s", err)
		return
	}
	log.Printf("recieved post.deleted event: %#v\n", e)

	eventData := &pb.PostDeleted{}
	if err := protojson.Unmarshal(e.EventData, eventData); err != nil {
		log.Printf("error wrong eventdata type: %s", err)
		return
	}

	if err := c.eventInteractor.PostDeleted(ctx, eventData.Post.Id); err != nil {
		log.Println(err)
		return
	}

	if err := m.Ack(); err != nil {
		log.Println(err)
	}
}

func (c *eventController) ApplyPostCreated(m *stan.Msg) {
	ctx := context.Background()

	e := &pb.Event{}
	if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
		log.Printf("error wrong subject data type : %s", err)
		return
	}
	log.Printf("recieved apply.post.created event: %#v\n", e)

	data := &pb.ApplyPostCreated{}
	if err := protojson.Unmarshal(e.EventData, data); err != nil {
		log.Printf("error wrong eventdata type: %s", err)
		return
	}

	if err := c.eventInteractor.ApplyPostCreated(ctx, data.ApplyPost.Id, data.ApplyPost.UserId); err != nil {
		log.Println(err)
		return
	}

	if err := m.Ack(); err != nil {
		log.Println(err)
	}
}

func (c *eventController) ApplyPostDeleted(m *stan.Msg) {
	ctx := context.Background()

	e := &pb.Event{}
	if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
		log.Printf("error wrong subject data type : %s", err)
		return
	}
	log.Printf("recieved apply.post.deleted event: %#v\n", e)

	data := &pb.ApplyPostDeleted{}
	if err := protojson.Unmarshal(e.EventData, data); err != nil {
		log.Printf("error wrong eventdata type: %s", err)
		return
	}

	if err := c.eventInteractor.ApplyPostDeleted(ctx, data.ApplyPost.Id, data.ApplyPost.UserId); err != nil {
		log.Println(err)
		return
	}
	if err := m.Ack(); err != nil {
		log.Println(err)
	}
}
