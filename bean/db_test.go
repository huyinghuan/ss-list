package bean

import (
	"fmt"
	"ss-list/schema"
	"testing"
)

func TestAll(t *testing.T) {
	engine, err := GetDBConenct()
	if err != nil {
		t.Fail()
	}
	vpsBean := &VpsBean{engine}
	vps := &schema.Vps{}
	vps.IP = "12.23"
	if f, e := vpsBean.add(vps); e != nil {
		fmt.Println(e)
		t.Fail()
	} else {
		fmt.Println(f)
	}
	list, _ := vpsBean.findAll()
	for _, item := range list {
		fmt.Println(item)
	}
}
