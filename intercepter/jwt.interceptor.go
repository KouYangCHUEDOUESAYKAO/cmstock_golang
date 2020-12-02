package intercepter

import (
	"fmt"
	"main/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "8@7654321"

// JwtSign - jwtsigned api
func JwtSign(payload model.User) string {
	atClaims := jwt.MapClaims{}
	// payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}

// JwtVerify - call this method
func JwtVerify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["jwt_username"])
		level := fmt.Sprintf("%v", claims["jwt_level"])
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("jwt_level", level)
		c.Next()
	} else {
		c.JSON(200, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}
