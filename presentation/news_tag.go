package presentation

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
