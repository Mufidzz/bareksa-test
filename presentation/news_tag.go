package presentation

type CreateNewsTagsAssoc struct {
	NewsID    int   `json:"news_id"`
	NewsTagID []int `json:"news_tag_id"`
}
type CreateNewsTagsRequest struct {
	Name string `db:"name"`
}

type GetNewsTagsResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type UpdateNewsTagsRequest struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
