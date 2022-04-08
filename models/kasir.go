package models

type RequestKasir struct {
	Name string `json:"name" binding:"required"`
}

type Kasir struct {
	Passcode  string      `json:"passcode"`
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt interface{} `json:"updatedAt"`
}

type ResponseKasir struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Kasir `json:"data"`
}

type DataKasir struct {
	CashierId int    `json:"cashierId"`
	Name      string `json:"name"`
}

type ResponseData struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *DataKasir `json:"data"`
}

type MetaData struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListKasir struct {
	Cashiers []DataKasir `json:"cashiers"`
	Meta     *MetaData   `json:"meta"`
}

type ResponseListData struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *ListKasir `json:"data"`
}
