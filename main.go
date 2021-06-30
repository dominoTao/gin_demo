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

func demo1(w http.ResponseWriter, r *http.Request)  {

	// 定义模板
	// 解析模板
	files, err := template.ParseFiles("./template/t.tmpl", "./template/ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "小王子"
	files.Execute(w, name)

}

func home(w http.ResponseWriter, r *http.Request)  {
	files, err := template.ParseFiles("./template/home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err : %v\n", err)
		return
	}
	name := "小王子"
	files.Execute(w, name)
}
//func index(w http.ResponseWriter, r *http.Request)  {
//	files, err := template.ParseFiles("./template/index.tmpl")
//	if err != nil {
//		fmt.Printf("parse template failed , err : %v\n", err)
//		return
//	}
//	name := "小王子"
//	files.Execute(w, name)
//
//}
func home2(w http.ResponseWriter, r *http.Request)  {
	files, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err : %v\n", err)
		return
	}
	name := "小王子"
	//files.Execute(w, name)
	files.ExecuteTemplate(w, "home.tmpl", name)
}
func index2(w http.ResponseWriter, r *http.Request)  {
	files, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err : %v\n", err)
		return
	}
	name := "小王子"
	files.ExecuteTemplate(w, "index.tmpl", name)
}
func index(w http.ResponseWriter, r *http.Request)  {
	files, err := template.New("index1.tmpl").
		Delims("{[", "]}").
		ParseFiles("./template/index1.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err : %v\n", err)
		return
	}
	name := "小王子"
	//files.ExecuteTemplate(w, "index.tmpl", name)
	err = files.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed , err : %v\n", err)
		return
	}
}
func xss(w http.ResponseWriter, r *http.Request) {
	files, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe" : func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./template/xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed , err : %v\n", err)
		return
	}
	str1 := `<script>alert(123);</script>`
	str2 := "<a href='http://liwenzhou.com'>李文周</a>"
	err = files.Execute(w, map[string]string{
		"str1":str1,
		"str2":str2,
	})
	if err != nil {
		fmt.Printf("render template failed , err : %v\n", err)
		return
	}
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	//http.HandleFunc("/index", index)
	//http.HandleFunc("/home", home)
	//http.HandleFunc("/index2", index2)
	//http.HandleFunc("/home2", home2)
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err%v\n", err)
		return
	}
}
