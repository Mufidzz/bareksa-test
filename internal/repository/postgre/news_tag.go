package postgre

import "Test_Bareksa/presentation"

func (db *Postgre) CreateBulkNewsTags(in []presentation.CreateBulkTagsRequest) (insertedID []int, err error) {
	return nil, nil
}

func (db *Postgre) GetBulkNewsTags(pagination *presentation.Pagination, filter *presentation.NewsTagsFilter) (res []presentation.GetNewsTagsResponse, err error) {
	return nil, nil
}

func (db *Postgre) UpdateBulkNewsTags(in []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	return nil, nil
}

func (db *Postgre) DeleteBulkNewsTags(newsTopicID []int) (deletedID []int, err error) {
	return nil, nil
}
