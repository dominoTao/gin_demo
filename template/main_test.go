package main__test

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
)
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err%v\n", err)
		return
	}
	//value := "world !"
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	hobbyList := []string{
		"抽烟","喝酒","烫头",
	}
	// 渲染模板
	err = files.Execute(w, map[string]interface{}{
		"user" : user,
		"hobby" : hobbyList,
	})
	if err != nil {
		fmt.Printf("render template failed, err%v\n", err)
		return
	}
}

func TestView(t *testing.T) {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err%v\n", err)
		return
	}
}
