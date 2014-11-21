package config

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"strconv"
)

type Js struct {
	data interface{}
}

//设置配置文件
func SetConfig (filePath string) (string, error) {

	fp, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	json, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

//解析json
func Json(jsonStr string) *Js {

	j := new(Js)
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return j
	}

	j.data = data
	return j
}

//获取json数据
func (j *Js) getJson() map[string]interface{} {

	if m, ok := j.data.(map[string]interface{}); ok {
	
		return m
	}

	return nil
}

//获取配置值
func (j *Js) GetValue(key string) *Js {

	m := j.getJson()
	if v, ok := m[key]; ok {
	
		j.data = v
		return j
	}

	j.data = nil
	return j
}

//转换格式
func (j *Js) ToString() string {

	if m, ok := j.data.(string); ok {
	
		return m
	}

	if m, ok := j.data.(float64); ok {
	
		return strconv.FormatFloat(m, 'f', -1, 64)
	}

	return ""
}
