package config

type SmsConfig struct {
	Secret      map[string]interface{} `json:"Secret"`
	SmsSdkAppId string                 `json:"SmsSdkAppId"`
	SignName    string                 `json:"SignName"`
	TemplateIds map[string]interface{} `json:"TemplateIds"`
}
