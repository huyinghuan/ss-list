package controller

import (
	"strings"

	"gopkg.in/kataras/iris.v6"
)

func AuthCtrl(ctx *iris.Context) {
	if strings.Index(ctx.Path(), "login") != -1 {
		ctx.Next()
		return
	}
	if auth, _ := ctx.Session().GetBoolean("logined"); !auth {
		ctx.SetStatusCode(401)
		ctx.Done()
	} else {
		ctx.Next()
	}
}
