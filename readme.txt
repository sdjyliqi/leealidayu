使用方式：

第一步：首先构建yaml格式的阿里大鱼配置文件，具体内容如下：
alimsm:
  Appkey: "24681239"
  TemplateId: "SMS_108020127"
  AppSecret: "a01ae1342cc8122da4da72c0a679dab0"
  Https: false

第二部 调用阿里大鱼的短信模块

1.利用阿里大鱼的api初始化，制定配置文件位置。
2:利用定义的api发送短信。

附加示例：
package main

import (
	"fmt"
	"github.com/sdjyliqi/leealidayu"
)

func main() {
	//其中初始化函数中的参数为阿里大鱼的配置文件
	alidayu.SettingInit("./pjconf/alidayu.yaml")
	success, resp := alidayu.SendSMS("15210510285", "e宠物医院", "SMS_107845026", `{"msg":"8934"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
}



