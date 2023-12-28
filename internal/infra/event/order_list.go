package event

import "time"

type OrderList struct {
	Name    string
	Payload interface{}
}

func NewOrderList() *OrderList {
	return &OrderList{
		Name: "OrderList",
	}
}

func (e *OrderList) GetName() string {
	return e.Name
}

func (e *OrderList) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderList) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderList) GetDateTime() time.Time {
	return time.Now()
}
