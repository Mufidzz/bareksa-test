package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
)

func (uc *Usecase) CreateNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error) {
	return uc.repositories.CreateBulkNewsTopics(in)
}
func (uc *Usecase) DeleteNewsTopics(newsTopicID []int) (deletedID []int, err error) {
	return uc.repositories.DeleteBulkNewsTopics(newsTopicID)
}
func (uc *Usecase) UpdateNewsTopics(newNewsTopics []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkNewsTopics(newNewsTopics)
}
func (uc *Usecase) GetNewsTopics(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsTopicsResponse, err error) {
	// Get From Redis First
	var redisData []presentation.GetNewsTopicsResponse
	err = uc.repositories.GetObject(ctx.Request.RequestURI, &redisData)
	if err == nil {
		return redisData, nil
	}

	var pagination *presentation.Pagination
	var newsFilter *presentation.NewsTopicFilter

	if paginationString != "" {
		err = urlutils.DecodeEncodedString(paginationString, &pagination)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNewsTopics",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	if filterString != "" {
		err = urlutils.DecodeEncodedString(filterString, &newsFilter)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNewsTopics",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	res, err = uc.repositories.GetBulkNewsTopics(pagination, newsFilter)
	if err != nil {
		return nil, response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetNewsTopics",
			Description:  "Failed Store data to redis",
			Trace:        err,
		}.Error()
	}

	// Save Result to Redis
	err = uc.repositories.SaveObject(ctx.Request.RequestURI, res)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetNewsTopics",
			Description:  "Failed Store data to redis",
			Trace:        err,
		}.Error())
	}

	return res, nil
}
