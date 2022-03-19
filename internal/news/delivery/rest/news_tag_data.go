package rest

import (
	"github.com/Mufidzz/bareksa-test/pkg/logger"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handler *HTTPHandler) HandleGetNewsTag(ctx *gin.Context) {
	filterString, paginationString := ctx.Query("filter"), ctx.Query("pagination")

	news, err := handler.usecases.GetNewsTags(ctx, paginationString, filterString)
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

func (handler *HTTPHandler) HandleUpdateNewsTag(ctx *gin.Context) {
	var newNewsTag []presentation.UpdateNewsTagsRequest

	err := ctx.BindJSON(&newNewsTag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    newNewsTag,
		})
		return
	}

	updatedIDs, err := handler.usecases.UpdateNewsTags(newNewsTag)
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
			Message: "Failed Update News Tags",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Updated News Tags",
		Data: gin.H{
			"updated_id": updatedIDs,
		},
	})
}

func (handler *HTTPHandler) HandleCreateNewsTag(ctx *gin.Context) {
	var newNewsTag []presentation.CreateNewsTagsRequest

	err := ctx.BindJSON(&newNewsTag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Binding JSON",
			Type:    0,
			Data:    nil,
		})
		return
	}

	_, err = handler.usecases.CreateNewsTags(newNewsTag)
	if err != nil {
		logger.Error(response.InternalError{
			Type:         "Handler",
			Name:         "News Data",
			FunctionName: "HandleCreateNewsTag",
			Description:  "error running usecase",
			Trace:        err,
		}.Error())

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Failed Run Create News Tags",
			Type:    0,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleDeleteNewsTag(ctx *gin.Context) {
	topicsIds := ctx.QueryArray("id")
	var intTagIds []int

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
				Message: "Failed Parsing Tag ID, Please check all Tag ID is valid Number",
				Type:    0,
				Data:    topicsIds,
			})
			return
		}

		intTagIds = append(intTagIds, _t)
	}

	deletedIds, err := handler.usecases.DeleteNewsTags(intTagIds)

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
			Message: "Failed Run Delete News Tag",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success Delete News Tags",
		Data: gin.H{
			"deleted_id": deletedIds,
		},
	})
}
