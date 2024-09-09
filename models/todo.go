package models
import "gorm.io/gorm"
import "time"
type Todo struct{
    gorm.Model
    Title     string  	 `json:"title" gorm:"not null"`
    Completed bool       `json:"completed" gorm:"default:false"`
    UserID    uint       `json:"user_id"` 
	DeletedAt *time.Time `json:"deleted_at"`
}