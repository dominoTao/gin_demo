package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// router.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// read request
		//file, err := c.FormFile("f1") // 从请求中获取携带的参数
		form, err := c.MultipartForm()
		files := form.File["file"]
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		}else{
			// read file to local
			//dst := fmt.Sprintf("./%s", file.Filename)
			////dst2 := path.Join("./", file.Filename)
			//c.SaveUploadedFile(file, dst)
			//c.JSON(http.StatusOK, gin.H{
			//	"status" : "ok",
			//})

			// batch


			for index,file:=range files {
				log.Println(file.Filename)
				dst:=fmt.Sprintf("D:/tmp/%d_%s", index, file.Filename)
				// 上传到指定目录
				c.SaveUploadedFile(file, dst)
			}
			c.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		}
	})
	r.Run(":80")
}
