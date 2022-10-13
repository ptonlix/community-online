package tencentsms

import (
	"github/community-online/app/sms/cmd/rpc/internal/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsClient struct {
	Credential *common.Credential
	Region     string
	Cpf        *profile.ClientProfile
	Request    SmsRequest
}

type Option func(*SmsClient)

func NewSmsClient(options ...func(client *SmsClient)) *SmsClient {
	client := &SmsClient{
		Region: "ap-guangzhou",
		Cpf:    profile.NewClientProfile(),
	}
	for _, option := range options {
		option(client)
	}
	return client

}

func WithRequest(request SmsRequest) Option {
	return func(smsClient *SmsClient) {
		smsClient.Request = request
	}
}

func WithCredential(options config.SmsConfig) Option {
	return func(smsClient *SmsClient) {
		smsClient.Credential = common.NewCredential(options.Secret["SecretId"].(string), options.Secret["SecretKey"].(string))
	}
}
func WithCpfReqMethod(method string) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.HttpProfile.ReqMethod = method
	}
}
func WithCpfReqTimeout(timeout int) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.HttpProfile.ReqTimeout = timeout
	}
}
func WithCpfSignMethod(method string) Option {
	return func(smsClient *SmsClient) {
		smsClient.Cpf.SignMethod = method
	}
}

func (s *SmsClient) Send() bool {
	sendClient, _ := sms.NewClient(s.Credential, s.Region, s.Cpf)
	_, err := sendClient.SendSms(s.Request.request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logx.Errorf("An API error has returned: %s", err)
		return false
	}

	if err != nil {
		logx.Errorf("发送短信失败:%s", err)
		return false

	}
	return true
}
