package controller

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

func AuthCtrl(ctx *iris.Context) {
	if auth, _ := ctx.Session().GetBoolean("logined"); !auth {
		fmt.Println(2)
		ctx.Redirect("/login", 302)
		ctx.Done()
	} else {
		ctx.Next()
	}
}
