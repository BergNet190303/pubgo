package publice

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendMsgAliyun(sign string,mobile string,template string,templateParam string,accessKeyId string,accessSecret string) {
	client,err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",accessKeyId,accessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = sign
	request.TemplateCode = template
	request.TemplateParam = templateParam

	response,err := client.SendSms(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("response is",response)
}
