package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	// 路由组
	// 把公用的前缀提取出来 创建一个路由组
	group := r.Group("/view")
	{
		group.GET("/a", func(c *gin.Context) {

		})
		group.GET("/b", func(c *gin.Context) {

		})
		group.GET("/c", func(c *gin.Context) {

		})
	}

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"DELETE",
		})
	})
	// any 请求方法的大集合
	r.Any("/ii", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method":"GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method":"POST",
			})
			// ...
		}
	})
	r.HEAD("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"HEAD",
		})
	})
	r.LoadHTMLFiles("./views/404.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "./views/404.html", nil)
	})
	r.Run(":8080")
}
