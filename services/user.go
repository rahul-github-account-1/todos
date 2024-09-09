package services

import (
	"errors"
	"time"

	"github.com/rahul-github-account-1/todo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB,email, password string)(*models.User,error) {
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);
	if err!=nil{
		return nil,err;
	}
	user:=models.User{
		Email: email,
		Password: string(hashedPassword),
	}
	if err:=db.Create(&user).Error; err!=nil{
		return nil,err;
	}
	return &user,nil;
}

func SoftDeleteUser(db *gorm.DB,userID uint)error {
	var user models.User;
	if err:=db.First(&user,userID).Error; err!=nil{
		return errors.New("Not able to find user with userID while deleting the user");
	}
	deletedAt:=time.Now();
	if err:=db.Model(&user).Update("deleted_at",deletedAt).Error; err!=nil{
		return err;
	}
	if err:=db.Model(&models.Todo{}).Where("user_id = ?",userID).Update("deleted_at",deletedAt).Error; err!=nil{
		return err;
	}
	return nil;
}
func GetUser(db *gorm.DB,userID uint)(*models.User,error) {
	var user models.User;
	if err:=db.Where("id = ? AND deleted_at is NULL",userID).First(&user).Error;err!=nil{
		return nil,err;
	}
	return &user,nil;
}

func CheckUserExists(db *gorm.DB,email string)(bool,error) {
	var user models.User;
	if err := db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
        return false, err
    }

    return true, nil

}
