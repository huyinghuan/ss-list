package bean

import "github.com/go-xorm/xorm"
import "schema"

type VpsBean struct {
	engine *xorm.Engine
}

func (vpsBean *VpsBean) add(vps *schema.Vps) (int64, error) {
	return vpsBean.engine.Insert(vps)
}

func (vpsBean *VpsBean) findAll() ([]schema.Vps, error) {
	var vpsList []schema.Vps
	if err := vpsBean.engine.Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}
