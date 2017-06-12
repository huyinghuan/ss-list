package main

import (
	"ss-list/bean"
	"ss-list/controller"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
)

func main() {
	bean.InitConnect()
	app := iris.New()
	sessionAdapt := sessions.New(sessions.Config{})

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		sessionAdapt)

	admin := app.Party("/admin", controller.AuthCtrl)

	vpsCtrl := &controller.VpsCtrl{}
	admin.Get("/vps", vpsCtrl.GetAll)
	admin.Post("/vps", vpsCtrl.Post)
	admin.Put("/vps/:id", vpsCtrl.Put)
	admin.Get("/vps/:id", vpsCtrl.Get)

	admin.Post("/login", controller.LoginPost)

	public := app.Party("/public", func(ctx *iris.Context) { ctx.Next() })
	public.Get("/vps", vpsCtrl.GetPublic)

	app.Listen(":6300")
}
