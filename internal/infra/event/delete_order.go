package event

import "time"

type DeleteOrder struct {
	Name    string
	Payload interface{}
}

func NewDeleteOrder() *DeleteOrder {
	return &DeleteOrder{
		Name: "DeleteOrder",
	}
}

func (e *DeleteOrder) GetName() string {
	return e.Name
}

func (e *DeleteOrder) GetPayload() interface{} {
	return e.Payload
}

func (e *DeleteOrder) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *DeleteOrder) GetDateTime() time.Time {
	return time.Now()
}
