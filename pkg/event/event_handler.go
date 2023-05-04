package event

type EventHandler interface {
	Handle(params interface{})
}
