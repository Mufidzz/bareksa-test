package postgre

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

//TODO : UPDATE
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
		in := []presentation.CreateNewsTopicsRequest{
			{
				"AAA",
			},
		}

		mock.ExpectQuery("INSERT INTO news_topics (.+) VALUES (.+) RETURNING id").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.CreateBulkNewsTopics(in)

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
		in := []presentation.CreateNewsTopicsRequest{
			{
				"AAA",
			},
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("INSERT INTO news_topics (.+) VALUES (.+) RETURNING id").
			WithArgs(in[0].Name).
			WillReturnRows(rows)

		_, err = pgDB.CreateBulkNewsTopics(in)
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
func Test_GetBulkNewsTopics(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on Initalize sqlmock, trace %v", err)
	}
	defer db.Close()

	pgDB := NewWithDBObject(&PgDB{
		Master: sqlx.NewDb(db, DB_DRIVER_NAME_SQLMOCK),
	})

	defaultPagination := presentation.Pagination{
		Offset: 0,
		Count:  5,
	}

	// Failed - SQL Return Error
	t.Run("Failed - SQL Return Error", func(tt *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM news_topics").
			WillReturnError(fmt.Errorf("hello"))

		_, err = pgDB.GetBulkNewsTopics(nil, nil)

		if err == nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Expected From Function",
				Trace:        nil,
			}.Error())
		}
	})

	// Success #1 - No Filter
	t.Run("Success #1 - No Filter", func(tt *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "AA")

		mock.ExpectQuery("SELECT (.+) FROM news_topics").
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(nil, nil)

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #2 - Filter Name
	t.Run("Success #2 - Filter Name", func(tt *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "ABCDE")

		mock.ExpectQuery("SELECT (.+) FROM news_topics (.+)").
			WithArgs("%X%", defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(&defaultPagination, &presentation.NewsTopicFilter{
			Name: "X",
		})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #3 - Filter Topics ID
	t.Run("Success #3 - Filter Topics ID", func(tt *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "ABCDE")

		mock.ExpectQuery("SELECT (.+) FROM news_topics (.+)").
			WithArgs(1, defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(&defaultPagination, &presentation.NewsTopicFilter{
			NewsTopicID: 1,
		})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #4 - Filter Topics ID & Name
	t.Run("Success #4 - Filter Topics ID & Name", func(tt *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "ABCDE")

		mock.ExpectQuery("SELECT (.+) FROM news_topics (.+)").
			WithArgs(1, "%ABCDE%", defaultPagination.Count, defaultPagination.Offset).
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(&defaultPagination, &presentation.NewsTopicFilter{
			NewsTopicID: 1,
			Name:        "ABCDE",
		})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #5 - Filter Topics ID & Name No Pagination
	t.Run("Success #5 - Filter Topics ID & Name No Pagination", func(tt *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "ABCDE")

		mock.ExpectQuery("SELECT (.+) FROM news_topics (.+)").
			WithArgs(1, "%ABCDE%").
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(nil, &presentation.NewsTopicFilter{
			NewsTopicID: 1,
			Name:        "ABCDE",
		})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})

	// Success #5 - Filter Topics ID & Name No Pagination
	t.Run("Success #5 - Filter Topics ID & Name No Pagination", func(tt *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "ABCDE")

		mock.ExpectQuery("SELECT (.+) FROM news_topics (.+)").
			WithArgs(1, "%ABCDE%").
			WillReturnRows(rows)

		_, err = pgDB.GetBulkNewsTopics(nil, &presentation.NewsTopicFilter{
			NewsTopicID: 1,
			Name:        "ABCDE",
		})

		if err != nil {
			tt.Error(response.InternalTestError{
				Name:         tt.Name(),
				FunctionName: "Test_GetBulkNewsTopics",
				Description:  "Error Not Expected From Function",
				Trace:        err,
			}.Error())
		}
	})
}

func Test_UpdateBulkNewsTopics(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	type tests struct {
		name       string
		in         []presentation.UpdateNewsTopicsRequest
		mockExp    func(mm sqlmock.Sqlmock)
		mustReturn []int
		mustErr    bool
	}

	testcase := []tests{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectQuery("UPDATE news_topics SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
					WillReturnError(fmt.Errorf("hello"))
			},
			in: []presentation.UpdateNewsTopicsRequest{
				{
					1, "Test",
				},
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Update All Rows",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1).
					AddRow(2).
					AddRow(3).
					AddRow(4)
				mm.ExpectQuery("UPDATE news_topics SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
					WillReturnRows(rows)
			},
			in: []presentation.UpdateNewsTopicsRequest{
				{1, "Test"},
				{2, "Test"},
				{3, "Test"},
				{4, "Test"},
			},
			mustReturn: []int{1, 2, 3, 4},
			mustErr:    false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockExp(mock)
			res, err := pgDB.UpdateBulkNewsTopics(tc.in)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_Bulk_UpdateNewsTopics",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}

		})
	}

}

func Test_DeleteBulkNewsTopics(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	testcase := []struct {
		name       string
		in         []int
		mockExp    func(mm sqlmock.Sqlmock)
		mustReturn []int
		mustErr    bool
	}{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectQuery("DELETE FROM news_topics WHERE (.+)").
					WillReturnError(fmt.Errorf("hello"))
			},
			in:         []int{1, 2},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Delete All Rows",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1).
					AddRow(2)
				mm.ExpectQuery("DELETE FROM news_topics WHERE (.+)").
					WillReturnRows(rows)
			},
			in:         []int{1, 2},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
		{
			name: "Success - Delete Partial Rows",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1).
					AddRow(2)
				mm.ExpectQuery("DELETE FROM news_topics WHERE (.+)").
					WillReturnRows(rows)
			},
			in:         []int{1, 2, 3, 4},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockExp(mock)
			res, err := pgDB.DeleteBulkNewsTopics(tc.in)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_Bulk_UpdateNewsTopics",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_CreateBulkNewsTopicsAssoc(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	testcase := []struct {
		name    string
		in      []presentation.CreateNewsTopicsAssoc
		mockExp func(mm sqlmock.Sqlmock)
		mustErr bool
	}{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectExec("INSERT INTO assoc_news_topics (.+) VALUES (.+)").
					WillReturnError(fmt.Errorf("hello"))
			},
			in: []presentation.CreateNewsTopicsAssoc{
				{
					NewsID:       1,
					NewsTopicsID: []int{1, 2, 3, 4},
				},
			},
			mustErr: true,
		},
		{
			name: "Failed - Not All Data Inserted",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectExec("INSERT INTO assoc_news_topics (.+) VALUES (.+)").
					WillReturnResult(sqlmock.NewResult(4, 12341))
			},
			in: []presentation.CreateNewsTopicsAssoc{
				{
					NewsID:       1,
					NewsTopicsID: []int{1, 2, 3, 4},
				},
			},
			mustErr: true,
		},
		{
			name: "Success - Success Create New Rows",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectExec("INSERT INTO assoc_news_topics (.+) VALUES (.+)").
					WillReturnResult(sqlmock.NewResult(4, 4))
			},
			in: []presentation.CreateNewsTopicsAssoc{
				{
					NewsID:       1,
					NewsTopicsID: []int{1, 2, 3, 4},
				},
			},
			mustErr: false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockExp(mock)
			err := pgDB.CreateBulkNewsTopicsAssoc(tc.in)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateBulkNewsTopicsAssoc",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}

}
