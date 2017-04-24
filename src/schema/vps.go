package schema

//Vps 虚拟机配置
type Vps struct {
	Id         int64  `xorm:"unique"`
	IP         string `xorm:"ip"`
	Port       string
	Password   string
	Encryption string
	Veid       string
	Apikey     string
	Private    bool
}
