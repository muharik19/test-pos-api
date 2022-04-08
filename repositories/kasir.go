package repositories

import (
	"fmt"
	"log"

	"github.com/devcode-pos/models"

	dbconnect "github.com/devcode-pos/databases"
)

func lastID(id int) (response *models.ResponseKasir) {
	var name, passcode, createdAt string
	var updatedAt interface{}
	query := fmt.Sprintf(`SELECT id, passcode, name, createdAt, updatedAt FROM master_kasir WHERE id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&passcode,
		&name,
		&createdAt,
		&updatedAt,
	)
	if passcode != "" {
		d := &models.ResponseKasir{
			Success: true,
			Message: "Success",
			Data: &models.Kasir{
				Passcode:  passcode,
				ID:        id,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseKasir{
		Success: false,
		Message: "Cashier Not Found",
		Data:    &models.Kasir{},
	}
	response = d
	return
}

func CreatedKasir(request *models.RequestKasir) (response *models.ResponseKasir) {
	var name, passcode, createdAt string
	var updatedAt interface{}
	var id int
	check := fmt.Sprintf(`SELECT id, passcode, name, createdAt, updatedAt FROM master_kasir WHERE name = '%s' AND deletedAt is null;`, request.Name)
	dbconnect.QueryRow(check).Scan(
		&id,
		&passcode,
		&name,
		&createdAt,
		&updatedAt,
	)
	if id > 0 {
		d := &models.ResponseKasir{
			Success: false,
			Message: "Cashier Already Exist",
			Data: &models.Kasir{
				Passcode:  passcode,
				ID:        id,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		INSERT INTO master_kasir (passcode, name, createdAt) VALUES('123456', '%s', now());
	`,
		request.Name,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM master_kasir WHERE deletedAt is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastID(id)
	return
}

func UpdateKasir(request *models.RequestKasir, id int) (response *models.ResponseKasir) {
	query := fmt.Sprintf(`
		UPDATE master_kasir SET name = '%s', updatedAt = now() WHERE id = %d;
	`,
		request.Name,
		id,
	)
	dbconnect.Exec(query)
	response = lastID(id)
	return
}

func DeleteKasir(id int) (response *models.ResponseKasir) {
	query := fmt.Sprintf(`
		UPDATE master_kasir SET deletedAt = now() WHERE id = %d;
	`,
		id,
	)
	dbconnect.Exec(query)
	response = lastID(id)
	return
}

func DetailKasir(id int) (response *models.ResponseData) {
	var name string
	query := fmt.Sprintf(`SELECT id, name FROM master_kasir WHERE id = %d AND deletedAt is null;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
	)
	if name != "" {
		d := &models.ResponseData{
			Success: true,
			Message: "Success",
			Data: &models.DataKasir{
				CashierId: id,
				Name:      name,
			},
		}
		response = d
		return
	}
	d := &models.ResponseData{
		Success: false,
		Message: "Cashier Not Found",
		Data:    &models.DataKasir{},
	}
	response = d
	return
}

func ListKasir(limit, skip int) (response *models.ResponseListData) {
	arrData := []models.DataKasir{}
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM master_kasir WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	query := fmt.Sprintf(`
		SELECT
			id,
			name
		FROM
			master_kasir
		WHERE
			deletedAt is null
		LIMIT %d OFFSET %d;
	`,
		limit,
		skip,
	)

	rowsQ, err := dbconnect.Query(query)
	rowData := models.DataKasir{}

	if err != nil {
		log.Printf(err.Error())
		return
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.CashierId,
			&rowData.Name,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}
		arrData = append(arrData, rowData)
	}

	d := &models.ResponseListData{
		Success: true,
		Message: "Success",
		Data: &models.ListKasir{
			Cashiers: arrData,
			Meta: &models.MetaData{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
