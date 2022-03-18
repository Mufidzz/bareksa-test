package rest

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handler *HTTPHandler) HandleGetNews(ctx *gin.Context) {
	filterString, paginationString := ctx.Query("filter"), ctx.Query("pagination")

	if paginationString == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Pagination Cannot be Blank, use URL Query to assign",
			Type:    0,
			Data:    nil,
		})
		return
	}

	news, err := handler.usecases.GetNews(paginationString, filterString)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleGetNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Get News",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Getting News",
		Data:    news,
	})
}

func (handler *HTTPHandler) HandleGetSingleNews(ctx *gin.Context) {
	newsID := ctx.Param("id")
	if newsID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "News ID Cannot be Blank, use URL Parameter to assign",
			Type:    0,
			Data:    nil,
		})
		return
	}

	intNewsID, err := strconv.Atoi(newsID)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleGetSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Parsing News ID, Please check news id is valid Number",
			Type:    0,
			Data:    nil,
		})
		return
	}

	news, err := handler.usecases.GetSingleNews(intNewsID)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleGetSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Get Single News",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Getting News",
		Data:    news,
	})
}

func (handler *HTTPHandler) HandleUpdateSingleNews(ctx *gin.Context) {
	newsID := ctx.Param("id")
	if newsID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "News ID Cannot be Blank, use URL Parameter to assign",
			Type:    0,
			Data:    nil,
		})
		return
	}

	intNewsID, err := strconv.Atoi(newsID)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleUpdateSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Parsing News ID, Please check news id is valid Number",
			Type:    0,
			Data:    nil,
		})
		return
	}

	var newNews presentation.UpdateNewsRequest

	err = ctx.BindJSON(&newNews)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    newNews,
		})
		return
	}

	newNews.ID = intNewsID
	err = handler.usecases.UpdateSingleNews(newNews)

	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleUpdateSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Get Single News",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleCreateSingleNews(ctx *gin.Context) {
	var newNews presentation.CreateNewsRequest

	err := ctx.BindJSON(&newNews)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    newNews,
		})
		return
	}

	err = handler.usecases.CreateSingleNews(newNews)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleCreateSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Get Single News",
			Type:    0,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleDeleteSingleNews(ctx *gin.Context) {
	newsID := ctx.Param("id")
	if newsID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "News ID Cannot be Blank, use URL Parameter to assign",
			Type:    0,
			Data:    nil,
		})
		return
	}

	intNewsID, err := strconv.Atoi(newsID)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleDeleteSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Parsing News ID, Please check news id is valid Number",
			Type:    0,
			Data:    nil,
		})
		return
	}

	err = handler.usecases.DeleteSingleNews(intNewsID)

	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleDeleteSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Get Single News",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}