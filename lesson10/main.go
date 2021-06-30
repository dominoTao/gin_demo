package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// method 1：使用map
		c.JSON(http.StatusOK, gin.H{
			"name" : "小王子",
			"message" : "hello world !",
			"age" : 18,
		})
		// method 2: 结构体
		type msg struct{
			Name string
			Message string
			Age int
		}
	})
	type msg struct{  // 灵活使用tag做定制化json
		Name string `json:"name"`
		Message string `bson:"xxx"`
		Age int `json:"age"`
	}
	r.GET("/type_json", func(c *gin.Context) {
		data := msg{
			"name",
			"hello golang",
			22,
		}
		c.JSON(http.StatusOK, data) // json的序列化  结构体中小写开头的属性不会被序列化
	})

	r.Run(":80")
}