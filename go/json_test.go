package _go

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

type Data struct {
	Id int64 `json:"id"`
}

var int64Data = []Data{
	{
		Id: 1 << 53,
	},
	{
		Id: 1<<53 + 1,
	},
	{
		Id: 1<<53 + 2,
	},
	{
		Id: 1<<53 + 3,
	},
}

func TestHttpJson(t *testing.T) {
	var httpJsonTestHandler = func(w http.ResponseWriter, r *http.Request) {
		log.Println("com here")

		r.Header.Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("content-type", "text/json")
		if resByte, jsonErr := json.Marshal(int64Data); jsonErr != nil {
			w.Write([]byte(jsonErr.Error()))
			log.Fatal(jsonErr)
		} else {
			w.Write(resByte)
		}
		return
	}
	http.HandleFunc("/json_test", httpJsonTestHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {}
}

func TestJsonNumberWithGin(t *testing.T) {
	r := gin.Default()
	r.GET("json_test", func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, int64Data)
	})
	r.Run(":9000")
}
