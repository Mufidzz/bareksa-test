package presentation

type CreateNewsTopicsRequest struct {
	Name string `db:"name"`
}

type GetNewsTopicsResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type UpdateNewsTopicsRequest struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
