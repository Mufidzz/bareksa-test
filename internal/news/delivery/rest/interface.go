package rest

import "github.com/Mufidzz/bareksa-test/presentation"

type NewsDataUC interface {
	CreateSingleNews(newNews presentation.CreateNewsRequest) error
	UpdateSingleNews(updatedNews presentation.UpdateNewsRequest) error
	DeleteSingleNews(newsId int) error
	GetSingleNews(newsId int) (presentation.GetNewsResponse, error)
	GetNews(paginationString, filterString string) (res []presentation.GetNewsResponse, err error)
}
