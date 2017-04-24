package schema

//Vps 虚拟机配置
type Vps struct {
	IP         string
	Port       string
	Password   string
	Encryption string
	Veid       string
	Apikey     string
	Private    bool
}
