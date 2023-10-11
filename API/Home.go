package home

import (
	"GolangTestAPI/Database"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AddUser(c *gin.Context) {

	var requestBody Database.UsersArr

	if err := c.BindJSON(&requestBody); err != nil {
		// Handle the error, for example, return an error response to the client
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Database.Insert(requestBody)
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func GetUser(c *gin.Context) {

	user, err := Database.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {

	var deleteId Database.UserDelete

	if err := c.BindJSON(&deleteId); err != nil {
		// Handle the error, for example, return an error response to the client
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Database.Delete(deleteId)
	c.JSON(http.StatusOK, gin.H{"message": "User delete successfully"})

}

func UpdateUser(c *gin.Context) {

	var requestBody Database.UsersArr

	if err := c.BindJSON(&requestBody); err != nil {
		// Handle the error, for example, return an error response to the client
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Database.Update(requestBody)
	c.JSON(http.StatusOK, gin.H{"message": "User Update successfully"})
}
