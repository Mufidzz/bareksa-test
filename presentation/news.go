package presentation

import (
	"time"
)

// Constant for News Status
const NEWS_STATUS_DRAFT = 1
const NEWS_STATUS_PUBLISHED = 2
const NEWS_STATUS_DELETED = 3

type GetNewsResponse struct {
	ID         int64     `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Title      string    `db:"title" json:"title"`
	Content    string    `db:"content" json:"content"`
	TopicsName string    `db:"topics_name" json:"topics_name"`
	TagsName   string    `db:"tags_name" json:"tags_name"`
	Status     int       `db:"status" json:"status"`
}

type CreateNewsRequest struct {
	Title   string `db:"title"`
	Content string `db:"content"`
	Status  int    `db:"status"`
}

type UpdateNewsRequest struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Status  int    `db:"status"`
}
