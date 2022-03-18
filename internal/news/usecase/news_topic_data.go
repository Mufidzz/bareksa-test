package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
)

func (uc *Usecase) CreateNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error) {
	return uc.repositories.CreateBulkNewsTopics(in)
}
func (uc *Usecase) DeleteNewsTopics(newsTopicID []int) (deletedID []int, err error) {
	return uc.repositories.DeleteBulkNewsTopics(newsTopicID)
}
func (uc *Usecase) UpdateNewsTopics(newNewsTopics []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkNewsTopics(newNewsTopics)
}
func (uc *Usecase) GetNewsTopics(paginationString, filterString string) (res []presentation.GetNewsTopicsResponse, err error) {
	var pagination *presentation.Pagination
	var newsFilter *presentation.NewsTopicFilter

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
		err = urlutils.DecodeEncodedString(filterString, &newsFilter)
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

	return uc.repositories.GetBulkNewsTopics(pagination, newsFilter)
}
