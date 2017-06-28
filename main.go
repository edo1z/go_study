package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/gorilla/sessions"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("hoge"))
var sess_name = "sess"

func main() {
	app := iris.New()
	app.Get("/", topPage)
	app.Get("up", upPage)
	app.Get("cl", clearPage)
	app.Run(iris.Addr(":8080"))
}

func topPage (ctx context.Context) {
	f := getFlash(ctx)
	cnt := fmt.Sprintf(
		"<h2>count is %d</h2>",
		getCount(ctx),
	)
	html := `<br><a href="/up">Count Up</a>
	<br><a href="/cl">Count Clear</a>`
	ctx.HTML(f + cnt + html)
}

func upPage (ctx context.Context) {
	countUp(ctx)
	addFlash(ctx, "count up")
	ctx.Redirect("/")
}

func clearPage (ctx context.Context) {
	sessClear(ctx)
	addFlash(ctx,"session clear")
	ctx.Redirect("/")
}