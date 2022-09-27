package dto

import (
	"time"
)

type BlogDTO struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Author_ID   int64        `json:"author_id"`
	Author_Name string       `json:"author_name"`
	Title       string       `json:"title"`
	Slug        string       `json:"slug" gorm:"unique"`
	Body        string       `json:"body"`
	Comments    []CommentDTO `json:"comments"`
}
