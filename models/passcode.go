package models

type DataPasscode struct {
	Passcode string `json:"passcode"`
}

type ResponsePasscode struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *DataPasscode `json:"data"`
}
