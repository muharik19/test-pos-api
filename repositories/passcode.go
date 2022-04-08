package repositories

import (
	"fmt"

	"github.com/devcode-pos/models"

	dbconnect "github.com/devcode-pos/databases"
)

func DetailPasscode(id int) (response *models.ResponsePasscode) {
	var passcode string
	query := fmt.Sprintf(`SELECT passcode FROM master_kasir WHERE id = %d AND deletedAt is null;`, id)
	dbconnect.QueryRow(query).Scan(
		&passcode,
	)
	if passcode != "" {
		d := &models.ResponsePasscode{
			Success: true,
			Message: "Success",
			Data: &models.DataPasscode{
				Passcode: passcode,
			},
		}
		response = d
		return
	}
	d := &models.ResponsePasscode{
		Success: false,
		Message: "Cashier Not Found",
		Data:    &models.DataPasscode{},
	}
	response = d
	return
}
