package sms

import (
	"github.com/davveo/lemonShop/pkg/sms/ali"
	"github.com/davveo/lemonShop/pkg/sms/tx"
)

type SmsPlatform interface {
	OnSend(phone, content string, m map[string]interface{}) bool
}

func NewSms(smsType string) SmsPlatform {
	switch smsType {
	case "ALI":
		return ali.NewAliSms()
	case "TX":
		return tx.NewTxSms()
	default:
		return tx.NewTxSms()
	}
}
