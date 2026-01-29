package server

import (
	"GinIntroduction/internal/controller"
	"GinIntroduction/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pingHandler := handler.NewPingHandler()
	router.GET("/ping", pingHandler.Ping)
	router.GET("/products", controller.ProductIndex)
	router.POST("/product", controller.ProductPost)
	router.PUT("/product", controller.ProductPut)
	router.DELETE("/product", controller.ProductDelete)
	router.GET("/product", controller.ProductGetByID)

	return router
}
