package controller

import iris "gopkg.in/kataras/iris.v6"

func LoginCtrl(ctx *iris.Context) {
	ctx.SetStatusCode(500)
	ctx.Done()
}
