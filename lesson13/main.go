package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/:year/:month", func(c *gin.Context) {
		// 获取路径参数
		year := c.Param("year")  // string
		month := c.Param("month")  // string
		c.JSON(http.StatusOK, gin.H{
			"year":year,
			"month":month,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		// 获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":year,
			"month":month,
		})
	})
	r.Run(":80")
}