package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-relaylog/conf"
	"github.com/ezio1119/fishapp-relaylog/domain"
	"github.com/ezio1119/fishapp-relaylog/logminer"
	"github.com/ezio1119/fishapp-relaylog/publisher"
	"github.com/nats-io/stan.go"
	"github.com/siddontang/go-mysql/mysql"
)

func main() {
	ctx := context.Background()
	conn, err := newNatsConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	postDBPos := mysql.Position{}
	chatDBPos := mysql.Position{}

	postDBPosCh := make(chan mysql.Position)
	chatDBPosCh := make(chan mysql.Position)
	defer close(postDBPosCh)
	defer close(chatDBPosCh)

	go getLastPosPostDB(conn, postDBPosCh)
	go getLastPosChatDB(conn, chatDBPosCh)

Loop:
	for {
		select {
		case postDBPos = <-postDBPosCh:
		case chatDBPos = <-chatDBPosCh:
		case <-time.After(time.Second * 2):
			break Loop
		}
	}

	eventChan := make(chan domain.BinlogEvent)

	for i := 0; i < conf.C.Nats.PublisherNum; i++ {
		go publisher.StartEventPublishing(conn, eventChan)
	}

	go logminer.StartPostDBLogMining(ctx, eventChan, postDBPos)
	go logminer.StartChatDBLogMining(ctx, eventChan, chatDBPos)

	http.HandleFunc(conf.C.Sv.HealthCheck.Path, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "relaylog is healthy!\n")
	})

	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.HealthCheck.Port, nil))
}

func newNatsConn() (stan.Conn, error) {
	return stan.Connect(conf.C.Nats.ClusterID, conf.C.Nats.ClientID, stan.NatsURL(conf.C.Nats.URL))
}

func getLastPosPostDB(conn stan.Conn, ch chan mysql.Position) {

	if _, err := conn.Subscribe(conf.C.Nats.Subject.PosPostDB, func(msg *stan.Msg) {

		pos := mysql.Position{}
		if err := json.Unmarshal(msg.MsgProto.Data, &pos); err != nil {
			panic(err)
		}

		if err := msg.Sub.Close(); err != nil {
			panic(err)
		}

		ch <- pos
	}, stan.StartWithLastReceived()); err != nil {
		panic(err)
	}
}

func getLastPosChatDB(conn stan.Conn, ch chan mysql.Position) {

	if _, err := conn.Subscribe(conf.C.Nats.Subject.PosChatDB, func(msg *stan.Msg) {

		pos := mysql.Position{}
		if err := json.Unmarshal(msg.MsgProto.Data, &pos); err != nil {
			panic(err)
		}

		if err := msg.Sub.Close(); err != nil {
			panic(err)
		}

		ch <- pos
	}, stan.StartWithLastReceived()); err != nil {
		panic(err)
	}
}
