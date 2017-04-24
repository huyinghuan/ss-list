package utils

import "testing"
import "log"
import "fmt"
import "path/filepath"

func TestReadConfig(T *testing.T) {
	config := Config{}
	path, _ := filepath.Abs("../conf/config.yaml")
	err := ReadConfig(path, &config)
	if err != nil {
		log.Fatalln(err)
		T.Fail()
	} else {
		fmt.Println(config.Db.Driver)
		fmt.Println(config.Db.Connect)
	}
}

func TestGetConfig(T *testing.T) {
	config := GetConfig()
	fmt.Println(config.Db.Driver)
	fmt.Println(config.Db.Connect)
}
