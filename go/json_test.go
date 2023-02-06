package _go

import (
	"encoding/json"
	"fmt"
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
		Id: 1,
	},
	{
		Id: 2,
	},
	{
		Id: 3,
	},
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

func TestJsonMarshal(t *testing.T) {
	type Student struct {
		// 这里的首字母必须是大写，不然无法导出，也没办法marshal出来
		//  这里的 json tag 可以自定义输出json字段的名字
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var stu = &Student{
		Age:  30,
		Name: "Elliot",
	}

	bytes, err := json.Marshal(stu)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))
}
