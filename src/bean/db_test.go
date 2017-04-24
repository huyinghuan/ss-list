package bean

import "testing"
import "fmt"

func TestAll(t *testing.T) {
	engine, err := GetDBConenct()
	if err != nil {
		t.Fail()
	}
	results, queryErr := engine.Query("select * from vps")
	if queryErr != nil {
		t.Fail()
	}
	for index, value := range results {
		fmt.Println(index, value)
	}
}
