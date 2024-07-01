package capstermodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Capster {
	rows, err := config.DB.Query(`SELECT * FROM capsters`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var capsters []entities.Capster

	for rows.Next() {
		var capster entities.Capster
		if err := rows.Scan(&capster.ID, &capster.Name, &capster.CreatedAt, &capster.UpdatedAt); err != nil {
			panic(err)
		}

		capsters = append(capsters, capster)
	}

	return capsters
}

func Create(capster entities.Capster) bool {
	result, err := config.DB.Exec(`
		INSERT INTO capsters (name, created_at, updated_at) 
		VALUES (?, ?, ?)`,
		capster.Name,
		capster.CreatedAt,
		capster.UpdatedAt,
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

func Detail(id int) entities.Capster {
	row := config.DB.QueryRow(`SELECT id, name FROM capsters WHERE id = ?`, id)

	var capster entities.Capster

	if err := row.Scan(&capster.ID, &capster.Name); err != nil {
		panic(err.Error())
	}

	return capster
}

func Update(id int, capster entities.Capster) bool {
	query, err := config.DB.Exec(`UPDATE capsters SET name = ?, updated_at = ? WHERE id = ?`, capster.Name, capster.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM capsters WHERE id = ?", id)
	return err
}
