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
	id := c.Param("id")
	var product entity.Product

	config.DB.First(&product, id)

	var updateProduct entity.Product
	c.Bind(&updateProduct)
	config.DB.Model(&product).Updates(entity.Product{
		Code:  updateProduct.Code,
		Price: updateProduct.Price,
	})
	c.JSON(200, gin.H{
		"PUT succeed": product,
	})
}

func ProductGetByID(c *gin.Context) {
	id := c.Param("id")
	var product entity.Product
	config.DB.First(&product, id)

	c.JSON(200, gin.H{
		"GET by id succeed": product,
	})
}

func ProductDelete(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&entity.Product{}, id)
	c.JSON(200, gin.H{
		"DELETE succeed": id,
	})
}
