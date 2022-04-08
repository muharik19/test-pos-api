package utils

import (
	"fmt"
	"log"
	"time"

	cm "github.com/devcode-pos/common"
	"github.com/devcode-pos/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// UserClaims struct
type UserClaims struct {
	models.UserModel
	jwt.StandardClaims
}

// GenerateToken func
func GenerateToken(modelUserType models.UserTypeModel) string {
	// jwtExpired := cm.Config.JwtExpired
	LOGIN_EXPIRATION_DURATION := time.Duration(5) * time.Minute
	// uom := jwtExpired[len(jwtExpired)-1:]
	// expiredToken, _ := strconv.Atoi(strings.Replace(jwtExpired, uom, "", -1))

	// timeHourMinutes := time.Hour
	// if uom == "d" {
	// 	timeHourMinutes = time.Hour * 24
	// }

	// expiredAt := time.Now().Add(timeHourMinutes * time.Duration(expiredToken)).Unix()
	expiredAt := time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix()

	claims := UserClaims{
		models.UserModel{
			Passcode: modelUserType.Passcode,
		},
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    modelUserType.Passcode,
		},
	}

	var signingKey = []byte(cm.Config.JwtKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		log.Println("Generate Token Failed..")
		return ""
	}

	return tokenString
}

// JwtDecode func
func JwtDecode(token string) (*jwt.Token, error) {
	var signingKey = []byte(cm.Config.JwtKey)

	return jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
}

// IsAuthorized func
func IsAuthorized(tokenString string) bool {
	token, err := JwtDecode(tokenString)
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims == nil {
			return false
		}
		return true
	}

	return false
}

// JwtClaim func
func JwtClaim(tokenString string) models.UserModel {
	var signingKey = []byte(cm.Config.JwtKey)

	users := models.UserModel{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		fmt.Printf("Error JWT Claim")
	}

	users.Passcode = claims["passcode"].(string)

	return users
}
