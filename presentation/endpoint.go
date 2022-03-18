package presentation

type NewsFilter struct {
	Status int    `json:"status"`
	Topics []int  `json:"topics"`
	NewsID int    `json:"news_id"`
	Title  string `json:"title"`
}

type NewsTopicFilter struct {
	Name        string `json:"name"`
	NewsTopicID int    `json:"news_topic_id"`
}

type NewsTagsFilter struct {
	Name      string `json:"name"`
	NewsTagID int    `json:"news_tag_id"`
}

type Pagination struct {
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
}
