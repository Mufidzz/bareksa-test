package postgre

import (
	"Test_Bareksa/pkg/response"
	"Test_Bareksa/presentation"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func Test_CreateBulkNewsTags(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	testcase := []struct {
		name       string
		in         []presentation.CreateBulkTagsRequest
		mockExp    func(mm sqlmock.Sqlmock)
		mustReturn []int
		mustErr    bool
	}{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectQuery("INSERT INTO news_tags (.+) VALUES (.+) RETURNING id").
					WillReturnError(fmt.Errorf("hello"))
			},
			in: []presentation.CreateBulkTagsRequest{
				{"A"},
				{"B"},
				{"C"},
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Success Create New Rows",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1).
					AddRow(2)

				mm.ExpectQuery("INSERT INTO news_tags (.+) VALUES (.+) RETURNING id").
					WillReturnRows(rows)
			},
			in: []presentation.CreateBulkTagsRequest{
				{"A"},
				{"B"},
			},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockExp(mock)
			res, err := pgDB.CreateBulkNewsTags(tc.in)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateBulkNewsTags",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}

}

func Test_GetBulkNewsTags(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	testcase := []struct {
		name       string
		filter     *presentation.NewsTagsFilter
		pagination *presentation.Pagination
		mockExp    func(mm sqlmock.Sqlmock)
		mustReturn []presentation.GetNewsTagsResponse
		mustErr    bool
	}{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnError(fmt.Errorf("hello"))
			},
			filter:     nil,
			pagination: nil,
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success #1 - No Filter, No Pagination",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "AVC").
					AddRow(2, "AVC")

				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnRows(rows)
			},
			filter:     nil,
			pagination: nil,
			mustReturn: []presentation.GetNewsTagsResponse{
				{ID: 1, Name: "AVC"},
				{ID: 2, Name: "AVC"},
			},
			mustErr: false,
		},
		{
			name: "Success #2 - No Filter",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "AVC")

				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnRows(rows)
			},
			filter: nil,
			pagination: &presentation.Pagination{
				Offset: 0,
				Count:  1,
			},
			mustReturn: []presentation.GetNewsTagsResponse{
				{ID: 1, Name: "AVC"},
			},
			mustErr: false,
		},
		{
			name: "Success #3 - No Pagination",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "AVC")

				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnRows(rows)
			},
			filter: &presentation.NewsTagsFilter{
				Name:        "ASDASD",
				NewsTopicID: 1,
			},
			pagination: nil,
			mustReturn: []presentation.GetNewsTagsResponse{
				{ID: 1, Name: "AVC"},
			},
			mustErr: false,
		},
		{
			name: "Success #4 - Filter Name",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "ASDASD")

				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnRows(rows)
			},
			filter: &presentation.NewsTagsFilter{
				Name: "ASDASD",
			},
			pagination: nil,
			mustReturn: []presentation.GetNewsTagsResponse{
				{ID: 1, Name: "ASDASD"},
			},
			mustErr: false,
		},
		{
			name: "Success #5 - Filter ID",
			mockExp: func(mm sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "ASDASD")

				mm.ExpectQuery("SELECT (.+) FROM news_tags").
					WillReturnRows(rows)
			},
			filter: &presentation.NewsTagsFilter{
				NewsTopicID: 1,
			},
			pagination: nil,
			mustReturn: []presentation.GetNewsTagsResponse{
				{ID: 1, Name: "ASDASD"},
			},
			mustErr: false,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(tt *testing.T) {
			tc.mockExp(mock)
			res, err := pgDB.GetBulkNewsTags(tc.pagination, tc.filter)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_GetBulkNewsTags",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_UpdateBulkNewsTags(t *testing.T) {
	pgDB, db, mock, err := initDB()
	if err != nil {
		t.Fatalf("Failed init Mock Database")
	}
	defer db.Close()

	testcase := []struct {
		name       string
		in         []presentation.UpdateNewsTagsRequest
		mockExp    func(mm sqlmock.Sqlmock)
		mustReturn []int
		mustErr    bool
	}{
		{
			name: "Failed - SQL Return Error",
			mockExp: func(mm sqlmock.Sqlmock) {
				mm.ExpectQuery("UPDATE news_tags SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
					WillReturnError(fmt.Errorf("hello"))
			},
			in: []presentation.UpdateNewsTagsRequest{
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
				mm.ExpectQuery("UPDATE news_tags SET (.+) FROM (.+) WHERE (.+) RETURNING (.+)").
					WillReturnRows(rows)
			},
			in: []presentation.UpdateNewsTagsRequest{
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
			res, err := pgDB.UpdateBulkNewsTags(tc.in)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_UpdateBulkNewsTags",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}

}
func Test_DeleteBulkNewsTags(t *testing.T) {
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
				mm.ExpectQuery("DELETE FROM news_tags WHERE (.+)").
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
				mm.ExpectQuery("DELETE FROM news_tags WHERE (.+)").
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
				mm.ExpectQuery("DELETE FROM news_tags WHERE (.+)").
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
			res, err := pgDB.DeleteBulkNewsTags(tc.in)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_DeleteBulkNewsTags",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}
