package repositories

import (
	"fmt"

	"github.com/devcode-pos/utils"

	dbconnect "github.com/devcode-pos/databases"
	"github.com/devcode-pos/models"
)

func VerifyLoginPasscode(request *models.RequestPasscode, id int) (response *models.LoginModel) {
	var token string
	var passcode string
	query := fmt.Sprintf(`SELECT passcode FROM master_kasir WHERE passcode = '%s' AND deletedAt is null;`, request.Passcode)
	dbconnect.QueryRow(query).Scan(
		&passcode,
	)

	if passcode == "" {
		d := &models.LoginModel{
			Success: false,
			Message: "Passcode Not Match",
			Data:    &models.DataToken{},
		}
		response = d
		return
	}

	attrUser := models.UserTypeModel{
		Passcode: request.Passcode,
	}

	token = utils.GenerateToken(attrUser)

	d := &models.LoginModel{
		Success: true,
		Message: "Success",
		Data: &models.DataToken{
			Token: token,
		},
	}
	response = d
	return
}

func VerifyLogoutPasscode(request *models.RequestPasscode, id int) (response *models.Logout) {
	var passcode string
	query := fmt.Sprintf(`SELECT passcode FROM master_kasir WHERE passcode = '%s' AND deletedAt is null;`, request.Passcode)
	dbconnect.QueryRow(query).Scan(
		&passcode,
	)

	if passcode == "" {
		d := &models.Logout{
			Success: false,
			Message: "Passcode Not Match",
		}
		response = d
		return
	}
	d := &models.Logout{
		Success: true,
		Message: "Success",
	}
	response = d
	return
}
