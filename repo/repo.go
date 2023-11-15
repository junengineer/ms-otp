package repo

import "errors"

type OTPSession struct {
	OtpCode string
	Count   int
}

var storageMap = make(map[string]OTPSession)

type SessionStorage interface {
	GetByID(id string) (OTPSession, error)
	DeleteByID(id string) error
	Save(sessionId string, o *OTPSession) error
}

func (o OTPSession) GetByID(id string) (OTPSession, error) {
	otpSession, ok := storageMap[id]
	if !ok {
		return otpSession, errors.New("invalid session")
	}
	return otpSession, nil
}

func (o OTPSession) DeleteByID(id string) error {
	_, err := o.GetByID(id)
	if err != nil {
		return err
	}
	delete(storageMap, id)
	return nil
}

func (o OTPSession) Save(sessionId string, otp *OTPSession) error {
	storageMap[sessionId] = *otp
	return nil
}
