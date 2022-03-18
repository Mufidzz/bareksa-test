package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type NewsDataRepository interface {
	CreateBulkNews(in []presentation.CreateNewsRequest) (insertedID []int, err error)
	GetBulkNews(pagination presentation.Pagination, filter *presentation.NewsFilter) (res []presentation.GetNewsResponse, err error)
	UpdateBulkNews(in []presentation.UpdateNewsRequest) (updatedID []int, err error)
	DeleteBulkNews(newsID []int) (deletedID []int, err error)
}

type NewsTopicDataRepository interface {
	CreateBulkNewsTopics(in []presentation.CreateBulkTopicsRequest) (insertedID []int, err error)
	GetBulkNewsTopics(pagination *presentation.Pagination, filter *presentation.NewsTopicFilter) (res []presentation.GetNewsTopicsResponse, err error)
	UpdateBulkNewsTopics(in []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error)
	DeleteBulkNewsTopics(newsTopicID []int) (deletedID []int, err error)
}
