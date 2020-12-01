package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func myInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
	}
}

// SetupProductAPI - setup api product
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", myInterceptor, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

// Llis - list product api
func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product"})
}

// Create - create product api
func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}
