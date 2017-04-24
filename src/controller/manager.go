package controller

import (
	"fmt"

	iris "gopkg.in/kataras/iris.v6"
)

func ManagerCtrl(ctx *iris.Context) {
	err := ctx.Render("m.hbs", iris.Map{"Title": "Page Title"}, iris.RenderOptions{"gzip": true})
	if err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(500)
		ctx.Done()
	}
}
