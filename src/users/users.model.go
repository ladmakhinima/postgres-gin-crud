package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	FirstName string         `json:"firstname" gorm:"type: varchar(100);not null;" binding:"required"`
	LastName  string         `json:"lastname" gorm:"type: varchar(100);not null;" binding:"required"`
	Age       int            `json:"age" gorm:"type: int;not null;" binding:"required"`
	Password  string         `json:"password" gorm:"type: text;not null;sql:-" binding:"required"`
	Email     string         `json:"email" gorm:"type: varchar(100);unique;not null;" binding:"required,email"`
	ID        uint           `json:"id" gorm:"autoIncrement;primaryKey;unique"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
