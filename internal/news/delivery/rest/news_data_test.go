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
	"time"
)

func Test_HandleGetNews(t *testing.T) {
	now := time.Now()

	defaultPagination, err := urlutils.EncodeStruct(presentation.Pagination{
		Offset: 0,
		Count:  1,
	})

	defaultFilter, err := urlutils.EncodeStruct(presentation.NewsFilter{
		Status: 1,
		Topics: []int{1, 2, 3},
		NewsID: 1,
	})

	if err != nil {
		t.Fatal("Failed Generate Pagination Encoded String")
	}

	testcases := []struct {
		name           string
		url            string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - No Pagination",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Pagination Cannot be Blank, use URL Query to assign",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Invalid Pagination",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusInternalServerError,
			url:            fmt.Sprintf("/news?pagination=%sawdad", defaultPagination),
			handler: NewHTTP(nil, &MockNewsDataUC{
				getNews: getNews{err: fmt.Errorf("Adwde")},
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
			url:            fmt.Sprintf("/news?pagination=%s", defaultPagination),
			handler: NewHTTP(nil, &MockNewsDataUC{
				getNews: getNews{err: fmt.Errorf("Adwde")},
			}),
		},
		{
			name: "Success #1",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: []presentation.GetNewsResponse{
					{
						ID:         1,
						CreatedAt:  now,
						UpdatedAt:  now,
						Title:      "A",
						Content:    "B",
						TopicsName: "C",
						TagsName:   "D",
						Status:     1,
					},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            fmt.Sprintf("/news?pagination=%s", defaultPagination),
			handler: NewHTTP(nil, &MockNewsDataUC{
				getNews: getNews{
					res: []presentation.GetNewsResponse{
						{
							ID:         1,
							CreatedAt:  now,
							UpdatedAt:  now,
							Title:      "A",
							Content:    "B",
							TopicsName: "C",
							TagsName:   "D",
							Status:     1,
						},
					},
					err: nil,
				},
			}),
		},
		{
			name: "Success #2 - With Filter",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: []presentation.GetNewsResponse{
					{
						ID:         1,
						CreatedAt:  now,
						UpdatedAt:  now,
						Title:      "A",
						Content:    "B",
						TopicsName: "C",
						TagsName:   "D",
						Status:     1,
					},
				},
			},
			mustReturnCode: http.StatusOK,
			url:            fmt.Sprintf("/news?pagination=%s&filter=%s", defaultPagination, defaultFilter),
			handler: NewHTTP(nil, &MockNewsDataUC{
				getNews: getNews{
					res: []presentation.GetNewsResponse{
						{
							ID:         1,
							CreatedAt:  now,
							UpdatedAt:  now,
							Title:      "A",
							Content:    "B",
							TopicsName: "C",
							TagsName:   "D",
							Status:     1,
						},
					},
					err: nil,
				},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.url, nil)

			router := gin.Default()
			router.GET("/news", tc.handler.HandleGetNews)
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

func Test_HandleGetSingleNews(t *testing.T) {
	now := time.Now()

	testcases := []struct {
		name           string
		url            string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid ID Param",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Parsing News ID, Please check news id is valid Number",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news/alkdjhaqwd",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get Single News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				getSingleNews: getSingleNews{err: fmt.Errorf("Adwde")},
			}),
		},
		{
			name: "Success",
			mustReturn: response.SuccessResponse{
				Success: true,
				Message: "Success Getting News",
				Data: presentation.GetNewsResponse{
					ID:         1,
					CreatedAt:  now,
					UpdatedAt:  now,
					Title:      "ABC",
					Content:    "DEF",
					TopicsName: "GHI",
					TagsName:   "JKL",
					Status:     1,
				},
			},
			mustReturnCode: http.StatusOK,
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				getSingleNews: getSingleNews{
					res: presentation.GetNewsResponse{
						ID:         1,
						CreatedAt:  now,
						UpdatedAt:  now,
						Title:      "ABC",
						Content:    "DEF",
						TopicsName: "GHI",
						TagsName:   "JKL",
						Status:     1,
					},
					err: nil,
				},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc.url, nil)

			router := gin.Default()
			router.GET("/news/:id", tc.handler.HandleGetSingleNews)
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
					FunctionName: "Test_HandleGetSingleNews",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v, code %v", w.Body.String(), string(jsonMustResponse), w.Code),
				}.Error())
			}
		})
	}
}

