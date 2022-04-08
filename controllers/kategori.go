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

func CreatedKategori(c *gin.Context) {
	var request *models.RequestKategori
	var response *models.ResponseKategori
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
	log.WithField("request", request).Info("post create kategori request received")

	response = repo.CreatedKategori(request)
	c.JSON(http.StatusOK, response)
	return
}

func UpdateKategori(c *gin.Context) {
	var request *models.RequestKategori
	var response *models.ResponseKategori
	var err error

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("categoryId")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("put update kategori request received")

	response = repo.UpdateKategori(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteKategori(c *gin.Context) {
	var response *models.ResponseKategori

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("categoryId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("delete remove kategori request received")
	response = repo.DeleteKategori(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailKategori(c *gin.Context) {
	var response *models.ResponseDataKategori

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("categoryId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail kategori request received")
	response = repo.DetailKategori(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListKategori(c *gin.Context) {
	var response *models.ResponseListDataKategori

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
	log.WithField("request", intLimit).Info("get list kategori request received")
	response = repo.ListKategori(intLimit, intSkip)
	c.JSON(http.StatusOK, response)
	return
}
