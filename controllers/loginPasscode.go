package controllers

import (
	"net/http"
	"strconv"

	"github.com/devcode-pos/models"
	repo "github.com/devcode-pos/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func VerifyLoginPasscode(c *gin.Context) {
	var request *models.RequestPasscode
	var response *models.LoginModel
	var err error

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("post verify login passcode request received")

	response = repo.VerifyLoginPasscode(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func VerifyLogoutPasscode(c *gin.Context) {
	var request *models.RequestPasscode
	var response *models.Logout
	var err error

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("post verify logout passcode request received")

	response = repo.VerifyLogoutPasscode(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