func Test_HandleCreateSingleNews(t *testing.T) {
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
				Data:    presentation.CreateNewsRequest{},
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "",
			url:            "/news",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get Single News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "{\"title\":\"awdawd\",\"content\":\"adwdwa\",\"status\":1}",
			url:            "/news",
			handler: NewHTTP(nil, &MockNewsDataUC{
				createSingleNews: createSingleNews{err: fmt.Errorf("Adwde")},
			}),
		},
		{
			name:           "Success",
			mustReturn:     "",
			mustReturnCode: http.StatusNoContent,
			body:           "{\"title\":\"AAA\",\"content\":\"XCZX\",\"status\":1}",
			url:            "/news",
			handler: NewHTTP(nil, &MockNewsDataUC{
				createSingleNews: createSingleNews{err: nil},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.POST("/news", tc.handler.HandleCreateSingleNews)
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
					FunctionName: "Test_HandleCreateSingleNews",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleUpdateSingleNews(t *testing.T) {
	testcases := []struct {
		name           string
		url            string
		body           string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid ID Param",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Parsing News ID, Please check news id is valid Number",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "",
			url:            "/news/alkdjhaqwd",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Invalid JSON",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Binding JSON",
				Type:    0,
				Data:    presentation.UpdateNewsRequest{},
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "",
			url:            "/news/1",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get Single News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			body:           "{\"title\":\"AAA\",\"content\":\"XCZX\",\"status\":1}",
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				updateSingleNews: updateSingleNews{err: fmt.Errorf("Adwde")},
			}),
		},
		{
			name:           "Success",
			mustReturn:     "",
			mustReturnCode: http.StatusNoContent,
			body:           "{\"title\":\"AAA\",\"content\":\"XCZX\",\"status\":1}",
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				updateSingleNews: updateSingleNews{err: nil},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", tc.url, bytes.NewBuffer([]byte(tc.body)))

			router := gin.Default()
			router.PUT("/news/:id", tc.handler.HandleUpdateSingleNews)
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
					FunctionName: "Test_HandleUpdateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %s, expected %v", w.Body.String(), string(jsonMustResponse)),
				}.Error())
			}
		})
	}
}

func Test_HandleDeleteSingleNews(t *testing.T) {
	testcases := []struct {
		name           string
		url            string
		mustReturn     interface{}
		mustReturnCode int
		handler        *HTTPHandler
	}{
		{
			name: "Failed - Invalid ID Param",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Parsing News ID, Please check news id is valid Number",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news/alkdjhaqwd",
			handler:        NewHTTP(nil, &MockNewsDataUC{}),
		},
		{
			name: "Failed - Usecase Return Error",
			mustReturn: response.ErrorResponse{
				Success: false,
				Message: "Failed Run Get Single News",
				Type:    0,
				Data:    nil,
			},
			mustReturnCode: http.StatusBadRequest,
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				deleteSingleNews: deleteSingleNews{err: fmt.Errorf("Adwde")},
			}),
		},
		{
			name:           "Success",
			mustReturn:     "",
			mustReturnCode: http.StatusNoContent,
			url:            "/news/1",
			handler: NewHTTP(nil, &MockNewsDataUC{
				createSingleNews: createSingleNews{err: nil},
			}),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", tc.url, nil)

			router := gin.Default()
			router.DELETE("/news/:id", tc.handler.HandleDeleteSingleNews)
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
					FunctionName: "Test_HandleDeleteSingleNews",
					Description:  "Testcase run Test_HandleCreateSingleNews",
					Trace:        fmt.Sprintf("got %s, expected %v, code %v", w.Body.String(), string(jsonMustResponse), w.Code),
				}.Error())
			}
		})
	}
}
