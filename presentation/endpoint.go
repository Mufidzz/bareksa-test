package presentation

type NewsFilter struct {
	Status int   `json:"status"`
	Topics []int `json:"topics"`
	NewsID int   `json:"news_id"`
}

type Pagination struct {
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
}
