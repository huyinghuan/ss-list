package bean

import (
	"ss-list/schema"
	"ss-list/utils"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func GetDBConenct() (*xorm.Engine, error) {
	config := utils.GetConfig()
	engine, connectErr := xorm.NewEngine(config.Db.Driver, config.Db.Connect)
	if connectErr != nil {
		return nil, connectErr
	}
	err := engine.Sync2(new(schema.Vps))
	if err != nil {
		return nil, err
	}
	return engine, nil
}
