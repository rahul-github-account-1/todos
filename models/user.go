package models
import ("gorm.io/gorm"
"time")

type User struct {
    gorm.Model
    Email      string     `json:"email" gorm:"unique;not null"`
    Password   string     `json:"password" gorm:"not null"`
    DeletedAt  *time.Time `json:"deleted_at"` 
	Todos      []Todo     `json:"todos"`      
}
