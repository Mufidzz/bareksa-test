package postgre

import (
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_CreateBulkTopics(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on Initalize sqlmock, trace %v", err)
	}

	defer db.Close()

	pgDB := NewWithDBObject(&PgDB{
		Master: sqlx.NewDb(db, DB_DRIVER_NAME_SQLMOCK),
	})

	// Failed - SQL Return Error
	t.Run("Failed - SQL Return Error", func(tt *testing.T) {
		in := []presentation.CreateBulkTopicsRequest{
			{
				"AAA",
			},
		}

		mock.ExpectQuery("INSERT INTO news_topics (.+) VALUES (.+) RETURNING id").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.CreateBulkTopics(in)

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         "Failed - SQL Return Error",
				FunctionName: "Test_CreateBulkTopics",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - Success Create New Rows
	t.Run("Success #1 - Success Create New Rows", func(tt *testing.T) {
		in := []presentation.CreateBulkTopicsRequest{
			{
				"AAA",
			},
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("INSERT INTO news_topics (.+) VALUES (.+) RETURNING id").
			WithArgs(in[0].Name).
			WillReturnRows(rows)

		_, err = pgDB.CreateBulkTopics(in)
		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - Success Create New Rows",
				FunctionName: "Test_CreateNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})
}
