package main

import (
	home "GolangTestAPI/API"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.PUT("/adduser", home.AddUser)
	router.POST("/updateuser", home.UpdateUser)
	router.DELETE("/deleteuser", home.DeleteUser)
	router.GET("/getuser", home.GetUser)

	router.POST("/login", home.Login)
	// router.GET("/protected", home.AuthMiddleware(), home.Protected)

	router.Run("localhost:8080")

	// for Https
	// router.RunTLS(":8080", "cert.pem", "key.pem")
}
