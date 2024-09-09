package main

import (
	"fmt"
	"log"
	"github.com/rahul-github-account-1/todo/routes"
	"github.com/gin-gonic/gin"
	"github.com/rahul-github-account-1/todo/config"
	"github.com/rahul-github-account-1/todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	config,err:=config.LoadConf();
	if err!=nil{
		log.Fatal("Unable to Load config in main package: ",err);
	}

	// database_conn_string:=fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",config.DBUser,config.DBPassword,config.DBName,config.DBHost,config.DBPort);
	database_conn_string:="postgresql://neondb_owner:Cj9lMErfik1d@ep-aged-field-a5yms3hz.us-east-2.aws.neon.tech/neondb?sslmode=require";
	
	db,err:=gorm.Open(postgres.Open(database_conn_string),&gorm.Config{});
	if err!=nil{
		log.Fatal("Failed to connect to db in main package: ",err);
	}
	if err:=db.AutoMigrate(&models.User{},&models.Todo{}); err!=nil{
		log.Fatal("Unable to auto migrate the models in main package: ",err);
	}
	fmt.Println("JWT Key: ",config.JWTKey);

	router:=gin.Default();
	routes.InitializeRoutes(router,db);
	log.Fatal(router.Run(":8080"));

	
}