package redis

import (
	"encoding/json"
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/go-redis/redismock/v8"
	"reflect"
	"testing"
	"time"
)

func Test_SaveObject(t *testing.T) {
	exp := time.Duration(10) * time.Second

	testcases := []struct {
		name    string
		key     string
		value   string
		mock    func(redismock.ClientMock)
		mustErr bool
	}{
		{
			name:  "Success",
			key:   "A",
			value: "B",
			mock: func(mock redismock.ClientMock) {
				objectJson, _ := json.Marshal("B")
				mock.ExpectSet("A", objectJson, exp).
					SetVal("")
			},

			mustErr: false,
		},
		{
			name:  "Failed",
			key:   "A",
			value: "B",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectSet("A", "B", exp).
					SetErr(fmt.Errorf("FFF"))
			},
			mustErr: true,
		},
		{
			name:  "Failed - Is not JSON",
			key:   "A",
			value: "B",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectSet("A", "B", exp).
					SetVal("")
			},
			mustErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			db, mock := redismock.NewClientMock()
			tc.mock(mock)

			r := NewFromObject(db)
			err := r.SaveObject(tc.key, tc.value)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tc.name,
					FunctionName: "Test_SaveObject",
					Description:  "Failed Running Testcase",
					Trace:        err,
				})
			}
		})
	}
}

func Test_GetObject(t *testing.T) {
	testcases := []struct {
		name       string
		key        string
		mock       func(redismock.ClientMock)
		mustErr    bool
		mustReturn string
	}{
		{
			name: "Success",
			key:  "A",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectGet("A").
					SetVal(`{"name" : 4}`)
			},
			mustReturn: `{"name":4}`,
			mustErr:    false,
		},
		{
			name: "Failed",
			key:  "A",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectGet("A").
					SetErr(fmt.Errorf("FFF"))
			},
			mustReturn: "null",
			mustErr:    true,
		},
		{
			name: "Failed - Cant Bind JSON",
			key:  "A",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectGet("A").
					SetVal(`{"namae"adw : 4}`)
			},
			mustReturn: "null",
			mustErr:    true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			db, mock := redismock.NewClientMock()
			tc.mock(mock)

			r := NewFromObject(db)

			var jsonRes map[string]interface{}

			err := r.GetObject(tc.key, &jsonRes)

			jsonResString, _ := json.Marshal(jsonRes)

			fmt.Println(jsonResString)
			fmt.Println([]byte(tc.mustReturn))
			fmt.Println(string(jsonResString))
			fmt.Println(tc.mustReturn)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual([]byte(tc.mustReturn), jsonResString) {
				tt.Error(response.InternalTestError{
					Name:         tc.name,
					FunctionName: "Test_SaveObject",
					Description:  "Failed Running Testcase",
					Trace:        err,
				})
			}
		})
	}
}

func Test_FlushAll(t *testing.T) {
	testcases := []struct {
		name    string
		mock    func(redismock.ClientMock)
		mustErr bool
	}{
		{
			name: "Success",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectFlushAll().
					SetVal("")
			},
			mustErr: false,
		},
		{
			name: "Failed",
			mock: func(mock redismock.ClientMock) {
				mock.ExpectFlushAll().
					SetErr(fmt.Errorf("FFF"))
			},
			mustErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			db, mock := redismock.NewClientMock()
			tc.mock(mock)

			r := NewFromObject(db)

			err := r.FlushAll()

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tc.name,
					FunctionName: "Test_SaveObject",
					Description:  "Failed Running Testcase",
					Trace:        err,
				})
			}
		})
	}
}
