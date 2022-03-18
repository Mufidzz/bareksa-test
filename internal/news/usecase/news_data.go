package usecase

import (
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
)

func (uc *Usecase) CreateSingleNews(newNews presentation.CreateNewsRequest) error {
	_, err := uc.repositories.CreateBulkNews([]presentation.CreateNewsRequest{newNews})
	return err
}

func (uc *Usecase) UpdateSingleNews(updatedNews presentation.UpdateNewsRequest) error {
	_, err := uc.repositories.UpdateBulkNews([]presentation.UpdateNewsRequest{updatedNews})
	return err
}

func (uc *Usecase) DeleteSingleNews(newsId int) error {
	_, err := uc.repositories.DeleteBulkNews([]int{newsId})
	return err
}

func (uc *Usecase) GetSingleNews(newsId int) (presentation.GetNewsResponse, error) {
	news, err := uc.repositories.GetBulkNews(presentation.Pagination{
		Offset: 0,
		Count:  1,
	}, &presentation.NewsFilter{NewsID: newsId})

	if err != nil {
		return presentation.GetNewsResponse{}, response.InternalError{
			Type:         "Usecase",
			Name:         "News Data",
			FunctionName: "GetSingleNews",
			Description:  "Failed running repository",
			Trace:        err,
		}.Error()
	}

	return news[0], err
}

func (uc *Usecase) GetNews(paginationString, filterString string) (res []presentation.GetNewsResponse, err error) {
	var pagination presentation.Pagination
	var newsFilter *presentation.NewsFilter

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

	return uc.repositories.GetBulkNews(pagination, newsFilter)
}

func (uc *Usecase) AssignNewsWithNewsTopic(newsId int, topicIDs []int) {

}
