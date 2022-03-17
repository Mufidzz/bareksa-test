package presentation

import (
	"time"
)

type GetNewsResponse struct {
	ID         int64     `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	TopicsName string    `db:"topics_name"`
	TagsName   string    `db:"tags_name"`
}
