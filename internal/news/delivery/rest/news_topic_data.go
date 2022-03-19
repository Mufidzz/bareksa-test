package rest

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handler *HTTPHandler) HandleGetNewsTopic(ctx *gin.Context) {
	filterString, paginationString := ctx.Query("filter"), ctx.Query("pagination")

	news, err := handler.usecases.GetNewsTopics(ctx, paginationString, filterString)
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

func (handler *HTTPHandler) HandleUpdateNewsTopic(ctx *gin.Context) {
	var newNewsTopic []presentation.UpdateNewsTopicsRequest

	err := ctx.BindJSON(&newNewsTopic)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    newNewsTopic,
		})
		return
	}

	updatedIDs, err := handler.usecases.UpdateNewsTopics(newNewsTopic)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleUpdateSingleNews",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Failed Update News Topics",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Updated News Topics",
		Data: gin.H{
			"updated_id": updatedIDs,
		},
	})
}

func (handler *HTTPHandler) HandleCreateNewsTopic(ctx *gin.Context) {
	var newNewsTopic []presentation.CreateNewsTopicsRequest

	err := ctx.BindJSON(&newNewsTopic)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    nil,
		})
		return
	}

	_, err = handler.usecases.CreateNewsTopics(newNewsTopic)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleCreateNewsTopic",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Create News Topics",
			Type:    0,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleDeleteNewsTopic(ctx *gin.Context) {
	topicsIds := ctx.QueryArray("id")
	var intTopicIds []int

	if len(topicsIds) <= 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "ID Must more than 1",
			Type:    0,
			Data:    nil,
		})
		return
	}

	for _, v := range topicsIds {
		_t, err := strconv.Atoi(v)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Failed Parsing Topic ID, Please check all Topic ID is valid Number",
				Type:    0,
				Data:    topicsIds,
			})
			return
		}

		intTopicIds = append(intTopicIds, _t)
	}

	deletedIds, err := handler.usecases.DeleteNewsTopics(intTopicIds)

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
			Message: "Failed Run Delete News Topic",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Delete News Topics",
		Data: gin.H{
			"deleted_id": deletedIds,
		},
	})
}
