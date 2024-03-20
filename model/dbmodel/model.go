package dbmodel

import "gorm.io/gorm"

type RegisterUser struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
	Email    string `json:"email" gorm:"column:email"`
	Name     string `json:"name" gorm:"column:name"`
	Role     string `json:"role" gorm:"column:role"`
}

func (RegisterUser) TableName() string {
	return "register_user"
}

type LoginUser struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;unique;not null"`
}
