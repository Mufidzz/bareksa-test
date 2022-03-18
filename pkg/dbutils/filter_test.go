package dbutils

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"reflect"
	"testing"
)

func Test_AddFilter(t *testing.T) {
	type inputParam struct {
		str        string
		connector  string
		key        string
		comparator string
		paramCount int
	}

	testcases := []struct {
		name string
		in   inputParam
		out  string
	}{
		{
			name: "String not include Where",
			in: inputParam{
				str:        "A",
				connector:  CONNECTOR_OR,
				key:        "x",
				comparator: COMPARATOR_EQUAL,
				paramCount: 1,
			},
			out: "A WHERE x = $1",
		},

		{
			name: "String include Where",
			in: inputParam{
				str:        "A WHERE y = 1",
				connector:  CONNECTOR_AND,
				key:        "x",
				comparator: COMPARATOR_EQUAL,
				paramCount: 1,
			},
			out: "A WHERE y = 1 AND x = $1",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			res := AddFilter(tc.in.str, tc.in.connector, tc.in.key, tc.in.comparator, tc.in.paramCount)

			if !reflect.DeepEqual(res, tc.out) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_AddFilter",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v", res, tc.out),
				}.Error())
			}
		})
	}
}

func Test_AddCustomFilter(t *testing.T) {
	type inputParam struct {
		str        string
		connector  string
		key        string
		comparator string
		param      string
	}

	testcases := []struct {
		name string
		in   inputParam
		out  string
	}{
		{
			name: "String not include Where",
			in: inputParam{
				str:        "A",
				connector:  CONNECTOR_OR,
				key:        "x",
				comparator: COMPARATOR_EQUAL,
				param:      "X",
			},
			out: "A WHERE x = X",
		},

		{
			name: "String include Where",
			in: inputParam{
				str:        "A WHERE y = 1",
				connector:  CONNECTOR_AND,
				key:        "x",
				comparator: COMPARATOR_EQUAL,
				param:      "ZZ",
			},
			out: "A WHERE y = 1 AND x = ZZ",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			res := AddCustomFilter(tc.in.str, tc.in.connector, tc.in.key, tc.in.comparator, tc.in.param)

			if !reflect.DeepEqual(res, tc.out) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_AddCustomFilter",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v", res, tc.out),
				}.Error())
			}
		})
	}
}
