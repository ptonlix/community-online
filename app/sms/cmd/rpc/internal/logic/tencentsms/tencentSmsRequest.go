package tencentsms

import (
	"github/community-online/app/sms/cmd/rpc/internal/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type SmsRequest struct {
	request *sms.SendSmsRequest
}

func NewSmsRequest(options *config.SmsConfig, withOptions ...func(smsRequest *SmsRequest)) *SmsRequest {
	request := sms.NewSendSmsRequest()

	request.SmsSdkAppId = &options.SmsSdkAppId
	request.SignName = &options.SignName
	smsRequest := &SmsRequest{request: request}
	for _, option := range withOptions {
		option(smsRequest)
	}
	return smsRequest

}

type RequestOption func(*SmsRequest)

func WithTemplateIdSet(id string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.TemplateId = common.StringPtr(id)
	}
}

func WithPhoneNumberSet(phoneSet []string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.PhoneNumberSet = common.StringPtrs(phoneSet)
	}
}

func WithTemplateParamSet(templateSet []string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.TemplateParamSet = common.StringPtrs(templateSet)
	}
}

func WitSmsSdkAppId(smsSdkAppId string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.SmsSdkAppId = common.StringPtr(smsSdkAppId)
	}
}

func WitSignName(signName string) RequestOption {
	return func(smsRequest *SmsRequest) {
		smsRequest.request.SignName = common.StringPtr(signName)
	}
}
