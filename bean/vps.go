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
	if err := engine.SQL("SELECT * FROM vps WHERE private = ?", "false").Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}

func (vpsBean *VpsBean) Update(id int64, update map[string]interface{}) error {
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return err
	}
	if _, err = engine.Table(new(schema.Vps)).Id(id).Update(update); err != nil {
		return err
	}
	return nil
}

func (vpsBean *VpsBean) Get(id int64) (*schema.Vps, error) {
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return nil, err
	}
	vps := new(schema.Vps)
	if _, err = engine.Id(id).Get(vps); err != nil {
		return nil, err
	}
	return vps, nil
}
