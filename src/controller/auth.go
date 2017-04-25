package controller

import iris "gopkg.in/kataras/iris.v6"
import "fmt"

func AuthCtrl(ctx *iris.Context) {
	fmt.Println(1)
	if auth, _ := ctx.Session().GetBoolean("logined"); !auth {
		fmt.Println(2)
		ctx.RedirectTo("/login")
	} else {
		ctx.Next()
	}
}
