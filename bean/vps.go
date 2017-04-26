package bean

import (
	"ss-list/schema"

	"github.com/go-xorm/xorm"
)

type VpsBean struct{}

func (vpsBean *VpsBean) Add(vps *schema.Vps) (int64, error) {
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return 0, err
	}
	return engine.Insert(vps)
}

func (vpsBean *VpsBean) FindAll() ([]schema.Vps, error) {
	var vpsList []schema.Vps
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return nil, err
	}
	if err := engine.Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}

func (vpsBean *VpsBean) FindPublic() ([]schema.Vps, error) {
	var vpsList []schema.Vps
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return nil, err
	}
	if err := engine.Cols("ip", "Password", "Port", "Encryption").Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}
