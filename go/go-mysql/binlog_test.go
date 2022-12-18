package go_mysql

import (
	"github.com/go-mysql-org/go-mysql/replication"
	"testing"
)

func TestBinlog(t *testing.T) {
	parser := &replication.BinlogParser{}
	var binlog []byte

	var timestamp = []byte{0x00, 0x00, 0x00, 0x00}
	var eventType = byte(0x1e) // WRITE_ROWS_EVENTv2
	var serverID = []byte{0x00, 0x00, 0x00, 0x01}
	var logPos = []byte{0x00, 0x00, 0x00, 0x01}
	var flags = []byte{0x00, 0x00}

	binlog = append(binlog, timestamp...)
	binlog = append(binlog, eventType)
	binlog = append(binlog, serverID...)
	binlog = append(binlog, logPos...)
	binlog = append(binlog, flags...)

	e, err := parser.Parse(binlog)
	if err != nil {
		t.Error(err)
	}
	t.Logf("e: %+v", e)
}
