package logminer

import (
	"context"
	"log"
	"time"

	"github.com/ezio1119/fishapp-relaylog/conf"
	"github.com/ezio1119/fishapp-relaylog/domain"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

func StartPostDBLogMining(ctx context.Context, ch chan domain.BinlogEvent, pos mysql.Position) {
	syncer := newPostDBSyncer()
	defer syncer.Close()

	streamer, err := syncer.StartSync(pos)
	if err != nil {
		panic(err)
	}

	for {
		binlogEvent, err := streamer.GetEvent(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		// binlogEvent.Dump(os.Stdout)
		rowsEvent, ok := binlogEvent.Event.(*replication.RowsEvent)

		if ok {
			if string(rowsEvent.Table.Table) == "outbox" {
				nextPos := syncer.GetNextPosition()                                           // index nameを取るため
				lastPos := mysql.Position{Name: nextPos.Name, Pos: binlogEvent.Header.LogPos} //binlogEvent.Header.LogPosにoutboxのrowEventのポジションが入ってる

				ch <- domain.BinlogEvent{
					LastPos:    lastPos,
					PosSubject: conf.C.Nats.Subject.PosPostDB,
					Event:      rowsEvent.Rows[0],
				}
			}
		}
	}
}

func StartChatDBLogMining(ctx context.Context, ch chan domain.BinlogEvent, pos mysql.Position) {
	syncer := newChatDBSyncer()
	defer syncer.Close()

	streamer, err := syncer.StartSync(pos)
	if err != nil {
		panic(err)
	}

	for {
		binlogEvent, err := streamer.GetEvent(ctx)
		if err != nil {
			return
		}
		// binlogEvent.Dump(os.Stdout)
		rowsEvent, ok := binlogEvent.Event.(*replication.RowsEvent)
		if ok {
			if string(rowsEvent.Table.Table) == "outbox" {
				nextPos := syncer.GetNextPosition()
				lastPos := mysql.Position{Name: nextPos.Name, Pos: binlogEvent.Header.LogPos}

				ch <- domain.BinlogEvent{
					LastPos:    lastPos,
					PosSubject: conf.C.Nats.Subject.PosChatDB,
					Event:      rowsEvent.Rows[0],
				}
			}
		}
	}
}

func newPostDBSyncer() *replication.BinlogSyncer {
	return replication.NewBinlogSyncer(replication.BinlogSyncerConfig{
		ServerID:                1,
		Flavor:                  conf.C.PostDB.Dbms,
		Host:                    conf.C.PostDB.Host,
		Port:                    conf.C.PostDB.Port,
		User:                    conf.C.PostDB.User,
		Password:                conf.C.PostDB.Pass,
		Charset:                 conf.C.PostDB.Charset,
		ParseTime:               true,
		TimestampStringLocation: time.Local,
		MaxReconnectAttempts:    10,
	})
}

func newChatDBSyncer() *replication.BinlogSyncer {
	return replication.NewBinlogSyncer(replication.BinlogSyncerConfig{
		ServerID:                1,
		Flavor:                  conf.C.ChatDB.Dbms,
		Host:                    conf.C.ChatDB.Host,
		Port:                    conf.C.ChatDB.Port,
		User:                    conf.C.ChatDB.User,
		Password:                conf.C.ChatDB.Pass,
		Charset:                 conf.C.ChatDB.Charset,
		ParseTime:               true,
		TimestampStringLocation: time.Local,
		MaxReconnectAttempts:    10,
	})
}
