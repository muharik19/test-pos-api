package models

type RequestPasscode struct {
	Passcode string `json:"passcode" binding:"required"`
}

type DataToken struct {
	Token string `json:"token"`
}

type LoginModel struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *DataToken `json:"data"`
}

type UserModel struct {
	Passcode string `json:"passcode"`
}

type UserTypeModel struct {
	Passcode string `json:"passcode"`
}

type Logout struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
