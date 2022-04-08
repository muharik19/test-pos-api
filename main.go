package main

import (
	cm "github.com/devcode-pos/common"
	"github.com/devcode-pos/databases"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	controller "github.com/devcode-pos/controllers"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	cm.InitConfig()

	db, err := databases.ConnectDb()
	if err == nil {
		defer db.Close()
		log.Info("MySQL Database Connected..")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	g := r.Group(cm.Config.RootURL)
	{
		g.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		g.POST("cashiers", controller.CreatedKasir)
		g.PUT("/cashiers/:cashierId", controller.UpdateKasir)
		g.DELETE("/cashiers/:cashierId", controller.DeleteKasir)
		g.GET("/cashiers/:cashierId", controller.DetailKasir)
		g.GET("/cashiers", controller.ListKasir)
		g.POST("categories", controller.CreatedKategori)
		g.PUT("/categories/:categoryId", controller.UpdateKategori)
		g.DELETE("/categories/:categoryId", controller.DeleteKategori)
		g.GET("/categories/:categoryId", controller.DetailKategori)
		g.GET("/categories", controller.ListKategori)
		g.POST("products", controller.CreatedProduk)
		g.PUT("/products/:productId", controller.UpdateProduk)
		g.DELETE("/products/:productId", controller.DeleteProduk)
		g.GET("/products/:productId", controller.DetailProduk)
		g.GET("/products", controller.ListProduk)
		g.GET("/cashiers/:cashierId/passcode", controller.DetailPasscode)
		g.POST("/cashiers/:cashierId/login", controller.VerifyLoginPasscode)
		g.POST("/cashiers/:cashierId/logout", controller.VerifyLogoutPasscode)
	}
	// Start serverlog.Info()
	log.Info("Staring server ...")
	r.Run(":" + cm.Config.Port)
}
