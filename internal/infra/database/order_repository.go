package database

import (
	"database/sql"

	"github.com/NayronFerreira/cleanArq_challenge/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	createTableStmt := `CREATE TABLE IF NOT EXISTS orders (
		id VARCHAR(255) PRIMARY KEY,
		price DECIMAL(10,2),
		tax DECIMAL(10,2),
		final_price DECIMAL(10,2)
	)`
	_, err := r.Db.Exec(createTableStmt)
	if err != nil {
		return err
	}

	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetOrders() ([]*entity.Order, error) {
	rows, err := r.Db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	var orders []*entity.Order
	for rows.Next() {
		order := &entity.Order{}
		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id string) (*entity.Order, error) {
	row := r.Db.QueryRow("SELECT * FROM orders WHERE id = ?", id)
	order := &entity.Order{}
	if err := row.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
		if err == sql.ErrNoRows {
			// Não há linhas correspondentes na tabela
			return nil, nil
		}
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) UpdateOrder(order *entity.Order) (*entity.Order, error) {
	stmt, err := r.Db.Prepare("UPDATE orders SET price = ?, tax = ?, final_price = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(order.Price, order.Tax, order.FinalPrice, order.ID)
	if err != nil {
		return nil, err
	}
	orderUpdate, err := r.GetOrderByID(order.ID)
	if err != nil {
		return nil, err
	}
	return orderUpdate, nil
}

func (r *OrderRepository) DeleteOrder(id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM orders WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
