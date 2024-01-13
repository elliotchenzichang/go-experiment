package es

import (
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

type DemoData1 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DemoData2 struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Score   int    `json:"score"`
}

func getESClient() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth("user", "secret"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
}

func TestGetClient(t *testing.T) {
	client, err := getESClient()
	assert.Nil(t, err)
	assert.NotNil(t, client)
}
