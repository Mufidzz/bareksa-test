package rest

import "github.com/Mufidzz/bareksa-test/presentation"

type MockNewsDataUC struct {
	createSingleNews createSingleNews
	updateSingleNews updateSingleNews
	deleteSingleNews deleteSingleNews
	getSingleNews    getSingleNews
	getNews          getNews
}

type createSingleNews struct {
	err error
}

type updateSingleNews struct {
	err error
}

type deleteSingleNews struct {
	err error
}

type getSingleNews struct {
	res presentation.GetNewsResponse
	err error
}

type getNews struct {
	res []presentation.GetNewsResponse
	err error
}

func (mnduc *MockNewsDataUC) CreateSingleNews(newNews presentation.CreateNewsRequest) error {
	return mnduc.createSingleNews.err
}
func (mnduc *MockNewsDataUC) UpdateSingleNews(updatedNews presentation.UpdateNewsRequest) error {
	return mnduc.updateSingleNews.err
}
func (mnduc *MockNewsDataUC) DeleteSingleNews(newsId int) error {
	return mnduc.deleteSingleNews.err
}
func (mnduc *MockNewsDataUC) GetSingleNews(newsId int) (presentation.GetNewsResponse, error) {
	return mnduc.getSingleNews.res, mnduc.getSingleNews.err
}
func (mnduc *MockNewsDataUC) GetNews(paginationString, filterString string) (res []presentation.GetNewsResponse, err error) {
	return mnduc.getNews.res, mnduc.getNews.err
}
