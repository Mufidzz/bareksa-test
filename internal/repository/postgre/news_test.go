package postgre

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"reflect"
	"testing"
	"time"
)

//TODO : UPDATE
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
		in := []presentation.CreateNewsRequest{
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
		in := []presentation.CreateNewsRequest{
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

//TODO : UPDATE
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
		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.GetBulkNews(defaultPagination, &presentation.NewsFilter{
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

		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNews(defaultPagination, nil)

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

		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(1, defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNews(defaultPagination, &presentation.NewsFilter{Status: 1})

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

		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(pq.Array([]int{1}), defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNews(defaultPagination, &presentation.NewsFilter{Topics: []int{1}})

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

		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs(1, pq.Array([]int{2}), defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNews(defaultPagination, &presentation.NewsFilter{Status: 1, Topics: []int{2}})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #4 - With All Filters",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #5 - With Filter Title
	t.Run("Success #5 - With Filter Title", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "content", "topics_name", "tags_name", "status"}).
			AddRow(1, now, now, "abc", "xyz", "asd", "asd", 1)

		mock.ExpectQuery("SELECT (.+) FROM news LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) LEFT JOIN (.+) GROUP BY (.+) LIMIT (.+) OFFSET (.+)").
			WithArgs("%TEST%", defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNews(defaultPagination, &presentation.NewsFilter{Title: "TEST"})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #5 - With Filter Title",
				FunctionName: "Test_GetNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})
}

//TODO : UPDATE
func Test_UpdateNews(t *testing.T) {
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
		in := []presentation.UpdateNewsRequest{
			{
				ID:      1,
				Title:   "A",
				Content: "B",
				Status:  2,
			},
		}

		mock.ExpectQuery("UPDATE news SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.UpdateBulkNews(in)

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         "Failed - SQL Return Error",
				FunctionName: "Test_UpdateNews",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - Success Update Some Rows
	t.Run("Success #1 - Success Update Some Rows", func(tt *testing.T) {
		in := []presentation.UpdateNewsRequest{
			{
				ID:      1,
				Title:   "A",
				Content: "B",
				Status:  2,
			},
			{
				ID:      2,
				Title:   "A",
				Content: "B",
				Status:  2,
			},
			{
				ID:      3,
				Title:   "A",
				Content: "B",
				Status:  2,
			},
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("UPDATE news SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
			WithArgs(in[0].ID, in[0].Title, in[0].Content, in[0].Status, in[1].ID, in[1].Title, in[1].Content, in[1].Status, in[2].ID, in[2].Title, in[2].Content, in[2].Status).
			WillReturnRows(rows)

		res, err := pgDB.UpdateBulkNews(in)
		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - Success Update Some Rows",
				FunctionName: "Test_UpdateNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}

		expectedResponse := []int{1}

		if !reflect.DeepEqual(res, expectedResponse) {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - Success Update Some Rows",
				FunctionName: "Test_UpdateNews",
				Description:  "Invalid Function Response",
				Trace:        fmt.Errorf("got %v, expected %v", res, expectedResponse),
			}.Error())
		}
	})

	// Success #2 - Success Update All Rows
	t.Run("Success #2 - Success Update All Rows", func(tt *testing.T) {
		in := []presentation.UpdateNewsRequest{
			{
				ID:      1,
				Title:   "A",
				Content: "B",
				Status:  2,
			},
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1).
			AddRow(2)

		mock.ExpectQuery("UPDATE news SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
			WithArgs(in[0].ID, in[0].Title, in[0].Content, in[0].Status).
			WillReturnRows(rows)

		res, err := pgDB.UpdateBulkNews(in)
		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #2 - Success Update All Rows",
				FunctionName: "Test_UpdateNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}

		expectedResponse := []int{1, 2}

		if !reflect.DeepEqual(res, expectedResponse) {
			tt.Error(response.InternalTestError{
				Name:         "Success #2 - Success Update All Rows",
				FunctionName: "Test_UpdateNews",
				Description:  "Invalid Function Response",
				Trace:        fmt.Errorf("got %v, expected %v", res, expectedResponse),
			}.Error())
		}
	})
}

//TODO : UPDATE
func Test_DeleteBulkNews(t *testing.T) {
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
		in := []int{1, 2, 3}

		mock.ExpectQuery("DELETE FROM (.+) WHERE (.+) RETURNING (.+)").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.DeleteBulkNews(in)

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         "Failed - SQL Return Error",
				FunctionName: "Test_DeleteBulkNews",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - Success Delete Some Rows
	t.Run("Success #1 - Success Delete Some Rows", func(tt *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3).
			AddRow(4)

		mock.ExpectQuery("DELETE FROM (.+) WHERE (.+) RETURNING (.+)").
			WithArgs(pq.Array(in)).
			WillReturnRows(rows)

		res, err := pgDB.DeleteBulkNews(in)
		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - Success Delete Some Rows",
				FunctionName: "Test_DeleteBulkNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}

		expectedResponse := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(res, expectedResponse) {
			tt.Error(response.InternalTestError{
				Name:         "Success #1 - Success Delete Some Rows",
				FunctionName: "Test_DeleteBulkNews",
				Description:  "Invalid Function Response",
				Trace:        fmt.Errorf("got %v, expected %v", res, expectedResponse),
			}.Error())
		}
	})

	// Success #2 - Success Delete All Rows
	t.Run("Success #2 - Success Delete All Rows", func(tt *testing.T) {
		in := []int{1, 2, 3, 4}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1).
			AddRow(2).
			AddRow(3).
			AddRow(4)

		mock.ExpectQuery("DELETE FROM (.+) WHERE (.+) RETURNING (.+)").
			WithArgs(pq.Array(in)).
			WillReturnRows(rows)

		res, err := pgDB.DeleteBulkNews(in)
		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         "Success #2 - Success Delete All Rows",
				FunctionName: "Test_DeleteBulkNews",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}

		expectedResponse := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(res, expectedResponse) {
			tt.Error(response.InternalTestError{
				Name:         "Success #2 - Success Delete All Rows",
				FunctionName: "Test_DeleteBulkNews",
				Description:  "Invalid Function Response",
				Trace:        fmt.Errorf("got %v, expected %v", res, expectedResponse),
			}.Error())
		}
	})

}
