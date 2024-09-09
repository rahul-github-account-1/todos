package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahul-github-account-1/todo/models"
	"github.com/rahul-github-account-1/todo/services"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/users", func(c *gin.Context) {
		fmt.Println("working");
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in binding the user body with struct in routes package": err.Error()})
			return
		}
		createdUser, err := services.CreateUser(db, user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errror in creating user in routes package ": err})
			return
		}
		c.JSON(http.StatusOK, createdUser)
	})
	router.GET("users/:id",func(c *gin.Context){
		userID,err:= strconv.ParseUint(c.Param("id"),10,32);
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return	
		}
		userIDUint:=uint(userID)
		user,err:=services.GetUser(db,userIDUint);
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error in getting the user details: ":err});
		}
		c.JSON(http.StatusOK,user);
	})

	router.DELETE("users/:id", func(c *gin.Context) {
		// userID := c.Param("id")

		userID,err:= strconv.ParseUint(c.Param("id"),10,32);
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return	
		}
		userIDUint:=uint(userID)



		if err := services.SoftDeleteUser(db, userIDUint); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to soft delete user in routes package"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User and associated todos are soft deleted"})
	})

}
