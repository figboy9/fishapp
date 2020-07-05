package infrastructure

import (
	"context"
	"log"

	"github.com/ezio1119/fishapp-post/conf"
	"github.com/ezio1119/fishapp-post/interfaces/controllers"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewNatsStreamingConn() (stan.Conn, error) {
	clientID := uuid.New().String()
	log.Printf("nats clientID is %s", clientID)

	return stan.Connect(conf.C.Nats.ClusterID, "fishapp-post-"+clientID, stan.NatsURL(conf.C.Nats.URL))
}

func StartSubscribeCreatePostSagaReply(conn stan.Conn, c controllers.SagaReplyController) error {
	_, err := conn.QueueSubscribe("create.post.saga.reply", conf.C.Nats.QueueGroup, func(m *stan.Msg) {
		ctx := context.Background()
		e := &pb.Event{}
		if err := protojson.Unmarshal(m.MsgProto.Data, e); err != nil {
			log.Printf("error failed unmarshal protojson: %s", err)
			return
		}

		log.Printf("recieved event: %#v\n", e)

		switch e.EventType {
		case "room.created":
			data := &pb.RoomCreated{}
			if err := protojson.Unmarshal(e.EventData, data); err != nil {
				log.Printf("error failed unmarshal protojson: %s", err)
				return
			}
			if err := c.RoomCreated(ctx, data); err != nil {
				log.Println(err)
				return
			}
		case "create.room.failed":
			data := &pb.CreateRoomFailed{}
			if err := protojson.Unmarshal(e.EventData, data); err != nil {
				log.Printf("error failed unmarshal protojson: %s", err)
				return
			}
			if err := c.CreateRoomFailed(ctx, data); err != nil {
				log.Println(err)
				return
			}
		}

	}, stan.DurableName(conf.C.Nats.QueueGroup))
	if err != nil {
		return err
	}
	return nil
}
