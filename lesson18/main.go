package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func main() {
	//r := gin.Default() // 默认使用了Logger()和Recovery()中间件
	r := gin.New() // 没有中间件
	r.Use(m1, m2) // 全局注册中间件

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "user",
		})
	})

	group_xx := r.Group("/xx", authMiddleware_(true))
	{
		group_xx.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":"xxGroup",
			})
		})
	}
	group_xx2 := r.Group("/xx2")
	group_xx2.Use(authMiddleware_(true))
	{
		group_xx2.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":"xx2Group",
			})
		})
	}
	r.Run()
}
/**
中间件
 */
//func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in ...")
	name, ok := c.Get("name")
	if !ok {
		name="匿名用户"
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":name,
	})
	fmt.Println("indexHandler out ...")
}
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()

	//go funcxx(c.Copy())  // goroutine中 只能copy出来使用

	c.Next()     // 调用后续的处理函数indexHandler

	//c.Abort()  // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "m2中设置的值")
	//c.Abort()

	fmt.Println("m2 out ...")
	//return
}

func authMiddleware_(doCheck bool) gin.HandlerFunc{
	// 连接数据库 等一些准备工作
	return func(c *gin.Context) {
		// 存放具体的逻辑
		// 是否登录判断
		// if 登录用户
		// c.Next()
		// else
		// c.Abort()
		if doCheck {
			c.Next()
		}else{
			c.Abort()
		}
	}
}