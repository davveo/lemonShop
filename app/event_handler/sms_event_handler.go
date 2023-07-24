package event_handler

import (
	"encoding/json"
	"fmt"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/models/vo"
	"github.com/davveo/lemonShop/pkg/mq"
)

type SmsEventHandler struct {
}

func (smsEventHandle *SmsEventHandler) Handle(params interface{}) {
	rabbitMQ := mq.NewRabbitMQ()
	smsSendVo := params.(vo.SmsSendVo)
	messageJson, _ := json.Marshal(smsSendVo)
	if err := rabbitMQ.SendMessage(string(messageJson),
		consts.SEND_MESSAGE, consts.SEND_MESSAGE+"_QUEUE"); err != nil {
		fmt.Println("发送消息失败：", err.Error())
	}
}
