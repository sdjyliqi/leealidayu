package alidayu

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

//函数功能：向alidayu发送http请求
//参数说明：m为http请求参数，map字典格式
//返回值说明：如果发送成功，返回true+阿里大鱼的返回信息
//           否则，返回false+ 异常信息提示
func DoHttpPost(m map[string]string) (success bool, response string) {
	if Settings.AliMSM.AppKey == "" || Settings.AliMSM.AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}

	body, size := GetRequestBody(m)
	client := &http.Client{}
	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", "http://gw.api.taobao.com/router/rest", body)
	if err != nil {
		return false, err.Error()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	resp, err := client.Do(req)
	if err != nil {
		response = err.Error()
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response = string(data)
	if strings.Contains(response, "success") {
		return true, response
	}
	return false, response
}

//函数功能：针对请求参数做合法授权
//参数列表：m为客户端请求的原始请求参数，格式为map。
//         根据原始参数，进行拼接，然后在请求参数中增加签名的参数sign
//返回值：加工后的请求参数字符串和参数长度
func GetRequestBody(m map[string]string) (reader io.Reader, size int64) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}

	signString := Settings.AliMSM.AppSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += Settings.AliMSM.AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)
	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}
