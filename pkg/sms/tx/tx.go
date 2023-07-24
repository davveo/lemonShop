package tx

type TxSms struct {
}

func (tx *TxSms) OnSend(phone, content string, m map[string]interface{}) bool {
	return false
}

func NewTxSms() *TxSms {
	return &TxSms{}
}
