package controller

import (
	"GinIntroduction/internal/config"
	"GinIntroduction/internal/entity"
	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {
	var products []entity.Product
	config.DB.Find(&products)
	c.JSON(200, gin.H{
		"products": products,
	})
}

func ProductPost(c *gin.Context) {
	var product entity.Product
	c.Bind(&product)
	config.DB.Create(&product)
	c.JSON(200, gin.H{
		"POST succeed": product,
	})
}

func ProductPut(c *gin.Context) {
	var product entity.Product
	c.Bind(&product)
	config.DB.Save(&product)
	c.JSON(200, gin.H{
		"PUT succeed": product,
	})
}

func ProductGetByID(c *gin.Context) {

}

func ProductDelete(c *gin.Context) {

}
