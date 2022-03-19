package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
)

func (uc *Usecase) CreateSingleNews(newNews presentation.CreateNewsRequest) error {
	_, err := uc.repositories.CreateBulkNews([]presentation.CreateNewsRequest{newNews})
	if err != nil {
		return err
	}

	err = uc.repositories.FlushAll()
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "CreateSingleNews",
			Description:  "Failed Flush data redis",
			Trace:        err,
		}.Error())
	}

	return nil
}

func (uc *Usecase) UpdateSingleNews(updatedNews presentation.UpdateNewsRequest) error {
	_, err := uc.repositories.UpdateBulkNews([]presentation.UpdateNewsRequest{updatedNews})
	if err != nil {
		return err
	}

	err = uc.repositories.FlushAll()
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "CreateSingleNews",
			Description:  "Failed Flush data redis",
			Trace:        err,
		}.Error())
	}

	return nil
}

func (uc *Usecase) DeleteSingleNews(newsId int) error {
	_, err := uc.repositories.DeleteBulkNews([]int{newsId})
	if err != nil {
		return err
	}

	err = uc.repositories.CleanNewsTopicsAssoc([]int{newsId})
	if err != nil {
		return err
	}

	err = uc.repositories.CleanNewsTagAssoc([]int{newsId})
	if err != nil {
		return err
	}

	err = uc.repositories.FlushAll()
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "CreateSingleNews",
			Description:  "Failed Flush data redis",
			Trace:        err,
		}.Error())
	}

	return nil
}

func (uc *Usecase) AssignNewsWithNewsTopic(in presentation.CreateNewsTopicsAssoc) error {
	err := uc.repositories.CreateBulkNewsTopicsAssoc([]presentation.CreateNewsTopicsAssoc{in})
	if err != nil {
		return err
	}

	err = uc.repositories.FlushAll()
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "AssignNewsWithNewsTopic",
			Description:  "Failed Flush data redis",
			Trace:        err,
		}.Error())
	}
	return nil
}

func (uc *Usecase) AssignNewsWithNewsTag(in presentation.CreateNewsTagsAssoc) error {
	err := uc.repositories.CreateBulkNewsTagsAssoc([]presentation.CreateNewsTagsAssoc{in})
	if err != nil {
		return err
	}

	err = uc.repositories.FlushAll()
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "CreateSingleNews",
			Description:  "Failed Flush data redis",
			Trace:        err,
		}.Error())
	}
	return nil
}

func (uc *Usecase) GetSingleNews(ctx *gin.Context, newsId int) (presentation.GetNewsResponse, error) {
	// Get From Redis First
	var redisData presentation.GetNewsResponse
	err := uc.repositories.GetObject(ctx.Request.RequestURI, &redisData)
	if err == nil {
		return redisData, nil
	}

	// Get From Database
	news, err := uc.repositories.GetBulkNews(presentation.Pagination{
		Offset: 0,
		Count:  1,
	}, &presentation.NewsFilter{NewsID: newsId})
	if err != nil {
		return presentation.GetNewsResponse{}, response.InternalError{
			Type:         "Usecase",
			Name:         "News Data",
			FunctionName: "GetSingleNews",
			Description:  "Failed running repository",
			Trace:        err,
		}.Error()
	}

	// Save Result to Redis
	err = uc.repositories.SaveObject(ctx.Request.RequestURI, news[0])
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetSingleNews",
			Description:  "Failed Store data to redis",
			Trace:        err,
		}.Error())
	}

	return news[0], err
}

func (uc *Usecase) GetNews(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsResponse, err error) {
	var pagination presentation.Pagination
	var newsFilter *presentation.NewsFilter

	// Get From Redis First
	var redisData []presentation.GetNewsResponse
	err = uc.repositories.GetObject(ctx.Request.RequestURI, &redisData)
	if err == nil {
		return redisData, nil
	}

	// Get From Database
	err = urlutils.DecodeEncodedString(paginationString, &pagination)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Usecase",
			Name:         "News Data",
			FunctionName: "GetNews",
			Description:  "Failed running repository",
			Trace:        err,
		}.Error()
	}

	if filterString != "" {
		err = urlutils.DecodeEncodedString(filterString, &newsFilter)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNews",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	res, err = uc.repositories.GetBulkNews(pagination, newsFilter)
	if err != nil {
		return res, response.InternalError{
			Type:         "UC",
			Name:         "News Data",
			FunctionName: "GetNews",
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
			FunctionName: "GetNews",
			Description:  "Failed Store data to redis",
			Trace:        err,
		}.Error())
	}

	return res, nil
}
