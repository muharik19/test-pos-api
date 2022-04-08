package models

type RequestKategori struct {
	Name string `json:"name" binding:"required"`
}

type Kategori struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt interface{} `json:"updatedAt"`
}

type ResponseKategori struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *Kategori `json:"data"`
}

type DataKategori struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
}

type ResponseDataKategori struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *DataKategori `json:"data"`
}

type MetaDataKategori struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListKategori struct {
	Categories []DataKategori    `json:"categories"`
	Meta       *MetaDataKategori `json:"meta"`
}

type ResponseListDataKategori struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *ListKategori `json:"data"`
}
