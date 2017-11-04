//文件的功能：读取yaml 配置文件，然后初始化内存中。
package pjsettings

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type PJSetting struct {
	Alidayu struct {
		AppKey     string `yaml:"AppKey"`
		TemplateId string `yaml:"TemplateId"`
		AppSecret  string `yaml:"AppSecret"`
		UseHttps   bool   `yaml:"Https"`
	}
}

//备注，在进行文件读取的时候，目的都是从src，也就是gopath下的src目录开始的。
//完成时间：2017年10月
//备注1：冒号后面一定要有空格
//备注2：缩进要用空格
var Settings PJSetting

func init() {
	fileContent, _ := ioutil.ReadFile("./projectconf/conf.yaml")
	yaml.Unmarshal([]byte(fileContent), &Settings)
	fmt.Println("key:",Settings.Alidayu.AppKey)
	fmt.Println("sec:",Settings.Alidayu.AppSecret)

}
