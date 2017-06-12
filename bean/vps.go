package bean

import (
	"ss-list/schema"
)

type VpsBean struct{}

func (vpsBean *VpsBean) Add(vps *schema.Vps) (int64, error) {
	engine := GetDBConnect()
	return engine.Insert(vps)
}

func (vpsBean *VpsBean) FindAll() ([]schema.Vps, error) {
	var vpsList []schema.Vps
	engine := GetDBConnect()
	if err := engine.Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}

func (vpsBean *VpsBean) FindPublic() ([]schema.Vps, error) {
	var vpsList []schema.Vps
	engine := GetDBConnect()
	if err := engine.SQL("SELECT * FROM vps WHERE private = ?", "false").Find(&vpsList); err != nil {
		return nil, err
	}
	return vpsList, nil
}

func (vpsBean *VpsBean) Update(id int64, update map[string]interface{}) error {
	engine := GetDBConnect()
	if _, err := engine.Table(new(schema.Vps)).Id(id).Update(update); err != nil {
		return err
	}
	return nil
}

func (vpsBean *VpsBean) Get(id int64) (*schema.Vps, error) {
	engine := GetDBConnect()
	vps := new(schema.Vps)
	if _, err := engine.Id(id).Get(vps); err != nil {
		return nil, err
	}
	return vps, nil
}
