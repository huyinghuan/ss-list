package controller

import (
	"fmt"
	"ss-list/bean"
	"ss-list/utils"

	"ss-list/schema"

	"log"

	"encoding/json"

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
		for index, item := range list {
			if item.Veid == "" || item.Apikey == "" {
				continue
			}
			vpsInfo := utils.GetVPSNetwork(item.Veid, item.Apikey)
			list[index].Network = vpsInfo.Remaining
			list[index].DataNextReset = vpsInfo.DataReset
			list[index].IP = vpsInfo.IPAddresses[0]
		}
		ctx.JSON(200, list)
	}
}
func (vpsCtrl *VpsCtrl) Post(ctx *iris.Context) {
	vpsBean := &bean.VpsBean{}
	vpsModel := schema.Vps{}
	if err := ctx.ReadForm(&vpsModel); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(403)
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

func (vpsCtrl *VpsCtrl) Put(ctx *iris.Context) {
	id, paramsErr := ctx.ParamInt64("id")
	if paramsErr != nil {
		log.Fatal(paramsErr)
		ctx.SetStatusCode(403)
		ctx.Writef("数据不存在")
		return
	}

	vpsBean := &bean.VpsBean{}
	vpsModel := schema.Vps{}
	if err := ctx.ReadForm(&vpsModel); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(403)
		ctx.Writef("上传表单错误")
		return
	}
	fmt.Println(vpsModel)
	if r, e := json.Marshal(vpsModel); e == nil {
		fmt.Println(string(r))
	}
	if err := vpsBean.Update(id, &vpsModel); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(500)
		ctx.Writef("修改数据失败")
		return
	}
	ctx.SetStatusCode(200)
	ctx.Done()
}

func (vpsCtrl *VpsCtrl) GetPublic(ctx *iris.Context) {
	vpsBean := &bean.VpsBean{}
	var vpsList []schema.Vps
	if list, err := vpsBean.FindPublic(); err != nil {
		log.Fatal(err)
		ctx.SetStatusCode(500)
		ctx.Writef("查询数据失败")
		return
	} else {
		vpsList = list
	}

	for index, item := range vpsList {
		if item.Veid == "" || item.Apikey == "" {
			continue
		}
		vpsInfo := utils.GetVPSNetwork(item.Veid, item.Apikey)
		vpsList[index].Network = vpsInfo.Remaining
		vpsList[index].DataNextReset = vpsInfo.DataReset
		vpsList[index].IP = vpsInfo.IPAddresses[0]
	}

	ctx.JSON(200, vpsList)
}
