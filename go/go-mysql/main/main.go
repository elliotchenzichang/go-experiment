package main

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"math/rand"
	"os"
	"time"
)

type MyEventHandler struct {
}

func (h *MyEventHandler) OnRotate(header *replication.EventHeader, rotateEvent *replication.RotateEvent) error {
	fmt.Fprintf(os.Stdout, "reccive a OnRotate event, and the header is %+v,the event is %+v\n", header, rotateEvent)
	return nil
}

func (h *MyEventHandler) OnTableChanged(header *replication.EventHeader, schema string, table string) error {
	fmt.Fprintf(os.Stdout, "reccive a OnTableChanged event, and the header is %+v schema is %s, table is %s\n", header, schema, table)
	return nil
}

func (h *MyEventHandler) OnDDL(header *replication.EventHeader, nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	fmt.Fprintf(os.Stdout, "reccive a OnDDL event, and the header is %+v, the next position is %+v, queryEvent is %+v\n", nextPos, queryEvent)
	return nil
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Fprintf(os.Stdout, "reccive a OnRow event, and the content of this event is %+v\n", e)
	return nil
}

func (h *MyEventHandler) OnXID(header *replication.EventHeader, nextPos mysql.Position) error {
	fmt.Fprintf(os.Stdout, "reccive a OnXID event, and the header is %+v, the next position is %+v\n", header, nextPos)
	return nil
}

func (h *MyEventHandler) OnGTID(header *replication.EventHeader, gtid mysql.GTIDSet) error {
	fmt.Fprintf(os.Stdout, "reccive a OnGTID event, and header is %+v, the gtid is %+v\n", header, gtid)
	return nil
}

func (h *MyEventHandler) OnPosSynced(header *replication.EventHeader, pos mysql.Position, set mysql.GTIDSet, force bool) error {
	fmt.Fprintf(os.Stdout, "reccive a OnPosSynced event, and the header is %+v, the position is %+v, gtidSet is %+v\n, force is %+v", header, pos, set, force)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	cfg := new(canal.Config)
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Password = "123456"
	cfg.ServerID = uint32(rand.New(rand.NewSource(time.Now().Unix())).Intn(1000)) + 1001

	c, err := canal.NewCanal(cfg)
	if err != nil {
		fmt.Fprintf(os.Stdout, "encounter a error during init canal, and the error is %s", err.Error())
		return
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
	select {}
}
