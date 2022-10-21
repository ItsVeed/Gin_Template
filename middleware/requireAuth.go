package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ItsVeed/Gin_Template/initializers"
	"github.com/ItsVeed/Gin_Template/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get cookie off request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println(1)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Print(claims)
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println(2)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attatch to request
		c.Set("user", &user)

		// Continue
		c.Next()
	} else {
		fmt.Println(3)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
