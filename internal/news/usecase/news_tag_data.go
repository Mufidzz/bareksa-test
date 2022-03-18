package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
)

func (uc *Usecase) CreateNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error) {
	return uc.repositories.CreateBulkNewsTags(in)
}
func (uc *Usecase) DeleteNewsTags(newsTopicID []int) (deletedID []int, err error) {
	return uc.repositories.DeleteBulkNewsTags(newsTopicID)
}
func (uc *Usecase) UpdateNewsTags(newNewsTags []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkNewsTags(newNewsTags)
}
func (uc *Usecase) GetNewsTags(paginationString, filterString string) (res []presentation.GetNewsTagsResponse, err error) {
	var pagination *presentation.Pagination
	var filter *presentation.NewsTagsFilter

	if paginationString != "" {
		err = urlutils.DecodeEncodedString(paginationString, &pagination)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNews",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	if filterString != "" {
		err = urlutils.DecodeEncodedString(filterString, &filter)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Usecase",
				Name:         "News Data",
				FunctionName: "GetNews",
				Description:  "Failed running repository",
				Trace:        err,
			}.Error()
		}
	}

	return uc.repositories.GetBulkNewsTags(pagination, filter)
}
