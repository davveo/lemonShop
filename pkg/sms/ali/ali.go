package ali

type AliSms struct {
}

func (ali *AliSms) OnSend(phone, content string, m map[string]interface{}) bool {
	return false
}

func NewAliSms() *AliSms {
	return &AliSms{}
}
