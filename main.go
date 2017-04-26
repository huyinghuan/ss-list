package main

import (
	"ss-list/controller"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
)

func main() {
	app := iris.New()
	sessionAdapt := sessions.New(sessions.Config{})

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		sessionAdapt,
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	admin := app.Party("/admin", controller.AuthCtrl)
	admin.Get("/vps", controller.ManagerCtrl)

	admin.Post("/login", controller.LoginPost)

	app.Listen(":6300")
}
