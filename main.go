package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数
	// 要么只有一个返回值；要么有两个返回值 第二个必须是error
	// 要在解析模板文件之前  定义好函数
	kua := func(name string) (string, error){
		return name + "联盟", nil
	}
	t := template.New("f.tmpl")  // 名字一定要个解析的名字对应上  或者解析的文件组中必须包含新建的文件名
	t.Funcs(template.FuncMap{
		"lakua":kua,
	})
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err%v\n", err)
		return
	}
	name := "英雄"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err%v\n", err)
		return
	}
}
