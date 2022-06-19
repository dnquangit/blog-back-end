package usermodel

import "time"

type UserRole struct {
	UserId   string    `json:"user_id" gorm:"column:user_id;"`
	RoleId   string    `json:"role_id" gorm:"column:role_id;"`
	CreateAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Deleted  bool      `json:"deleted" gorm:"column:deleted;"`
}

func (UserRole) TableName() string { return "user_roles" }

type Role struct {
	Id       string    `json:"id" gorm:"column:id;"`
	Name     string    `json:"name" gorm:"column:name;"`
	CreateAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Deleted  bool      `json:"deleted" gorm:"column:deleted;"`
}

func (Role) TableName() string { return "roles" }
