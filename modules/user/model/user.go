package usermodel

import "time"

type User struct {
	Id        string    `json:"id" gorm:"column:id;"`
	UserName  string    `json:"username" gorm:"column:username;"`
	Password  string    `json:"password" gorm:"column:password;"`
	FirstName string    `json:"first_name" gorm:"column:first_name;"`
	LastName  string    `json:"last_name" gorm:"column:last_name;"`
	Email     string    `json:"email" gorm:"column:email;"`
	CreateAt  time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt  time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Deleted   bool      `json:"deleted" gorm:"column:deleted;"`
}

func (User) TableName() string { return "users" }
