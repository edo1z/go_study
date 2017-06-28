package main

import (
	"github.com/kataras/iris/context"
	"fmt"
)

func getCount(ctx context.Context) int {
	sess, _ := store.Get(ctx.Request(), sess_name)
	cnt := 0
	if v, ok := sess.Values["cnt"]; ok {
		if c, ok := v.(int); ok {
			cnt = c
		}
	}
	return cnt
}

func countUp(ctx context.Context) {
	sess, _ := store.Get(ctx.Request(), sess_name)
	cnt := 1
	if v, ok := sess.Values["cnt"]; ok {
		if c, ok := v.(int); ok {
			cnt = c + 1
		}
	}
	sess.Values["cnt"] = cnt
	sess.Save(ctx.Request(), ctx.ResponseWriter())
}

func sessClear(ctx context.Context) {
	sess, _ := store.Get(ctx.Request(), sess_name)
	sess.Values = make(map[interface{}]interface{})
	sess.Save(ctx.Request(), ctx.ResponseWriter())
}

func getFlash(ctx context.Context) string {
	sess, _ := store.Get(ctx.Request(), sess_name)
	f := sess.Flashes()
	sess.Save(ctx.Request(), ctx.ResponseWriter())
	if len(f) > 0 {
		return fmt.Sprintf(
			`<p style="color:#090;padding:5px;">%s</p>`,
			f[0],
		)
	}
	return ""
}

func addFlash(ctx context.Context, msg string) {
	sess, _ := store.Get(ctx.Request(), sess_name)
	sess.AddFlash(msg)
	sess.Save(ctx.Request(), ctx.ResponseWriter())
}