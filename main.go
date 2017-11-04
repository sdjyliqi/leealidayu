package main

import (
	"alidayu"
	"fmt"
)

func main() {
	//第二个参数是配置短信签名名称
	success, resp := alidayu.SendSMS("18701516837", "e宠物医院", "SMS_107845026", `{"msg":"8934"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)

}
