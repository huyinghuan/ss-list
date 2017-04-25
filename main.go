package main

import (
	"ss-list/controller"

	iris "gopkg.in/kataras/iris.v6"
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
		view.HTML("./views", ".html").Reload(true),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))
	app.StaticWeb("/js", "./static/js")
	app.StaticWeb("/css", "./static/css")
	app.Party("/admin", controller.AuthCtrl).
		Get("/manager", controller.ManagerCtrl)

	login := app.Party("/login").Layout("layouts/layout.html")
	{
		login.Get("", controller.LoginView)
		login.Post("", controller.LoginPost)
	}

	app.Listen(":6300")
}
