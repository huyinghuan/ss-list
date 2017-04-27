package schema

//Vps 虚拟机配置
type Vps struct {
	ID            int64  `xorm:"unique" json:"id"`
	IP            string `xorm:"ip" json:"ip" form:"ip"`
	Port          string `json:"port" form:"port"`
	Password      string `json:"password" form:"password"`
	Encryption    string `json:"encryption" form:"encryption"`
	Veid          string `json:"veid" form:"veid"`
	Apikey        string `json:"apikey" form:"apikey"`
	Private       bool   `json:"private" form:"private"`
	Network       string `json:"network"`
	DataNextReset string `json:"data_next_reset"`
}
