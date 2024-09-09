package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahul-github-account-1/todo/models"
	"github.com/rahul-github-account-1/todo/services"
	"gorm.io/gorm"
)

func TODORoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/todos", func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in binding the todos in post in routes package": err.Error()})
			return
		}

		createdTodo, err := services.CreateTodo(db, todo.UserID, todo.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error ": "Unable to create todo in routes package"})
			return
		}

		c.JSON(http.StatusOK, createdTodo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		var todoUpdate map[string]interface{}
		if err := c.ShouldBindJSON(&todoUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in binding todoUpdate body with struct in routes package": err.Error()})
			return
		}

		todoId,err:= strconv.ParseUint(c.Param("id"),10,32);
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return	
		}
		todoIdUint:=uint(todoId)
		if err := services.UpdateTodo(db, todoIdUint, todoUpdate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update todo in routes package"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		todoId,err:= strconv.ParseUint(c.Param("id"),10,32);
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return	
		}
		todoIdUint:=uint(todoId)
		
		if err := services.SoftDeleteTodo(db, todoIdUint); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to soft delete todo in routes package"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo soft-deleted successfully"})
	})

}
