package servicemodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Service {
	rows, err := config.DB.Query(`SELECT id, name, price, created_at, updated_at FROM services`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var services []entities.Service

	for rows.Next() {
		var service entities.Service
		if err := rows.Scan(&service.ID, &service.Name, &service.Price, &service.CreatedAt, &service.UpdatedAt); err != nil {
			panic(err)
		}

		services = append(services, service)
	}

	return services
}

func Create(service entities.Service) bool {
	result, err := config.DB.Exec(`
		INSERT INTO services (name, price, created_at, updated_at) 
		VALUES (?, ?, ?, ?)`,
		service.Name,
		service.Price,
		service.CreatedAt,
		service.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Service {
	row := config.DB.QueryRow(`SELECT id, name, created_at, updated_at FROM services WHERE id = ?`, id)

	var service entities.Service

	if err := row.Scan(&service.ID, &service.Name, &service.CreatedAt, &service.UpdatedAt); err != nil {
		panic(err)
	}

	return service
}

// Update updates a service based on ID in the database.
func Update(id int, service entities.Service) bool {
	query, err := config.DB.Exec(`UPDATE services SET name = ?, price = ?, updated_at = ? WHERE id = ?`, service.Name, service.Price, service.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM services WHERE id = ?", id)
	return err
}
