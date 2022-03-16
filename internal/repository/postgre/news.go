package postgre

import (
	"Test_Bareksa/pkg/dbutils"
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
)

func (db *Postgre) GetNews(filter *presentation.NewsFilter) (res []presentation.GetNewsResponse, err error) {
	q := `SELECT id, created_at, updated_at, title, content 
			FROM news`

	paramCount := 0
	paramArgs := []interface{}{}

	if filter != nil {
		if filter.Status != 0 {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "status", paramCount)
			paramArgs = append(paramArgs, filter.Status)
		}

		if filter.Topic != 0 {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "news_topic_id", paramCount)
			paramArgs = append(paramArgs, filter.Topic)
		}
	}

	rows, err := db.newsDatabase.Master.Queryx(q, paramArgs)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetNews",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetNewsResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetNews",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}
