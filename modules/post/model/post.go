package postmodel

import (
	"time"
)

type Post struct {
	Id          string     `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Content     string     `json:"content" gorm:"column:content;"`
	Published   bool       `json:"published" gorm:"type:boolean;column:published;"`
	PublishedAt *time.Time `json:"published_at" gorm:"column:published_at;"`
	CreateAt    time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdateAt    time.Time  `json:"updated_at" gorm:"column:updated_at;"`
	Deleted     bool       `json:"deleted" gorm:"column:deleted;"`
}

func (Post) TableName() string { return "posts" }
