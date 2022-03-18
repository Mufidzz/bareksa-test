package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type MockNewsTopicDataRepository struct {
	createBulkNewsTopics createBulkNewsTopics
	getBulkNewsTopics    getBulkNewsTopics
	updateBulkNewsTopics updateBulkNewsTopics
	deleteBulkNewsTopics deleteBulkNewsTopics
}

type createBulkNewsTopics struct {
	insertedID []int
	err        error
}

type getBulkNewsTopics struct {
	res []presentation.GetNewsTopicsResponse
	err error
}

type updateBulkNewsTopics struct {
	updatedID []int
	err       error
}

type deleteBulkNewsTopics struct {
	deletedID []int
	err       error
}

func (mntdr *MockNewsTopicDataRepository) CreateBulkNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error) {
	return mntdr.createBulkNewsTopics.insertedID, mntdr.createBulkNewsTopics.err
}
func (mntdr *MockNewsTopicDataRepository) GetBulkNewsTopics(pagination *presentation.Pagination, filter *presentation.NewsTopicFilter) (res []presentation.GetNewsTopicsResponse, err error) {
	return mntdr.getBulkNewsTopics.res, mntdr.getBulkNewsTopics.err
}
func (mntdr *MockNewsTopicDataRepository) UpdateBulkNewsTopics(in []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error) {
	return mntdr.updateBulkNewsTopics.updatedID, mntdr.updateBulkNewsTopics.err
}
func (mntdr *MockNewsTopicDataRepository) DeleteBulkNewsTopics(newsTopicID []int) (deletedID []int, err error) {
	return mntdr.deleteBulkNewsTopics.deletedID, mntdr.deleteBulkNewsTopics.err
}
