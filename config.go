package conf4Fish

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"strconv"
)

//设置日志类型
var	filePath string

type Config struct {
	data interface{}
}

//设置配置文件路径
func SetConfig (filepath string) *Config {

	filePath = filepath

	c := readList()
	return c
}

//读取配置信息列表
func readList () *Config {

	fp, err := os.Open(filePath)
	if err != nil {

		checkErr(err)
	}
	defer fp.Close()

	buff, err := ioutil.ReadAll(fp)
	if err != nil {

		checkErr(err)
	}

	var data interface{}
	err = json.Unmarshal(buff, &data)
	if err != nil {

		checkErr(err)
	}

	c := new(Config)
	c.data = data

	return c
}

//支持解析json格式字符串
func ReadJson(str string) *Config {

	var data interface{}
	buff := []byte(str)
	err := json.Unmarshal(buff, &data)
	if err != nil {

		checkErr(err)
	}

	c := new(Config)
	c.data = data

	return c
}

//读取配置信息
func (c *Config) GetValue(section string) *Config{

	readList()

	list := c.data.(map[string]interface{})
	if config, ok := list[section]; ok{

		c.data = config
		return c
	}
	c.data = nil
	return c
}

//转换格式
func (c *Config) ToString() string {

	if m, ok := c.data.(string); ok {

		return m
	}

	if m, ok := c.data.(float64); ok {

		return strconv.FormatFloat(m, 'f', -1, 64)
	}

	return ""
}

//错误处理
func checkErr (err error) string {

	if err != nil {

		return fmt.Sprintf("You got a error: %s", err)
	}

	return ""
}
