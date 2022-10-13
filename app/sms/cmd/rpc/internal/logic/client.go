package logic

type SmsClient interface {
	Send() bool
}
