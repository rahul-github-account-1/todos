package config

import (
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)


type Config struct{
	DBUser string
	DBPassword string
	DBName	string
	DBHost	string
	DBPort	string
	JWTKey	string
}

func LoadConf()(*Config,error){
	if err:= godotenv.Load();err!=nil{
		log.Fatal("Error loading .env file in config package: ",err);
	}
	config:=&Config{
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		JWTKey: os.Getenv("JWTKey"),

	}
	if config.DBUser==""||config.DBName==""||config.DBPassword==""||config.DBPort==""||config.JWTKey==""||config.DBHost==""{
		return nil,fmt.Errorf("missing required env variables in config package");
	}
	return config,nil

}