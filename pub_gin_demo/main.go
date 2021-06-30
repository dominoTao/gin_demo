package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)
// 静态文件：
// html页面上用到的样式文件 .css  js  图片


func resume(){
	r := gin.Default()
	r.Static("/xxx", "./resume")// 加载静态文件
	//r.LoadHTMLGlob("resume/*")// 解析模板
	r.LoadHTMLFiles("resume/index.html")
	r.GET("/resume", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":80")
}
func main() {
	demo()
	//resume()
}

func demo() {
	r := gin.Default()
	// 加载静态文件  把./statics文件夹映射到/xxx  用以在静态文件中引用资源使用
	r.Static("/xxx", "./statics")
	// gin 框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 解析模板
	r.LoadHTMLFiles("statics/index.html")
	//r.LoadHTMLGlob("statics/*")

	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	//		"title":"posts/index.tmpl",
	//	})
	//})
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	//		"title":"<a href='http://liwenzhou.com'>李文周</a>",
	//	})
	//})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":80") // 启动server
}