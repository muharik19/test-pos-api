package controllers

import (
	"net/http"
	"strconv"

	"github.com/devcode-pos/models"
	repo "github.com/devcode-pos/repositories"
	"github.com/devcode-pos/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreatedKasir(c *gin.Context) {
	var request *models.RequestKasir
	var response *models.ResponseKasir
	var err error

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("post create kasir request received")

	response = repo.CreatedKasir(request)
	c.JSON(http.StatusOK, response)
	return
}

func UpdateKasir(c *gin.Context) {
	var request *models.RequestKasir
	var response *models.ResponseKasir
	var err error

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("put update kasir request received")

	response = repo.UpdateKasir(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteKasir(c *gin.Context) {
	var response *models.ResponseKasir

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("delete remove kasir request received")
	response = repo.DeleteKasir(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailKasir(c *gin.Context) {
	var response *models.ResponseData

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("cashierId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail kasir request received")
	response = repo.DetailKasir(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListKasir(c *gin.Context) {
	var response *models.ResponseListData

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	limit := c.Query("limit")
	intLimit, _ := strconv.Atoi(limit)
	skip := c.Query("skip")
	intSkip, _ := strconv.Atoi(skip)
	log.WithField("request", intLimit).Info("get list kasir request received")
	response = repo.ListKasir(intLimit, intSkip)
	c.JSON(http.StatusOK, response)
	return
}
