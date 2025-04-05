package main

import (
	"fmt"
	"ms-otp/client"
	"ms-otp/repo"
	"ms-otp/service"
)

func main() {
	temporary()
}
func temporary() {
	//otpService := service.OTPService{
	//	SmsSender: client.Sms{
	//		"test", "testNumber",
	//	},
	//	SessionStorage: repo.OTPSession{OtpCode: "123", Count: 1},
	//}
	//otpService.SendOtp("+994512297773")
	sms := client.Sms{}
	storage := repo.OTPSession{}

	otpService := service.OTPService{
		SmsSender: sms, SessionStorage: storage,
	}
	var s, err = otpService.SendOtp("+994512297773")
	if err == nil {
		println(s)
	}
	err = otpService.CheckOtp(s, "122")
	//if err != nil {
	//	println(err.Error())
	//}
	err = otpService.CheckOtp(s, "122")
	err = otpService.CheckOtp(s, "122")
	err = otpService.CheckOtp(s, "122")
	if err != nil {
		println(err.Error())
	}
	id, err := otpService.SessionStorage.GetByID(s)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("ID IS", id)
	s, err = otpService.SendOtp("+994512297773")

	err = otpService.CheckOtp(s, "555")
}
