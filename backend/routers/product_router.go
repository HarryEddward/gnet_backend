package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterProductsRouter(router *gin.Engine) {

	productsGroup := router.Group("/products")
	{
		productsGroup.GET("/", general_products)
		productsGroup.GET("/specific/:id", specific_products)
	}

}

func general_products(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Lista de productos",
	})
}

func specific_products(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("Producto Especifico: %s, IP: %s", c.Param("id"), c.ClientIP()),
	})
}
