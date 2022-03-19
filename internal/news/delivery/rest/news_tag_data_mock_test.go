package rest

import (
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
)

type MockNewsTagDataUC struct {
	createNewsTags createNewsTags
	deleteNewsTags deleteNewsTags
	updateNewsTags updateNewsTags
	getNewsTags    getNewsTags
}

type createNewsTags struct {
	insertedID []int
	err        error
}

type deleteNewsTags struct {
	deletedID []int
	err       error
}

type updateNewsTags struct {
	updatedID []int
	err       error
}

type getNewsTags struct {
	res []presentation.GetNewsTagsResponse
	err error
}

func (mntduc *MockNewsTagDataUC) CreateNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error) {
	return mntduc.createNewsTags.insertedID, mntduc.createNewsTags.err
}
func (mntduc *MockNewsTagDataUC) DeleteNewsTags(newsTopicID []int) (deletedID []int, err error) {
	return mntduc.deleteNewsTags.deletedID, mntduc.deleteNewsTags.err
}
func (mntduc *MockNewsTagDataUC) UpdateNewsTags(newNewsTags []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	return mntduc.updateNewsTags.updatedID, mntduc.updateNewsTags.err
}
func (mntduc *MockNewsTagDataUC) GetNewsTags(ctx *gin.Context, paginationString, filterString string) (res []presentation.GetNewsTagsResponse, err error) {
	return mntduc.getNewsTags.res, mntduc.getNewsTags.err
}
