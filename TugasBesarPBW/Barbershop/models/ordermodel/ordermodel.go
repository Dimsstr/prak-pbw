package ordermodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

// GetAll returns all orders from the database.
func GetAll() []entities.Order {
	rows, err := config.DB.Query(`
		SELECT 
			id,
			name,
			capster_id,
			service_id,
			date,
			time,
			description
		FROM orders
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(
			&order.ID,
			&order.Name,
			&order.CapsterID,
			&order.ServiceID,
			&order.Date,
			&order.Time,
			&order.Description,
		); err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return orders
}

// Create creates a new order in the database.
func Create(order entities.Order) bool {
	result, err := config.DB.Exec(`
		INSERT INTO orders(
			name, capster_id, service_id, date, time, description
		) VALUES (?, ?, ?, ?, ?, ?)`,
		order.Name,
		order.CapsterID,
		order.ServiceID,
		order.Date,
		order.Time,
		order.Description,
	)
	if err != nil {
		panic(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertID > 0
}

// Detail returns the detail of an order based on ID from the database.
func Detail(id int) entities.Order {
	row := config.DB.QueryRow(`
		SELECT 
			id,
			name,
			capster_id,
			service_id,
			date,
			time,
			description
		FROM orders
		WHERE id = ?`, id)

	var order entities.Order
	err := row.Scan(
		&order.ID,
		&order.Name,
		&order.CapsterID,
		&order.ServiceID,
		&order.Date,
		&order.Time,
		&order.Description,
	)
	if err != nil {
		panic(err)
	}
	return order
}

// Update updates an order based on ID in the database.
func Update(id int, order entities.Order) bool {
	query, err := config.DB.Exec(`
		UPDATE orders SET
			name = ?,
			capster_id = ?,
			service_id = ?,
			date = ?,
			time = ?,
			description = ?
		WHERE id = ?`,
		order.Name,
		order.CapsterID,
		order.ServiceID,
		order.Date,
		order.Time,
		order.Description,
		id,
	)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

// Delete deletes an order based on ID from the database.
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM orders WHERE id = ?", id)
	return err
}
