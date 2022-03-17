package postgre

import (
	"Test_Bareksa/pkg/dbutils"
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
	"fmt"
	"github.com/lib/pq"
)

func (db *Postgre) GetNews(pagination presentation.Pagination, filter *presentation.NewsFilter) (res []presentation.GetNewsResponse, err error) {
	q := `SELECT news.id, news.created_at, news.updated_at, news.title, news.content, string_agg(DISTINCT topics.name, ', ') as topics_name, string_agg(DISTINCT tags.name, ', ') as tags_name FROM news
			INNER JOIN assoc_news_topics aTopics on news.id = aTopics.news_id
            LEFT JOIN news_topics topics on aTopics.news_topic_id = topics.id
            INNER JOIN assoc_news_tags aTags on news.id = aTags.news_id
            LEFT JOIN news_tags tags on aTags.news_tag_id = tags.id`

	paramCount := 0
	paramArgs := []interface{}{}

	// Apply Filter if Available
	if filter != nil {
		if filter.NewsID != 0 {
			paramCount += 1
			q = dbutils.AddCustomFilter(q, dbutils.CONNECTOR_AND, "news.id", dbutils.COMPARATOR_EQUAL, fmt.Sprintf("ANY($%d)", paramCount))
			paramArgs = append(paramArgs, filter.Status)
		}

		if filter.Status != 0 {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "status", dbutils.COMPARATOR_EQUAL, paramCount)
			paramArgs = append(paramArgs, filter.Status)
		}

		if filter.Topics != nil && len(filter.Topics) > 0 {
			paramCount += 1
			q = dbutils.AddCustomFilter(q, dbutils.CONNECTOR_AND, "topics.id", dbutils.COMPARATOR_EQUAL, fmt.Sprintf("ANY($%d)", paramCount))
			paramArgs = append(paramArgs, pq.Array(filter.Topics))
		}
	}

	// Implement Grouping
	q = fmt.Sprintf("%s GROUP BY %s", q, "news.id")

	// Implement Pagination
	q = fmt.Sprintf("%s LIMIT $%d OFFSET $%d", q, paramCount+1, paramCount+2)
	paramArgs = append(paramArgs, pagination.Count, pagination.Offset)

	rows, err := db.newsDatabase.Master.Queryx(q, paramArgs...)
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
