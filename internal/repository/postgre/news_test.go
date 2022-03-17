package postgre

import (
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"testing"
	"time"
)

func Test_CreateNews(t *testing.T) {
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
		in := []presentation.CreateBulkNewsRequest{
			{
				Title:   "A",
				Content: "B",
				Status:  2,
			},
		}

		mock.ExpectQuery("INSERT INTO news (.+) VALUES (.+) RETURNING id").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.CreateBulkNews(in)

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         "Failed - SQL Return Error",
				FunctionName: "Test_CreateNews",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - Success Create New Rows
	t.Run("Success #1 - Success Create New Rows", func(tt *testing.T) {
		in := []presentation.CreateBulkNewsRequest{
			{
				Title:   "A",
				Content: "B",
				Status:  2,
			},
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("INSERT INTO news (.+) VALUES (.+) RETURNING id").
			WithArgs(in[0].Title, in[0].Content, in[0].Status).
			WillReturnRows(rows)

		_, err = pgDB.CreateBulkNews(in)
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

func Test_GetNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on Initalize sqlmock, trace %v", err)
	}
	defer db.Close()

	pgDB := NewWithDBObject(&PgDB{
		Master: sqlx.NewDb(db, DB_DRIVER_NAME_SQLMOCK),
	})

	now := time.Now()
	defaultPagination := presentation.Pagination{
		Offset: 0,
		Count:  5,
	}

	// Failed - SQL Return Error
	t.Run("Failed - SQL Return Error", func(tt *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM news INNER JOIN (.+) LEFT JOIN (.+) INNER JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.GetNews(defaultPagination, &presentation.NewsFilter{
			Status: 1,
			Topics: []int{1, 2, 3},
		})

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         "Failed - SQL Return Error",
				FunctionName: "Test_GetNews",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - No Filter
	t.Run("Success #1 - No Filter", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "content", "topics_name", "tags_name", "status"}).
			AddRow(1, now, now, "abc", "xyz", "asd", "asd", 1)

		mock.ExpectQuery("SELECT (.+) FROM news INNER JOIN (.+) LEFT JOIN (.+) INNER JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetNews(defaultPagination, nil)

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - No Filter",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #2 - With Filter Status
	t.Run("Success #2 - With Filter Status", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "content", "topics_name", "tags_name", "status"}).
			AddRow(1, now, now, "abc", "xyz", "asd", "asd", 1)

		mock.ExpectQuery("SELECT (.+) FROM news INNER JOIN (.+) LEFT JOIN (.+) INNER JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(1, defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetNews(defaultPagination, &presentation.NewsFilter{Status: 1})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #2 - With Filter Status",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #3 - With Filter Topics
	t.Run("Success #3 - With Filter Topics", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "content", "topics_name", "tags_name", "status"}).
			AddRow(1, now, now, "abc", "xyz", "asd", "asd", 1)

		mock.ExpectQuery("SELECT (.+) FROM news INNER JOIN (.+) LEFT JOIN (.+) INNER JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(pq.Array([]int{1}), defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetNews(defaultPagination, &presentation.NewsFilter{Topics: []int{1}})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #3 - With Filter Topics",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #4 - With All Filters
	t.Run("Success #4 - With All Filters", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "content", "topics_name", "tags_name", "status"}).
			AddRow(1, now, now, "abc", "xyz", "asd", "asd", 1)

		mock.ExpectQuery("SELECT (.+) FROM news INNER JOIN (.+) LEFT JOIN (.+) INNER JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(1, pq.Array([]int{2}), defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetNews(defaultPagination, &presentation.NewsFilter{Status: 1, Topics: []int{2}})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #4 - With All Filters",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})
}
