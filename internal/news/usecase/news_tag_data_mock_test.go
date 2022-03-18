package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type MockNewsTagDataRepository struct {
	createBulkNewsTags createBulkNewsTags
	getBulkNewsTags    getBulkNewsTags
	updateBulkNewsTags updateBulkNewsTags
	deleteBulkNewsTags deleteBulkNewsTags
}

type createBulkNewsTags struct {
	insertedID []int
	err        error
}

type getBulkNewsTags struct {
	res []presentation.GetNewsTagsResponse
	err error
}

type updateBulkNewsTags struct {
	updatedID []int
	err       error
}

type deleteBulkNewsTags struct {
	deletedID []int
	err       error
}

func (mntdr *MockNewsTagDataRepository) CreateBulkNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error) {
	return mntdr.createBulkNewsTags.insertedID, mntdr.createBulkNewsTags.err
}
func (mntdr *MockNewsTagDataRepository) GetBulkNewsTags(pagination *presentation.Pagination, filter *presentation.NewsTagsFilter) (res []presentation.GetNewsTagsResponse, err error) {
	return mntdr.getBulkNewsTags.res, mntdr.getBulkNewsTags.err
}
func (mntdr *MockNewsTagDataRepository) UpdateBulkNewsTags(in []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	return mntdr.updateBulkNewsTags.updatedID, mntdr.updateBulkNewsTags.err
}
func (mntdr *MockNewsTagDataRepository) DeleteBulkNewsTags(newsTagID []int) (deletedID []int, err error) {
	return mntdr.deleteBulkNewsTags.deletedID, mntdr.deleteBulkNewsTags.err
}
