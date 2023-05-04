package event_handler

import (
	"encoding/json"
	"fmt"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/pkg/mq"
)

type OrderEventHandler struct {
}

func (orderEventHandle *OrderEventHandler) Handle(params interface{}) {
	orderId := params.(string)
	rabbitMQ := mq.NewRabbitMQ()
	orderMap := map[string]string{"order_id": orderId}
	messageJson, _ := json.Marshal(orderMap)
	if err := rabbitMQ.SendMessage(string(messageJson), consts.ORDER_CREATE, consts.ORDER_CREATE+"_QUEUE"); err != nil {
		fmt.Println("发送消息失败：", err.Error())
	}
}
