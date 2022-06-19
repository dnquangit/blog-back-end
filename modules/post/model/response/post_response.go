package postresponsemodel

import "time"

type PostResponse struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
	CreateAt    time.Time  `json:"created_at"`
	UpdateAt    time.Time  `json:"updated_at"`
}
