package controller

import iris "gopkg.in/kataras/iris.v6"
import "fmt"

func LoginView(ctx *iris.Context) {
	//ctx.ViewLayout("layout.html")
	//iris.RenderOptions{"layout": "layout.html"}
	// iris.RenderOptions{"gzip": true}
	ctx.Render("login.html", iris.Map{})
}

func LoginPost(ctx *iris.Context) {
	form := ctx.FormValues()
	for key, value := range form {
		fmt.Println(key, value)
	}
	ctx.JSON(200, map[string]string{"hello": "world"})
}
