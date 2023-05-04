package event

// EventManager 事件管理器
type EventManager struct {
	events map[string]EventHandler
}

// Bind 绑定事件
func (eventManager *EventManager) Bind(eventName string, eventHandle EventHandler) *EventManager {
	isBind := false
	for name, _ := range eventManager.events {
		if eventName == name {
			isBind = true
		}
	}
	if !isBind {
		if eventManager.events == nil {
			eventManager.events = make(map[string]EventHandler)
		}
		// todo 增加并发安全控制
		eventManager.events[eventName] = eventHandle
	}
	return eventManager
}

// Trigger 触发事件
func (eventManager *EventManager) Trigger(eventName string, params interface{}) {
	for name, event := range eventManager.events {
		if eventName == name {
			event.Handle(params)
		}
	}
}
