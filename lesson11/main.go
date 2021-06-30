package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取请求中携带的query string参数
		//name := c.Query("query")  // 通过query获取请求中携带的query参数
		//name := c.DefaultQuery("query", "somebody") //取不到就用默认值
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	// 取不到
		//	name = "somebody"
		//}
//http://localhost/web?query=%E7%A8%8B%E6%BD%87&age=18
		name := c.Query("query")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name":name,
			"age":age,
		})
	})
	r.Run(":80")
}