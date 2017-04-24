package bean

import "github.com/go-xorm/xorm"
import "schema"

type VpsBean struct {
	engine *xorm.Engine
}

func (vpsBean *VpsBean) add(vps *schema.Vps) (int64, error) {
	return vpsBean.engine.Insert(vps)
}

func (vpsBean *VpsBean) findAll() []schema.Vps {
	var vpsList []schema.Vps
	vpsBean.engine.Find(&vpsList)
	return vpsList
}
