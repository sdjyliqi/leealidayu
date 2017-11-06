//文件的功能：读取yaml 配置文件，然后初始化内存中。
package alidayu

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type PJSetting struct {
	AliMSM struct {
		AppKey     string `yaml:"Appkey"`
		TemplateId string `yaml:"TemplateId"`
		AppSecret  string `yaml:"AppSecret"`
		Https      bool   `yaml:"Https"`
	}
}

//备注，在进行文件读取的时候，目的都是从src，也就是gopath下的src目录开始的。
//完成时间：2017年10月
//备注1：冒号后面一定要有空格
//备注2：缩进要用空格
var Settings PJSetting

func init() {

	fileContent, _ := ioutil.ReadFile("./alidayu/conf.yaml")
	yaml.Unmarshal([]byte(fileContent), &Settings)
	fmt.Println("key:", Settings.AliMSM.AppKey)
	fmt.Println("sec:", Settings.AliMSM.AppSecret)

}
