package main

import (
	"controller"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	app := iris.New()

	sessionAdapt := sessions.New(sessions.Config{})

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		sessionAdapt,
		view.Handlebars("./views", ".hbs").Reload(true),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	app.StaticWeb("/js", "./static/js")
	app.StaticWeb("/css", "./static/css")

	app.Party("/admin", controller.AuthCtrl).
		Get("/manager", controller.ManagerCtrl)

	app.Listen(":6300")
}
