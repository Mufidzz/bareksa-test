package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type MockNewsRepository struct {
	createBulkNews createBulkNews
	getBulkNews    getBulkNews
	updateBulkNews updateBulkNews
	deleteBulkNews deleteBulkNews
}

type createBulkNews struct {
	insertedID []int
	err        error
}

type getBulkNews struct {
	res []presentation.GetNewsResponse
	err error
}

type updateBulkNews struct {
	updatedID []int
	err       error
}

type deleteBulkNews struct {
	deletedID []int
	err       error
}

func (mnr *MockNewsRepository) CreateBulkNews(in []presentation.CreateNewsRequest) (insertedID []int, err error) {
	return mnr.createBulkNews.insertedID, mnr.createBulkNews.err
}

func (mnr *MockNewsRepository) GetBulkNews(pagination presentation.Pagination, filter *presentation.NewsFilter) (res []presentation.GetNewsResponse, err error) {
	return mnr.getBulkNews.res, mnr.getBulkNews.err
}

func (mnr *MockNewsRepository) UpdateBulkNews(in []presentation.UpdateNewsRequest) (updatedID []int, err error) {
	return mnr.updateBulkNews.updatedID, mnr.updateBulkNews.err
}

func (mnr *MockNewsRepository) DeleteBulkNews(newsID []int) (deletedID []int, err error) {
	return mnr.deleteBulkNews.deletedID, mnr.deleteBulkNews.err
}
