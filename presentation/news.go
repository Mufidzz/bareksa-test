package presentation

import (
	"time"
)

// Constant for News Status
const NEWS_STATUS_DRAFT = 1
const NEWS_STATUS_PUBLISHED = 2
const NEWS_STATUS_DELETED = 3

type GetNewsResponse struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	TopicsName string    `db:"topics_name"`
	TagsName   string    `db:"tags_name"`
	Status     int       `db:"status"`
}

type CreateBulkNewsRequest struct {
	Title   string `db:"title"`
	Content string `db:"content"`
	Status  int    `db:"status"`
}
