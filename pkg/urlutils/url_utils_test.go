package urlutils

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"reflect"
	"testing"
)

func Test_DecodeEncodedString(t *testing.T) {
	type inputParam struct {
		encodedFilterString string
	}

	testcases := []struct {
		name       string
		in         inputParam
		mustErr    bool
		mustReturn interface{}
	}{
		{
			name:       "Failed - Unknown base64",
			in:         inputParam{encodedFilterString: "adwadawjkhd"},
			mustErr:    true,
			mustReturn: presentation.Pagination{},
		},
		{
			name:    "Success - Parse Pagination",
			in:      inputParam{encodedFilterString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9"},
			mustErr: false,
			mustReturn: presentation.Pagination{
				Offset: 1,
				Count:  7,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			var res presentation.Pagination
			err := DecodeEncodedString(tc.in.encodedFilterString, &res)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, res) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_DecodeEncodedString",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", res, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_EncodeStruct(t *testing.T) {
	testcases := []struct {
		name       string
		in         interface{}
		mustErr    bool
		mustReturn string
	}{
		{
			name: "Success - Convert Pagination",
			in: presentation.Pagination{
				Offset: 0,
				Count:  1,
			},
			mustErr:    false,
			mustReturn: "eyJvZmZzZXQiOjAsImNvdW50IjoxfQ--",
		},

		{
			name: "Success - Convert News Filter",
			in: presentation.NewsFilter{
				Status: 1,
				Topics: []int{1, 2, 3, 4},
				NewsID: 2,
			},
			mustErr:    false,
			mustReturn: "eyJzdGF0dXMiOjEsInRvcGljcyI6WzEsMiwzLDRdLCJuZXdzX2lkIjoyfQ--",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			got, err := EncodeStruct(tc.in)

			if (tc.mustErr && err == nil) || !reflect.DeepEqual(tc.mustReturn, got) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_EncodeStruct",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}
