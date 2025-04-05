package client

import "fmt"

type Sms struct {
	Text        string //Hi, your OTP Code is \otpCode\
	PhoneNumber string
}

type SMSSender interface {
	SendSMS(s Sms)
}

func (s Sms) SendSMS(sms Sms) {
	fmt.Println(sms.Text)
}
