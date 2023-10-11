package home

import (
	"time"

	"GolangTestAPI/Database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("test01")

func Login(c *gin.Context) {

	username := c.Request.Header.Get("Username")
	password := c.Request.Header.Get("Password")

	user := Database.LoginAuth{Username: username, Password: password}

	token := GenerateToken(user)

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GenerateToken(user Database.LoginAuth) string {
	expTime := time.Now().Add(time.Hour * 1)
	expTimeString := expTime.Format(time.RFC3339)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "id":       user.ID,
		"username": user.Username,
		"password": user.Password,
		"exp":      expTime, // Token expiration time
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	Database.TokenDB(tokenString, expTimeString)

	return tokenString
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(401, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return secretKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			c.JSON(401, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func Protected(c *gin.Context) {
// 	c.JSON(200, gin.H{"message": "This is a protected route"})
// }
