package models

type RequestProduk struct {
	CategoryId int     `json:"categoryId" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Image      string  `json:"image" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Stock      float64 `json:"stock" binding:"required"`
}

type Produk struct {
	ID         int         `json:"id"`
	CategoryId int         `json:"categoryId"`
	Name       string      `json:"name"`
	Sku        string      `json:"sku"`
	Image      string      `json:"image"`
	Price      float64     `json:"price"`
	Stock      float64     `json:"stock"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  interface{} `json:"updatedAt"`
}

type ResponseProduk struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    *Produk `json:"data"`
}

type KategoriProduk struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
}

type DataProduk struct {
	ProductId int             `json:"productId"`
	Sku       string          `json:"sku"`
	Name      string          `json:"name"`
	Stock     float64         `json:"stock"`
	Price     float64         `json:"price"`
	Image     string          `json:"image"`
	Category  *KategoriProduk `json:"category"`
	Discount  interface{}     `json:"discount"`
}

type ResponseDataProduk struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *DataProduk `json:"data"`
}

type MetaDataProduk struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListProduk struct {
	Products *[]DataProduk   `json:"products"`
	Meta     *MetaDataProduk `json:"meta"`
}

type ResponseListDataProduk struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *ListProduk `json:"data"`
}
