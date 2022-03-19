package rest

import (
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
)

type NewsDataUC interface {
	CreateSingleNews(newNews presentation.CreateNewsRequest) error
	UpdateSingleNews(updatedNews presentation.UpdateNewsRequest) error
	DeleteSingleNews(newsId int) error
	GetSingleNews(ctx *gin.Context, newsId int) (presentation.GetNewsResponse, error)
	GetNews(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsResponse, err error)

	AssignNewsWithNewsTopic(in presentation.CreateNewsTopicsAssoc) error
	AssignNewsWithNewsTag(in presentation.CreateNewsTagsAssoc) error
}

type NewsTopicDataUC interface {
	CreateNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error)
	DeleteNewsTopics(newsTopicID []int) (deletedID []int, err error)
	UpdateNewsTopics(newNewsTopics []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error)
	GetNewsTopics(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsTopicsResponse, err error)
}

type NewsTagDataUC interface {
	CreateNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error)
	DeleteNewsTags(newsTopicID []int) (deletedID []int, err error)
	UpdateNewsTags(newNewsTags []presentation.UpdateNewsTagsRequest) (updatedID []int, err error)
	GetNewsTags(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsTagsResponse, err error)
}
