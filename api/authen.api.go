package api

import "github.com/gin-gonic/gin"

func SetupAuthenAPI(router *gin.Engine) 
{
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

// Login - login api
func login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "login"})
}

// Register - register api
func register(c *gin.Context) {
	c.JSON(200, gin.H{"result": "Register"})
}
