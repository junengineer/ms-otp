package service

import (
	"errors"
	"ms-otp/client"
	"ms-otp/repo"
	"ms-otp/util"
)

type OTPService struct {
	SmsSender      client.SMSSender
	SessionStorage repo.SessionStorage
}

func (o OTPService) SendOtp(phoneNumber string) (string, error) {
	otp, sessionId := util.BuildOTP(), util.BuildSessionId()
	otpSession := repo.OTPSession{
		OtpCode: otp,
		Count:   1,
	}

	err := o.SessionStorage.Save(sessionId, &otpSession)
	if err != nil {
		return "", err
	}

	sms := client.Sms{
		Text:        "Your OTP code is " + otp,
		PhoneNumber: phoneNumber,
	}
	o.SmsSender.SendSMS(sms)
	return sessionId, nil
}

func (o OTPService) CheckOtp(sessionId, otp string) error {
	otpSession, err := o.SessionStorage.GetByID(sessionId)

	if err != nil {
		return err
	}

	if otpSession.Count > 3 {
		err = o.SessionStorage.DeleteByID(sessionId)
		if err == nil {
			err = errors.New("please, request new otp code")
		}
	} else if otpSession.Count <= 3 && otpSession.OtpCode == otp {
		err = nil
	} else {
		otpSession.Count++
		err = o.SessionStorage.Save(sessionId, &otpSession)
	}
	return err
}
