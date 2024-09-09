package services

import (
	"errors"
	"time"

	"github.com/rahul-github-account-1/todo/models"
	"gorm.io/gorm"
)
func CreateTodo(db *gorm.DB, userID uint, title string)(*models.Todo,error){
	todo:=models.Todo{
		Title: title,
		UserID: userID,
	}
	if err:=db.Create(&todo).Error; err!=nil{
		return nil,err;
	}

	return &todo,nil;
}

func GetTodos(db *gorm.DB,userID uint)([]models.Todo,error){
	var todos []models.Todo;
	if err:=db.Where("user_id = ? AND deleted_at is NULL",userID).Find(&todos).Error;err!=nil{
		return nil,err;
	}
	return todos,nil;

}

func GetTodo(db *gorm.DB,todoID uint)(*models.Todo,error){
	var todo models.Todo;
	if err:=db.Where("id = ? AND deleted_at is NULL",todoID).First(&todo).Error; err!=nil{
		return nil,err;
	}
	return &todo,nil;
}

func UpdateTodo(db *gorm.DB, todoID uint, updates map[string]interface{})error{
	var todo models.Todo;
	if err:=db.First(&todo,todoID).Error; err!=nil{
		return errors.New("TODO not found for updating");
	}
	if err:=db.Model(&todo).Updates(updates).Error; err!=nil{
		return err;
	}
	return nil;
}
func SoftDeleteTodo(db *gorm.DB, todoID uint)error{
	var todo models.Todo;
	if err:=db.First(&todo,todoID).Error;err!=nil{
		return errors.New("todo not found for deleting");
	}
	deletedAt:=time.Now();
	if err:=db.Model(&todo).Update("deleted_at",deletedAt).Error;err!=nil{
		return err;
	}
	return nil;
}
