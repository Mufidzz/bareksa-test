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

func Test_HandleCreateNewsTopic(t *testing.T) {
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
			url:            "/news-topic",
			handler:        NewHTTP(nil, nil, &MockNewsTopicDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Create News Topics",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           `[{"name" : "A"}, {"name" : "B"}]`,
			url:            "/news-topic",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				createNewsTopics: createNewsTopics{err: fmt.Errorf("a")},
			}),
		},
		{
			name:           "Success",
			mustReturn:     "",
			mustReturnCode: http.StatusNoContent,
			body:           `[{"name" : "A"}, {"name" : "B"}]`,
			url:            "/news-topic",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				createNewsTopics: createNewsTopics{insertedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.POST("/news-topic", tc.handler.HandleCreateNewsTopic)
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
					FunctionName: "Test_HandleCreateNewsTopic",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleGetNewsTopic(t *testing.T) {
	defaultPagination, err := urlutils.EncodeStruct(presentation.Pagination{
		Offset: 0,
		Count:  1,
	})
	if err != nil {
		t.Fatal("Failed Generate Pagination Encoded String")
	}

	defaultFilter, err := urlutils.EncodeStruct(presentation.NewsTopicFilter{
		Name:        "ABC",
		NewsTopicID: 1,
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
			url:            fmt.Sprintf("/news-topic?pagination=%sawdad", defaultPagination),
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				getNewsTopics: getNewsTopics{err: fmt.Errorf("invalid pagination")},
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
			url:            "/news-topic",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				getNewsTopics: getNewsTopics{err: fmt.Errorf("other error")},
			}),
		},
		{
			name: "Success #1",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: []presentation.GetNewsTopicsResponse{
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
			url:            fmt.Sprintf("/news-topic?pagination=%s", defaultPagination),
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				getNewsTopics: getNewsTopics{
					res: []presentation.GetNewsTopicsResponse{
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
				Data: []presentation.GetNewsTopicsResponse{
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
			url:            fmt.Sprintf("/news-topic?pagination=%s&filter=%s", defaultPagination, defaultFilter),
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				getNewsTopics: getNewsTopics{
					res: []presentation.GetNewsTopicsResponse{
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
			router.GET("/news-topic", tc.handler.HandleGetNewsTopic)
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

func Test_HandleUpdateNewsTopic(t *testing.T) {
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
			url:            "/news-topic",
			handler:        NewHTTP(nil, nil, &MockNewsTopicDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Update News Topics",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusInternalServerError,
			body:           `[{"id" : 1, "name" : "A"}, {"id" : 2, "name" : "B"}]`,
			url:            "/news-topic",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				updateNewsTopics: updateNewsTopics{err: fmt.Errorf("a")},
			}),
		},
		{
			name: "Success",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Updated News Topics",
				Data: gin.H{
					"updated_id": []int{1, 2},
				},
			},
			mustReturnCode: http.StatusOK,
			body:           `[{"id" : 1, "name" : "A"}, {"id" : 2, "name" : "B"}]`,
			url:            "/news-topic",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				updateNewsTopics: updateNewsTopics{updatedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.PUT("/news-topic", tc.handler.HandleUpdateNewsTopic)
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
					FunctionName: "Test_HandleUpdateNewsTopic",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleDeleteNewsTopic(t *testing.T) {
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
			url:            "/news-topic",
			handler:        NewHTTP(nil, nil, &MockNewsTopicDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Delete News Topic",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news-topic?id=1&id=2",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				deleteNewsTopics: deleteNewsTopics{err: fmt.Errorf("a")},
			}),
		},
		{
			name: "Success",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Delete News Topics",
				Data: gin.H{
					"deleted_id": []int{1, 2},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            "/news-topic?id=1&id=2",
			handler: NewHTTP(nil, nil, &MockNewsTopicDataUC{
				deleteNewsTopics: deleteNewsTopics{deletedID: []int{1, 2}},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.url, nil)

			router := gin.Default()
			router.DELETE("/news-topic", tc.handler.HandleDeleteNewsTopic)
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
					FunctionName: "Test_HandleCreateNewsTopic",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}
