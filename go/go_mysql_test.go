package _go

import (
	"context"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/siddontang/go/log"
	"os"
	"testing"
	"time"
)

func TestMysqlReplication(t *testing.T) {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "12345678",
	}

	syncer := replication.NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(mysql.Position{
		Name: "",
		Pos:  0,
	})

	go func() {
		for {
			ev, _ := streamer.GetEvent(context.Background())
			// Dump event
			ev.Dump(os.Stdout)
		}
	}()

	// or we can use a timeout context
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		ev, err := streamer.GetEvent(ctx)
		cancel()

		if err == context.DeadlineExceeded {
			// meet timeout
			continue
		}

		ev.Dump(os.Stdout)
	}
}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Infof("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func TestMysqlCanal(t *testing.T) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Password = "12345678"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "elliot-test"
	cfg.Dump.Tables = []string{"test-table"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}
