package event

import "time"

type OrderUpdate struct {
	Name    string
	Payload interface{}
}

func NewOrderUpdate() *OrderUpdate {
	return &OrderUpdate{
		Name: "OrderUpdate",
	}
}

func (e *OrderUpdate) GetName() string {
	return e.Name
}

func (e *OrderUpdate) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderUpdate) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderUpdate) GetDateTime() time.Time {
	return time.Now()
}
