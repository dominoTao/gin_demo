package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username:=c.Query("username")
		//password:=c.Query("password")
		//u := UserInfo{
		//	username:username,
		//	password: password,
		//}
	//http://localhost/user?username=q1&password=123
		var u UserInfo   // 声明UserInfo类型的变量
		err := c.ShouldBind(&u)   // 值类型，必须用取址符
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message":"ok",
			})
		}
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo   // 声明UserInfo类型的变量
		err := c.ShouldBind(&u)   // 值类型，必须用取址符
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message":"ok",
			})
		}
	})
	r.Run(":80")
}

