package postgre

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/dbutils"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/lib/pq"
)

func (db *Postgre) CreateBulkNewsTags(in []presentation.CreateNewsTagsRequest) (insertedID []int, err error) {
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
		if filter.NewsTagID != 0 {
			paramCount += 1
			q = dbutils.AddFilter(q, dbutils.CONNECTOR_AND, "id", dbutils.COMPARATOR_EQUAL, paramCount)
			paramArgs = append(paramArgs, filter.NewsTagID)
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

func (db *Postgre) CreateBulkNewsTagsAssoc(in []presentation.CreateNewsTagsAssoc) (err error) {
	q := `INSERT INTO assoc_news_tags (news_id, news_tag_id) VALUES`

	queryParamLen := 2

	paramCount := 1
	paramArgs := []interface{}{}

	dataCount := 0

	for _, v := range in {
		dataCount += len(v.NewsTagID)
		for _, newsTagID := range v.NewsTagID {
			q = fmt.Sprintf("%s ($%d, $%d),", q, paramCount, paramCount+1)
			paramArgs = append(paramArgs, v.NewsID, newsTagID)
			paramCount += queryParamLen
		}
	}

	res, err := db.newsDatabase.Master.Exec(q[:len(q)-1], paramArgs...)
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateNewsTagsAssoc",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateNewsTagsAssoc",
			Description:  "failed get number of rows affected",
			Trace:        err,
		}.Error()
	}

	if rowsAffected != int64(dataCount) {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateNewsTagsAssoc",
			Description:  "not all data inserted",
			Trace:        fmt.Errorf("affected rows : %v, data count : %v", rowsAffected, dataCount),
		}.Error()
	}

	return nil
}

func (db *Postgre) CleanNewsTagAssoc(newsID []int) (err error) {
	q := `DELETE FROM assoc_news_tags WHERE news_id = ANY($1)`

	_, err = db.newsDatabase.Master.Exec(q, pq.Array(newsID))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CleanNewsTagAssoc",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	return nil
}
