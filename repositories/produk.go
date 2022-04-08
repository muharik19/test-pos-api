package repositories

import (
	"fmt"
	"log"

	"github.com/devcode-pos/models"

	dbconnect "github.com/devcode-pos/databases"
)

func lastProdukID(id int) (response *models.ResponseProduk) {
	var name, createdAt, sku, image string
	var updatedAt interface{}
	var kategori_id int
	var price, stock float64
	query := fmt.Sprintf(`SELECT id, kategori_id, name, sku, image, price, stock, createdAt, updatedAt FROM master_produk WHERE id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&kategori_id,
		&name,
		&sku,
		&image,
		&price,
		&stock,
		&createdAt,
		&updatedAt,
	)
	if name != "" {
		d := &models.ResponseProduk{
			Success: true,
			Message: "Success",
			Data: &models.Produk{
				ID:         id,
				CategoryId: kategori_id,
				Name:       name,
				Sku:        sku,
				Image:      image,
				Price:      price,
				Stock:      stock,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseProduk{
		Success: false,
		Message: "Product Not Found",
		Data:    &models.Produk{},
	}
	response = d
	return
}

func CreatedProduk(request *models.RequestProduk) (response *models.ResponseProduk) {
	var name, createdAt, sku, image string
	var updatedAt interface{}
	var kategori_id, id int
	var price, stock float64
	check := fmt.Sprintf(`SELECT id, kategori_id, name, sku, image, price, stock, createdAt, updatedAt FROM master_produk WHERE name = '%s' AND deletedAt is null;`, request.Name)
	dbconnect.QueryRow(check).Scan(
		&id,
		&kategori_id,
		&name,
		&sku,
		&image,
		&price,
		&stock,
		&createdAt,
		&updatedAt,
	)
	if id > 0 {
		d := &models.ResponseProduk{
			Success: false,
			Message: "Product Already Exist",
			Data: &models.Produk{
				ID:         id,
				CategoryId: kategori_id,
				Name:       name,
				Sku:        sku,
				Image:      image,
				Price:      price,
				Stock:      stock,
				CreatedAt:  createdAt,
				UpdatedAt:  updatedAt,
			},
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		INSERT INTO master_produk (kategori_id, name, sku, image, price, stock, createdAt) VALUES(%d, '%s', 'ID007', '%s', %f, %f, now());
	`,
		request.CategoryId,
		request.Name,
		request.Image,
		request.Price,
		request.Stock,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM master_produk WHERE deletedAt is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastProdukID(id)
	return
}

func UpdateProduk(request *models.RequestProduk, id int) (response *models.ResponseProduk) {
	query := fmt.Sprintf(`
		UPDATE master_produk SET kategori_id = %d, name = '%s', image = '%s', price = %f, stock = %f, updatedAt = now() WHERE id = %d;
	`,
		request.CategoryId,
		request.Name,
		request.Image,
		request.Price,
		request.Stock,
		id,
	)
	dbconnect.Exec(query)
	response = lastProdukID(id)
	return
}

func DeleteProduk(id int) (response *models.ResponseProduk) {
	query := fmt.Sprintf(`
		UPDATE master_produk SET deletedAt = now() WHERE id = %d;
	`,
		id,
	)
	dbconnect.Exec(query)
	response = lastProdukID(id)
	return
}

func DetailProduk(id int) (response *models.ResponseDataProduk) {
	var name, sku, image, kategori_name string
	var stock, price float64
	var kategori_id int
	query := fmt.Sprintf(`
		SELECT
			mp.id,
			mp.sku,
			mp.name,
			mp.stock,
			mp.price,
			mp.image,
			mk.kategori_id,
			mk.kategori_name
		FROM
			master_produk mp
		INNER JOIN (
			select
				mk.id as kategori_id,
				mk.name as kategori_name
			from
				master_kategori mk
			where mk.deletedAt is null) as mk on
			mk.kategori_id = mp.kategori_id
		where mp.deletedAt is null AND mp.id = %d;
	`,
		id,
	)
	dbconnect.QueryRow(query).Scan(
		&id,
		&sku,
		&name,
		&stock,
		&price,
		&image,
		&kategori_id,
		&kategori_name,
	)
	if name != "" {
		d := &models.ResponseDataProduk{
			Success: true,
			Message: "Success",
			Data: &models.DataProduk{
				ProductId: id,
				Sku:       sku,
				Name:      name,
				Stock:     stock,
				Price:     price,
				Image:     image,
				Category: &models.KategoriProduk{
					CategoryId: kategori_id,
					Name:       kategori_name,
				},
				Discount: nil,
			},
		}
		response = d
		return
	}
	d := &models.ResponseDataProduk{
		Success: false,
		Message: "Category Not Found",
		Data:    &models.DataProduk{},
	}
	response = d
	return
}

func ListProduk(limit, skip, categoryId int, q string) (response *models.ResponseListDataProduk) {
	arrData := []models.DataProduk{}
	catId := ""
	searchs := ""
	persen := "%"
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM master_produk WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	if categoryId > 0 {
		catId = fmt.Sprintf(`AND mk.kategori_id = %d`, categoryId)
	}
	if q != "" {
		searchs = fmt.Sprintf(`AND mp.name LIKE '%s%s%s'`, persen, q, persen)
	}
	query := fmt.Sprintf(`
		SELECT
			mp.id,
			mp.sku,
			mp.name,
			mp.stock,
			mp.price,
			mp.image,
			mk.kategori_id,
			mk.kategori_name
		FROM
			master_produk mp
		INNER JOIN (
			SELECT
				mk.id as kategori_id,
				mk.name as kategori_name
			FROM
				master_kategori mk
			WHERE mk.deletedAt is null) as mk on
			mk.kategori_id = mp.kategori_id
		WHERE mp.deletedAt is null %s %s
		LIMIT %d OFFSET %d;
	`,
		catId,
		searchs,
		limit,
		skip,
	)

	rowsQ, err := dbconnect.Query(query)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	var prodId, katId int
	var sku, prodName, prodImage, katName string
	var stock, price float64
	for rowsQ.Next() {
		err = rowsQ.Scan(
			&prodId,
			&sku,
			&prodName,
			&stock,
			&price,
			&prodImage,
			&katId,
			&katName,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		rowData := models.DataProduk{
			ProductId: prodId,
			Sku:       sku,
			Name:      prodName,
			Stock:     stock,
			Price:     price,
			Image:     prodImage,
			Category: &models.KategoriProduk{
				CategoryId: katId,
				Name:       katName,
			},
			Discount: nil,
		}
		arrData = append(arrData, rowData)
	}

	d := &models.ResponseListDataProduk{
		Success: true,
		Message: "Success",
		Data: &models.ListProduk{
			Products: &arrData,
			Meta: &models.MetaDataProduk{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
