package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/pkg/urlutils"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_HandleCreateNewsTag(t *testing.T) {
	testcases := []struct {
		name           string
		url            string
		body           string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid JSON",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Binding JSON",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "",
			url:            "/news-tag",
			handler:        NewHTTP(nil, nil, nil, &MockNewsTagDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Create News Tags",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           `[{"name" : "A"}, {"name" : "B"}]`,
			url:            "/news-tag",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				createNewsTags: createNewsTags{err: fmt.Errorf("a")},
			}),
		},
		{
			name:           "Success",
			mustReturn:     "",
			mustReturnCode: http.StatusNoContent,
			body:           `[{"name" : "A"}, {"name" : "B"}]`,
			url:            "/news-tag",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				createNewsTags: createNewsTags{insertedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.POST("/news-tag", tc.handler.HandleCreateNewsTag)
			router.ServeHTTP(w, req)

			var jsonMustResponse []byte
			var err error
			if tc.mustReturn != "" {
				jsonMustResponse, err = json.Marshal(tc.mustReturn)
				if err != nil {
					tt.Fatal("Failed Creating JSON String")
				}
			}

			if tc.mustReturnCode != w.Code || !reflect.DeepEqual(jsonMustResponse, w.Body.Bytes()) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_HandleCreateNewsTag",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleGetNewsTag(t *testing.T) {
	defaultPagination, err := urlutils.EncodeStruct(presentation.Pagination{
		Offset: 0,
		Count:  1,
	})
	if err != nil {
		t.Fatal("Failed Generate Pagination Encoded String")
	}

	defaultFilter, err := urlutils.EncodeStruct(presentation.NewsTagsFilter{
		Name:      "ABC",
		NewsTagID: 1,
	})

	if err != nil {
		t.Fatal("Failed Generate Filter Encoded String")
	}

	testcases := []struct {
		name           string
		url            string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid Pagination",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusInternalServerError,
			url:            fmt.Sprintf("/news-tag?pagination=%sawdad", defaultPagination),
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				getNewsTags: getNewsTags{err: fmt.Errorf("invalid pagination")},
			}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusInternalServerError,
			url:            "/news-tag",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				getNewsTags: getNewsTags{err: fmt.Errorf("other error")},
			}),
		},
		{
			name: "Success #1",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: []presentation.GetNewsTagsResponse{
					{
						ID:   1,
						Name: "A",
					},
					{
						ID:   2,
						Name: "B",
					},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            fmt.Sprintf("/news-tag?pagination=%s", defaultPagination),
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				getNewsTags: getNewsTags{
					res: []presentation.GetNewsTagsResponse{
						{
							ID:   1,
							Name: "A",
						},
						{
							ID:   2,
							Name: "B",
						},
					},
				},
			}),
		},
		{
			name: "Success #2 - With Filter",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: []presentation.GetNewsTagsResponse{
					{
						ID:   1,
						Name: "A",
					},
					{
						ID:   2,
						Name: "B",
					},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            fmt.Sprintf("/news-tag?pagination=%s&filter=%s", defaultPagination, defaultFilter),
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				getNewsTags: getNewsTags{
					res: []presentation.GetNewsTagsResponse{
						{
							ID:   1,
							Name: "A",
						},
						{
							ID:   2,
							Name: "B",
						},
					},
				},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.url, nil)

			router := gin.Default()
			router.GET("/news-tag", tc.handler.HandleGetNewsTag)
			router.ServeHTTP(w, req)

			var jsonMustResponse []byte
			var err error
			if tc.mustReturn != "" {
				jsonMustResponse, err = json.Marshal(tc.mustReturn)
				if err != nil {
					tt.Fatal("Failed Creating JSON String")
				}
			}

			if tc.mustReturnCode != w.Code || !reflect.DeepEqual(jsonMustResponse, w.Body.Bytes()) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_HandleGetNews",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v, code %v", w.Body.String(), string(jsonMustResponse), w.Code),
				}.Error())
			}
		})
	}
}

func Test_HandleUpdateNewsTag(t *testing.T) {
	testcases := []struct {
		name           string
		url            string
		body           string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid JSON",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Binding JSON",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "",
			url:            "/news-tag",
			handler:        NewHTTP(nil, nil, nil, &MockNewsTagDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Update News Tags",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusInternalServerError,
			body:           `[{"id" : 1, "name" : "A"}, {"id" : 2, "name" : "B"}]`,
			url:            "/news-tag",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				updateNewsTags: updateNewsTags{err: fmt.Errorf("a")},
			}),
		},
		{
			name: "Success",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Updated News Tags",
				Data: gin.H{
					"updated_id": []int{1, 2},
				},
			},
			mustReturnCode: http.StatusOK,
			body:           `[{"id" : 1, "name" : "A"}, {"id" : 2, "name" : "B"}]`,
			url:            "/news-tag",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				updateNewsTags: updateNewsTags{updatedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.PUT("/news-tag", tc.handler.HandleUpdateNewsTag)
			router.ServeHTTP(w, req)

			var jsonMustResponse []byte
			var err error
			if tc.mustReturn != "" {
				jsonMustResponse, err = json.Marshal(tc.mustReturn)
				if err != nil {
					tt.Fatal("Failed Creating JSON String")
				}
			}

			if tc.mustReturnCode != w.Code || !reflect.DeepEqual(jsonMustResponse, w.Body.Bytes()) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_HandleUpdateNewsTag",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleDeleteNewsTag(t *testing.T) {
	testcases := []struct {
		name           string
		url            string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - No ID",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "ID Must more than 1",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news-tag",
			handler:        NewHTTP(nil, nil, nil, &MockNewsTagDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Delete News Tag",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news-tag?id=1&id=2",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				deleteNewsTags: deleteNewsTags{err: fmt.Errorf("a")},
			}),
		},
		{
			name: "Success",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Delete News Tags",
				Data: gin.H{
					"deleted_id": []int{1, 2},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            "/news-tag?id=1&id=2",
			handler: NewHTTP(nil, nil, nil, &MockNewsTagDataUC{
				deleteNewsTags: deleteNewsTags{deletedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.url, nil)

			router := gin.Default()
			router.DELETE("/news-tag", tc.handler.HandleDeleteNewsTag)
			router.ServeHTTP(w, req)

			var jsonMustResponse []byte
			var err error
			if tc.mustReturn != "" {
				jsonMustResponse, err = json.Marshal(tc.mustReturn)
				if err != nil {
					tt.Fatal("Failed Creating JSON String")
				}
			}

			if tc.mustReturnCode != w.Code || !reflect.DeepEqual(jsonMustResponse, w.Body.Bytes()) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_HandleCreateNewsTag",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}
