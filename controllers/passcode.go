package controllers

import (
	"net/http"
	"strconv"

	"github.com/devcode-pos/models"
	repo "github.com/devcode-pos/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DetailPasscode(c *gin.Context) {
	var response *models.ResponsePasscode

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get passcode request received")
	response = repo.DetailPasscode(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
