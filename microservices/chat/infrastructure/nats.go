package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

type EventController interface {
	CreateRoom(m *stan.Msg)
	PostDeleted(m *stan.Msg)
	ApplyPostCreated(m *stan.Msg)
	ApplyPostDeleted(m *stan.Msg)
}

func NewNatsStreamingConn() (stan.Conn, error) {
	clientID := uuid.New().String()
	log.Printf("nats clientID is %s", clientID)
	conn, err := stan.Connect(conf.C.Nats.ClusterID, "fishapp-chat-"+uuid.New().String(), stan.NatsURL(conf.C.Nats.URL))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func StartSubscribeNats(c EventController, conn stan.Conn) error {
	_, err := conn.QueueSubscribe("create.room", conf.C.Nats.QueueGroup, c.CreateRoom, stan.DurableName(conf.C.Nats.QueueGroup))
	_, err = conn.QueueSubscribe("post.deleted", conf.C.Nats.QueueGroup, c.PostDeleted, stan.DurableName(conf.C.Nats.QueueGroup), stan.SetManualAckMode())
	_, err = conn.QueueSubscribe("apply.post.deleted", conf.C.Nats.QueueGroup, c.ApplyPostDeleted, stan.DurableName(conf.C.Nats.QueueGroup), stan.SetManualAckMode())
	_, err = conn.QueueSubscribe("apply.post.created", conf.C.Nats.QueueGroup, c.ApplyPostCreated, stan.DurableName(conf.C.Nats.QueueGroup), stan.SetManualAckMode())

	if err != nil {
		return err
	}

	return nil
}
