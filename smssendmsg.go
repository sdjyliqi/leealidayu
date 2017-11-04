package alidayu

import (
	"time"
)

func SendSMS(recNum, smsFreeSign, smsTemplateId, smsPaloadPars string) (success bool, response string) {
	if recNum == "" || smsFreeSign == "" || smsTemplateId == "" {
		return false, "Parameter not complete"
	}
	params := make(map[string]string)
	params["app_key"] = Settings.Alidayu.AppKey
	params["format"] = "json"
	params["method"] = "alibaba.aliqin.fc.sms.num.send"
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = smsFreeSign
	params["rec_num"] = recNum
	params["sms_template_code"] = smsTemplateId
	params["sms_param"] = smsPaloadPars

	return DoHttpPost(params)

}
