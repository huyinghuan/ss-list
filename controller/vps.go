package controller

import (
	"ss-list/bean"

	"ss-list/schema"

	"log"

	"fmt"

	iris "gopkg.in/kataras/iris.v6"
)

type VpsCtrl struct {
}

func (vpsCtrl *VpsCtrl) GetAll(ctx *iris.Context) {
	vpsBean := &bean.VpsBean{}
	if list, err := vpsBean.FindAll(); err != nil {
		ctx.SetStatusCode(500)
		ctx.Writef("查询数据失败")
		return
	} else {
		ctx.JSON(200, list)
	}
}
func (vpsCtrl *VpsCtrl) Post(ctx *iris.Context) {
	vpsBean := &bean.VpsBean{}
	fmt.Println(ctx.FormValue("ip"))
	fmt.Println(ctx.FormValues())
	vpsModel := schema.Vps{}
	if err := ctx.ReadForm(&vpsModel); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(401)
		ctx.Writef("上传表单错误")
		return
	}
	if _, err := vpsBean.Add(&vpsModel); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(500)
		ctx.Writef("插入数据错误")
		return
	}
	ctx.SetStatusCode(200)
	ctx.Done()
}

func (vpsCtrl *VpsCtrl) GetPublic(ctx *iris.Context) {

}
