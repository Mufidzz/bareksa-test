package postgre

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/dbutils"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/lib/pq"
)

func (db *Postgre) CreateBulkNewsTopics(in []presentation.CreateNewsTopicsRequest) (insertedID []int, err error) {
	q := `INSERT INTO news_topics (name) VALUES`

	queryParamLen := 1

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		q = fmt.Sprintf("%s ($%d),", q, paramCount)
		paramArgs = append(paramArgs, v.Name)
		paramCount += queryParamLen
	}

	// Remove Comma From end of line and Fetch ID after creation
	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	rows, err := db.newsDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateBulkNewsTopics",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "CreateBulkNewsTopics",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		insertedID = append(insertedID, id)
	}

	return insertedID, nil
}

func (db *Postgre) GetBulkNewsTopics(pagination *presentation.Pagination, filter *presentation.NewsTopicFilter) (res []presentation.GetNewsTopicsResponse, err error) {
	q := `SELECT id, name FROM news_topics `

	paramCount := 0
	paramArgs := []interface{}{}

	// Apply Filter if Available
	if filter != nil {
		if filter.NewsTopicID != 0 {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "id", dbutils.COMPARATOR_EQUAL, paramCount)
			paramArgs = append(paramArgs, filter.NewsTopicID)
		}

		if filter.Name != "" {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "name", dbutils.COMPARATOR_LIKE, paramCount)
			paramArgs = append(paramArgs, fmt.Sprintf("%%%s%%", filter.Name))
		}
	}

	// Implement Pagination if Any
	if pagination != nil {
		q = fmt.Sprintf("%s LIMIT $%d OFFSET $%d", q, paramCount+1, paramCount+2)
		paramArgs = append(paramArgs, pagination.Count, pagination.Offset)
	}

	rows, err := db.newsDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetBulkNewsTopics",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetNewsTopicsResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetBulkNewsTopics",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) UpdateBulkNewsTopics(in []presentation.UpdateNewsTopicsRequest) (updatedID []int, err error) {
	q := `UPDATE news_topics SET name = new_values.name FROM (VALUES %s) as new_values (id, name) WHERE news_topics.id = new_values.id RETURNING news_topics.id`

	queryParamLen := 2

	queryValues := ""
	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		queryValues = fmt.Sprintf("%s($%d::BIGINT, $%d::TEXT),", queryValues, paramCount, paramCount+1)
		paramArgs = append(paramArgs, v.ID, v.Name)
		paramCount += queryParamLen
	}

	q = fmt.Sprintf(q, queryValues[:len(queryValues)-1])

	rows, err := db.newsDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "UpdateBulkNewsTopics",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "UpdateBulkNewsTopics",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		updatedID = append(updatedID, id)
	}

	return updatedID, nil
}

func (db *Postgre) DeleteBulkNewsTopics(newsTopicID []int) (deletedID []int, err error) {
	q := `DELETE FROM news_topics WHERE id = ANY($1) RETURNING id`

	rows, err := db.newsDatabase.Master.Queryx(q, pq.Array(newsTopicID))
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "DeleteBulkNewsTopics",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "DeleteBulkNewsTopics",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		deletedID = append(deletedID, id)
	}

	return deletedID, nil
}
