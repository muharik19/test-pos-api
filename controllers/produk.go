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

func CreatedProduk(c *gin.Context) {
	var request *models.RequestProduk
	var response *models.ResponseProduk
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
	log.WithField("request", request).Info("post create produk request received")

	response = repo.CreatedProduk(request)
	c.JSON(http.StatusOK, response)
	return
}

func UpdateProduk(c *gin.Context) {
	var request *models.RequestProduk
	var response *models.ResponseProduk
	var err error

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("productId")
	intVar, _ := strconv.Atoi(id)
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithField("request", request).Info("put update produk request received")

	response = repo.UpdateProduk(request, intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DeleteProduk(c *gin.Context) {
	var response *models.ResponseProduk

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("productId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("delete remove produk request received")
	response = repo.DeleteProduk(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func DetailProduk(c *gin.Context) {
	var response *models.ResponseDataProduk

	token := c.Request.Header["Authorization"][0]
	isAuthorized := utils.IsAuthorized(token)

	if isAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	id := c.Param("productId")
	intVar, _ := strconv.Atoi(id)
	log.WithField("request", intVar).Info("get detail produk request received")
	response = repo.DetailProduk(intVar)
	if response.Success == false {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func ListProduk(c *gin.Context) {
	var response *models.ResponseListDataProduk

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
	id := c.Query("categoryId")
	intVar, _ := strconv.Atoi(id)
	q := c.Query("q")
	log.WithField("request", intLimit).Info("get list produk request received")
	response = repo.ListProduk(intLimit, intSkip, intVar, q)
	c.JSON(http.StatusOK, response)
	return
}
