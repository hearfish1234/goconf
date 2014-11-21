package config

import (

	"fmt"
	"testing"
)

func Test(t *testing.T) {

	json, _ := SetConfig("./config.json")
	user :=Json(json).GetValue("mysql").GetValue("user").ToString()
	pass := Json(json).GetValue("mysql").GetValue("pass").ToString()

	fmt.Println(user)
	fmt.Println(pass)

}
