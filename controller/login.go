package controller

import iris "gopkg.in/kataras/iris.v6"
import "fmt"

type LoginUser struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

func LoginPost(ctx *iris.Context) {
	user := LoginUser{}
	if err := ctx.ReadForm(&user); err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(500)
		return
	}
	if user.Name == "admin" && user.Password == "!admin" {
		ctx.Session().Set("logined", true)
		ctx.SetStatusCode(200)
	} else {
		ctx.SetStatusCode(403)
	}
}
