package event_handler

import (
	"encoding/json"
	"fmt"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/pkg/mq"
)

type CategoryEventHandler struct {
}

func (categoryEventHandle *CategoryEventHandler) Handle(params interface{}) {
	orderId := params.(string)
	rabbitMQ := mq.NewRabbitMQ()
	orderMap := map[string]string{"order_id": orderId}
	messageJson, _ := json.Marshal(orderMap)
	if err := rabbitMQ.SendMessage(string(messageJson),
		consts.GoodsCategoryChange, consts.GoodsCategoryChange+"_ROUTING"); err != nil {
		fmt.Println("发送消息失败：", err.Error())
	}
}
