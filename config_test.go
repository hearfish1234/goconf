package conf4Fish

import (

	"fmt"
	"testing"
)

func Test(t *testing.T) {

	conf := SetConfig("./config.json")
	user :=conf.GetValue("mysql").GetValue("user").ToString()
	json := `{"user":"name", "testBox": {"test1": "123", "test2": 456}}`
	buff := []byte(json)
	jsonConf := ReadJson(buff)

	test := jsonConf.GetValue("testBox").GetValue("test1").ToString()


	fmt.Println(user)
	fmt.Println(test)
}
