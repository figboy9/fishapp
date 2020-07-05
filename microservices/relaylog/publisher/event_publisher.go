package publisher

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ezio1119/fishapp-relaylog/domain"
	"github.com/ezio1119/fishapp-relaylog/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/encoding/protojson"
)

func StartEventPublishing(conn stan.Conn, ch chan domain.BinlogEvent) {
	for bEvent := range ch {
		e, err := convEventProto(bEvent.Event)
		if err != nil {
			log.Println(err)
			continue
		}

		eventByte, err := protojson.Marshal(e)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := conn.Publish(e.Channel, eventByte); err != nil {
			log.Println(err)
			continue
		}
		log.Printf("success: published event on %s subject: %s", e.Channel, e.EventType)

		posByte, err := json.Marshal(bEvent.LastPos)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := conn.Publish(bEvent.PosSubject, posByte); err != nil {
			log.Println(err)
			continue
		}

		log.Printf("success: published position: %+v", bEvent.LastPos)
	}
}

func convEventProto(raw []interface{}) (*pb.Event, error) {
	e := &pb.Event{}
	var ok bool
	var err error

	e.Id, ok = raw[0].(string)
	if !ok {
		return nil, errors.New("failed conv event.id")
	}

	e.EventType, ok = raw[1].(string)
	if !ok {
		return nil, errors.New("failed conv event.event_type")
	}

	e.EventData, ok = raw[2].([]byte)
	if !ok {
		return nil, errors.New("failed conv event.event_data")
	}

	e.AggregateId, ok = raw[3].(string)
	if !ok {
		return nil, errors.New("failed conv event.aggregate_id")
	}

	e.AggregateType, ok = raw[4].(string)
	if !ok {
		return nil, errors.New("failed conv event.aggregate_type")
	}

	e.Channel, ok = raw[5].(string)
	if !ok {
		return nil, errors.New("failed conv event.channel")
	}
	fmt.Printf("created_at: %T\n", raw[6])

	cAt, ok := raw[6].(time.Time)
	if !ok {
		return nil, errors.New("failed conv event.created_at")
	}

	e.CreatedAt, err = ptypes.TimestampProto(cAt)
	if err != nil {
		return nil, err
	}

	uAt, ok := raw[7].(time.Time)
	if !ok {
		return nil, errors.New("failed conv event.updated_at")
	}

	e.UpdatedAt, err = ptypes.TimestampProto(uAt)
	if err != nil {
		return nil, err
	}

	return e, nil
}
