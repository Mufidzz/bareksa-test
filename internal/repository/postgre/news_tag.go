package postgre

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/dbutils"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/lib/pq"
)

func (db *Postgre) CreateBulkNewsTags(in []presentation.CreateBulkTagsRequest) (insertedID []int, err error) {
	q := `INSERT INTO news_tags (name) VALUES`

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
			FunctionName: "CreateBulkNewsTags",
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
				FunctionName: "CreateBulkNewsTags",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		insertedID = append(insertedID, id)
	}

	return insertedID, nil
}

func (db *Postgre) GetBulkNewsTags(pagination *presentation.Pagination, filter *presentation.NewsTagsFilter) (res []presentation.GetNewsTagsResponse, err error) {
	q := `SELECT id, name FROM news_tags `

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
			q = dbutils.AddCustomFilter(q, dbutils.CONNECTOR_AND, "name", dbutils.COMPARATOR_LIKE, fmt.Sprintf("%%$%d%%", paramCount))
			paramArgs = append(paramArgs, filter.Name)
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
			FunctionName: "GetBulkNewsTags",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetNewsTagsResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetBulkNewsTags",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) UpdateBulkNewsTags(in []presentation.UpdateNewsTagsRequest) (updatedID []int, err error) {
	q := `UPDATE news_tags SET name = new_values.name FROM (VALUES %s) as new_values (id, name) WHERE news_tags.id = new_values.id RETURNING news_tags.id`

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
			FunctionName: "UpdateBulkNewsTags",
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
				FunctionName: "UpdateBulkNewsTags",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		updatedID = append(updatedID, id)
	}

	return updatedID, nil
}

func (db *Postgre) DeleteBulkNewsTags(newsTopicID []int) (deletedID []int, err error) {
	q := `DELETE FROM news_tags WHERE id = ANY($1) RETURNING id`

	rows, err := db.newsDatabase.Master.Queryx(q, pq.Array(newsTopicID))
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "DeleteBulkNewsTags",
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
				FunctionName: "DeleteBulkNewsTags",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		deletedID = append(deletedID, id)
	}

	return deletedID, nil
}
