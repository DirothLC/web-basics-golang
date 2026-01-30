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
	router.GET("/product", controller.ProductIndex)
	router.POST("/product", controller.ProductPost)
	router.PUT("/product/:id", controller.ProductPut)
	router.DELETE("/product/:id", controller.ProductDelete)
	router.GET("/product/:id", controller.ProductGetByID)

	return router
}
