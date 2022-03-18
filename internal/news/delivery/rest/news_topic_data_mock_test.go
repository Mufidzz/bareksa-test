package rest

import "github.com/Mufidzz/bareksa-test/presentation"

type MockNewsTopicDataUC struct {
	createNewsTopics createNewsTopics
	deleteNewsTopics deleteNewsTopics
	updateNewsTopics updateNewsTopics
	getNewsTopics    getNewsTopics
}

type createNewsTopics struct {
	insertedID []int
	err        error
}

type deleteNewsTopics struct {
	deletedID []int
	err       error
}

type updateNewsTopics struct {
	updatedID []int
	err       error
}

type getNewsTopics struct {
	res []presentation.GetNewsTopicsResponse
	err error
}

func (mntduc *MockNewsTopicDataUC) CreateNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error) {
	return mntduc.createNewsTopics.insertedID, mntduc.createNewsTopics.err
}
func (mntduc *MockNewsTopicDataUC) DeleteNewsTopics(newsTopicID []int) (deletedID []int, err error) {
	return mntduc.deleteNewsTopics.deletedID, mntduc.deleteNewsTopics.err
}
func (mntduc *MockNewsTopicDataUC) UpdateNewsTopics(newNewsTopics []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error) {
	return mntduc.updateNewsTopics.updatedID, mntduc.updateNewsTopics.err
}
func (mntduc *MockNewsTopicDataUC) GetNewsTopics(paginationString, filterString string) (res []presentation.GetNewsTopicsResponse, err error) {
	return mntduc.getNewsTopics.res, mntduc.getNewsTopics.err
}
