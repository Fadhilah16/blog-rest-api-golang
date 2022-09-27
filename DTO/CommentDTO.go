package dto

import "time"

type CommentDTO struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Author_ID       int64     `json:"author_id"`
	Author_Username string    `json:"author_username"`
	Author_Name     string    `json:"author_name"`
	Blog_ID         int64     `json:"blog_id"`
	Blog_Slug       string    `json:"blog_slug"`
	Body            string    `json:"body"`
}
