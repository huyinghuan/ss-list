package main

import (
	"controller"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		view.Handlebars("./views", ".hbs").Reload(true),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	app.StaticWeb("/js", "./static/js")
	app.StaticWeb("/css", "./static/css")
	app.Get("/login", controller.LoginCtrl)
	admin := app.Party("/admin", controller.authCtrl)
	admin.Get("/admin/manager", controller.ManagerCtrl)

	app.Listen(":6300")
}
