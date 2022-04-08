package repositories

import (
	"fmt"
	"log"

	"github.com/devcode-pos/models"

	dbconnect "github.com/devcode-pos/databases"
)

func lastKategoriID(id int) (response *models.ResponseKategori) {
	var name, createdAt string
	var updatedAt interface{}
	query := fmt.Sprintf(`SELECT id, name, createdAt, updatedAt FROM master_kategori WHERE id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)
	if name != "" {
		d := &models.ResponseKategori{
			Success: true,
			Message: "Success",
			Data: &models.Kategori{
				ID:        id,
				Name:      name,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseKategori{
		Success: false,
		Message: "Category Not Found",
		Data:    &models.Kategori{},
	}
	response = d
	return
}

func CreatedKategori(request *models.RequestKategori) (response *models.ResponseKategori) {
	var name, createdAt string
	var updatedAt interface{}
	var id int
	check := fmt.Sprintf(`SELECT id, name, createdAt, updatedAt FROM master_kategori WHERE name = '%s' AND deletedAt is null;`, request.Name)
	dbconnect.QueryRow(check).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)
	if id > 0 {
		d := &models.ResponseKategori{
			Success: false,
			Message: "Category Already Exist",
			Data: &models.Kategori{
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
		INSERT INTO master_kategori (name, createdAt) VALUES('%s', now());
	`,
		request.Name,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM master_kategori WHERE deletedAt is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastKategoriID(id)
	return
}

func UpdateKategori(request *models.RequestKategori, id int) (response *models.ResponseKategori) {
	query := fmt.Sprintf(`
		UPDATE master_kategori SET name = '%s', updatedAt = now() WHERE id = %d;
	`,
		request.Name,
		id,
	)
	dbconnect.Exec(query)
	response = lastKategoriID(id)
	return
}

func DeleteKategori(id int) (response *models.ResponseKategori) {
	query := fmt.Sprintf(`
		UPDATE master_kategori SET deletedAt = now() WHERE id = %d;
	`,
		id,
	)
	dbconnect.Exec(query)
	response = lastKategoriID(id)
	return
}

func DetailKategori(id int) (response *models.ResponseDataKategori) {
	var name string
	query := fmt.Sprintf(`SELECT id, name FROM master_kategori WHERE id = %d AND deletedAt is null;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
	)
	if name != "" {
		d := &models.ResponseDataKategori{
			Success: true,
			Message: "Success",
			Data: &models.DataKategori{
				CategoryId: id,
				Name:       name,
			},
		}
		response = d
		return
	}
	d := &models.ResponseDataKategori{
		Success: false,
		Message: "Category Not Found",
		Data:    &models.DataKategori{},
	}
	response = d
	return
}

func ListKategori(limit, skip int) (response *models.ResponseListDataKategori) {
	arrData := []models.DataKategori{}
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM master_kategori WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	query := fmt.Sprintf(`
		SELECT
			id,
			name
		FROM
			master_kategori
		WHERE
			deletedAt is null
		LIMIT %d OFFSET %d;
	`,
		limit,
		skip,
	)

	rowsQ, err := dbconnect.Query(query)
	rowData := models.DataKategori{}

	if err != nil {
		log.Printf(err.Error())
		return
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.CategoryId,
			&rowData.Name,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}
		arrData = append(arrData, rowData)
	}

	d := &models.ResponseListDataKategori{
		Success: true,
		Message: "Success",
		Data: &models.ListKategori{
			Categories: arrData,
			Meta: &models.MetaDataKategori{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
