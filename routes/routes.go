package routes

import (

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine,db *gorm.DB){

	UserRoutes(router,db);
	TODORoutes(router,db);

}