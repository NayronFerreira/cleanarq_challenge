package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetOrders() ([]*Order, error)
	GetOrderByID(id string) (*Order, error)
	UpdateOrder(order *Order) (*Order, error)
	DeleteOrder(id string) error
}
