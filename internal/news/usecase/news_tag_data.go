package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
)

func (uc *Usecase) CreateNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error) {
	return uc.repositories.CreateBulkNewsTags(in)
}
func (uc *Usecase) DeleteNewsTags(newsTopicID []int) (deletedID []int, err error) {
	return uc.repositories.DeleteBulkNewsTags(newsTopicID)
}
func (uc *Usecase) UpdateNewsTags(newNewsTags []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkNewsTags(newNewsTags)
}
func (uc *Usecase) GetNewsTags(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsTagsResponse, err error) {
	// Get From Redis First
	var redisData []presentation.GetNewsTagsResponse
	err = uc.repositories.GetObject(ctx.Request.RequestURI, &redisData)
	if err == nil {
		return redisData, nil
	}

	// Get From Database
	var pagination *presentation.Pagination
	var filter *presentation.NewsTagsFilter

	if paginationString != "" {
		err = urlutils.DecodeEncodedString(paginationString, &pagination)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNewsTags",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	if filterString != "" {
		err = urlutils.DecodeEncodedString(filterString, &filter)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNewsTags",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	res, err = uc.repositories.GetBulkNewsTags(pagination, filter)
	if err != nil {
		return nil, response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetNewsTags",
			Description:  "Failed run Repository",
			Trace:        err,
		}.Error()
	}

	// Save Result to Redis
	err = uc.repositories.SaveObject(ctx.Request.RequestURI, res)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetNewsTags",
			Description:  "Failed Store data to redis",
			Trace:        err,
		}.Error())
	}

	return res, nil
}
