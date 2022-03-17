package postgre

import (
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
	"fmt"
)

func (db *Postgre) CreateBulkTopics(in []presentation.CreateBulkTopicsRequest) (insertedID []int, err error) {
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
			FunctionName: "CreateBulkTopics",
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
				FunctionName: "CreateBulkTopics",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		insertedID = append(insertedID, id)
	}

	return insertedID, nil
}
