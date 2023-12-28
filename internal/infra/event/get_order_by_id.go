package event

import "time"

type GetOrderByID struct {
	Name    string
	Payload interface{}
}

func NewGetOrderByID() *GetOrderByID {
	return &GetOrderByID{
		Name: "GetOrderByID",
	}
}

func (e *GetOrderByID) GetName() string {
	return e.Name
}

func (e *GetOrderByID) GetPayload() interface{} {
	return e.Payload
}

func (e *GetOrderByID) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *GetOrderByID) GetDateTime() time.Time {
	return time.Now()
}
