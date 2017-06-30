package main

import (
	"html/template"
	"bytes"
	"fmt"
)

var viewPath = "../../views/"
var contentType = "text/plain"

type Contact struct {
	Company	string
	Name	string
	Mail	string
	Content	string
}

func main() {
	hoge()
	hoge()
}

func hoge() {
	c := Contact{
		Company: "会社",
		Name: "太郎",
		Mail: "hoge@example.com",
		Content: "どうしたらいいですか？",
	}
	t, err := template.ParseFiles("./mail.tpl", "./welcome.html")
	if err != nil {
		fmt.Println(err)
	}else{
		buff := &bytes.Buffer{}
		t.Execute(buff, c)
		fmt.Println(buff.String())
	}
}

