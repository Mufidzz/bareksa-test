package usecase

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_CreateSingleNews(t *testing.T) {
	type inputParam struct {
		newNews presentation.CreateNewsRequest
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					createBulkNews: createBulkNews{err: fmt.Errorf("AXDCZ")},
				},
			},
			in: inputParam{newNews: presentation.CreateNewsRequest{
				Title:   "A",
				Content: "B",
				Status:  1,
			}},
			mustErr: true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					createBulkNews: createBulkNews{
						insertedID: []int{123},
						err:        nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{flushAll: flushAll{
					nil,
				}},
			},
			in: inputParam{newNews: presentation.CreateNewsRequest{
				Title:   "A",
				Content: "B",
				Status:  1,
			}},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			err := uc.CreateSingleNews(tc.in.newNews)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_UpdateSingleNews(t *testing.T) {
	type inputParam struct {
		updatedNews presentation.UpdateNewsRequest
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					updateBulkNews: updateBulkNews{err: fmt.Errorf("AXDCZ")},
				},
			},
			in: inputParam{updatedNews: presentation.UpdateNewsRequest{
				ID:      1,
				Title:   "A",
				Content: "B",
				Status:  1,
			}},
			mustErr: true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					updateBulkNews: updateBulkNews{
						updatedID: []int{123},
						err:       nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{flushAll: flushAll{nil}},
			},
			in: inputParam{updatedNews: presentation.UpdateNewsRequest{
				ID:      1,
				Title:   "A",
				Content: "B",
				Status:  1,
			}},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			err := uc.UpdateSingleNews(tc.in.updatedNews)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_UpdateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_DeleteSingleNews(t *testing.T) {
	type inputParam struct {
		newsID int
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					deleteBulkNews: deleteBulkNews{err: fmt.Errorf("AXDCZ")},
				},
			},
			in:      inputParam{newsID: 123},
			mustErr: true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					deleteBulkNews: deleteBulkNews{
						deletedID: []int{123},
						err:       nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{flushAll: flushAll{
					nil,
				}},
			},
			in:      inputParam{newsID: 124312},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			err := uc.DeleteSingleNews(tc.in.newsID)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_DeleteSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_GetSingleNews(t *testing.T) {
	type inputParam struct {
		newsID int
	}

	now := time.Now()

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustReturn presentation.GetNewsResponse
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{err: fmt.Errorf("AXDCZ")},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{err: fmt.Errorf("any")},
					saveObject: saveObject{err: fmt.Errorf("any")},
				},
			},
			in:         inputParam{newsID: 123},
			mustReturn: presentation.GetNewsResponse{},
			mustErr:    true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{
						res: []presentation.GetNewsResponse{
							{
								ID:         1,
								CreatedAt:  now,
								UpdatedAt:  now,
								Title:      "ABCDE",
								Content:    "EFGH",
								TopicsName: "AAA",
								TagsName:   "ADWD",
								Status:     1,
							},
						},
						err: nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{fmt.Errorf("any")},
					saveObject: saveObject{nil},
				},
			},
			in: inputParam{newsID: 124312},
			mustReturn: presentation.GetNewsResponse{
				ID:         1,
				CreatedAt:  now,
				UpdatedAt:  now,
				Title:      "ABCDE",
				Content:    "EFGH",
				TopicsName: "AAA",
				TagsName:   "ADWD",
				Status:     1,
			},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}
			c := gin.Context{
				Request: &http.Request{RequestURI: "ASD"},
			}
			log.Println(c.Request.RequestURI)
			got, err := uc.GetSingleNews(&c, tc.in.newsID)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(tc.mustReturn, got) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_GetSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_GetNews(t *testing.T) {
	type inputParam struct {
		paginationString string
		filterString     string
	}

	now := time.Now()

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustReturn []presentation.GetNewsResponse
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{err: fmt.Errorf("AXDCZ")},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{err: fmt.Errorf("any")},
					saveObject: saveObject{err: fmt.Errorf("any")},
				},
			},
			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "",
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Failed - No Pagination & No Filter",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{
						res: []presentation.GetNewsResponse{
							{
								ID:         1,
								CreatedAt:  now,
								UpdatedAt:  now,
								Title:      "ABCDE",
								Content:    "EFGH",
								TopicsName: "AAA",
								TagsName:   "ADWD",
								Status:     1,
							},
						},
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{err: fmt.Errorf("any")},
					saveObject: saveObject{err: fmt.Errorf("any")},
				},
			},
			in: inputParam{
				paginationString: "",
				filterString:     "",
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{
						res: []presentation.GetNewsResponse{
							{
								ID:         1,
								CreatedAt:  now,
								UpdatedAt:  now,
								Title:      "ABCDE",
								Content:    "EFGH",
								TopicsName: "AAA",
								TagsName:   "ADWD",
								Status:     1,
							},
						},
						err: nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{err: fmt.Errorf("any")},
					saveObject: saveObject{err: fmt.Errorf("any")},
				},
			},
			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "",
			},
			mustReturn: []presentation.GetNewsResponse{
				{
					ID:         1,
					CreatedAt:  now,
					UpdatedAt:  now,
					Title:      "ABCDE",
					Content:    "EFGH",
					TopicsName: "AAA",
					TagsName:   "ADWD",
					Status:     1,
				},
			},
			mustErr: false,
		},
		{
			name: "Success - With Filter String",
			repository: &Repositories{
				NewsDataRepository: &MockNewsRepository{
					getBulkNews: getBulkNews{
						res: []presentation.GetNewsResponse{
							{
								ID:         1,
								CreatedAt:  now,
								UpdatedAt:  now,
								Title:      "ABCDE",
								Content:    "EFGH",
								TopicsName: "AAA",
								TagsName:   "ADWD",
								Status:     1,
							},
						},
						err: nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					getObject:  getObject{err: fmt.Errorf("any")},
					saveObject: saveObject{err: fmt.Errorf("any")},
				},
			},
			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "eyJ0aXRsZSIgOiAiYXZkZSJ9",
			},
			mustReturn: []presentation.GetNewsResponse{
				{
					ID:         1,
					CreatedAt:  now,
					UpdatedAt:  now,
					Title:      "ABCDE",
					Content:    "EFGH",
					TopicsName: "AAA",
					TagsName:   "ADWD",
					Status:     1,
				},
			},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			c := gin.Context{
				Request: &http.Request{RequestURI: "Aasda"},
			}
			got, err := uc.GetNews(&c, tc.in.paginationString, tc.in.filterString)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(tc.mustReturn, got) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_GetNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_AssignNewsWithNewsTopic(t *testing.T) {
	testcases := []struct {
		name       string
		repository *Repositories
		in         presentation.CreateNewsTopicsAssoc
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				AssignNewsAssocRepository: &MockAssignNewsAssocRepository{
					createBulkNewsTopicsAssoc: createBulkNewsTopicsAssoc{
						err: fmt.Errorf("awd"),
					},
				},
			},
			in: presentation.CreateNewsTopicsAssoc{
				NewsID:       1,
				NewsTopicsID: []int{1, 2, 3},
			},
			mustErr: true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				AssignNewsAssocRepository: &MockAssignNewsAssocRepository{
					createBulkNewsTopicsAssoc: createBulkNewsTopicsAssoc{
						err: nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					flushAll: flushAll{nil},
				},
			},
			in: presentation.CreateNewsTopicsAssoc{
				NewsID:       1,
				NewsTopicsID: []int{1, 2, 3},
			},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			err := uc.AssignNewsWithNewsTopic(tc.in)

			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_AssignNewsWithNewsTags(t *testing.T) {
	testcases := []struct {
		name       string
		repository *Repositories
		in         presentation.CreateNewsTagsAssoc
		mustErr    bool
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				AssignNewsAssocRepository: &MockAssignNewsAssocRepository{
					createBulkNewsTagsAssoc: createBulkNewsTagsAssoc{
						err: fmt.Errorf("awd"),
					},
				},
			},
			in: presentation.CreateNewsTagsAssoc{
				NewsID:    1,
				NewsTagID: []int{1, 2, 3},
			},
			mustErr: true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				AssignNewsAssocRepository: &MockAssignNewsAssocRepository{
					createBulkNewsTagsAssoc: createBulkNewsTagsAssoc{
						err: nil,
					},
				},
				NewsRedisRepository: &MockNewsRedisRepository{
					flushAll: flushAll{nil},
				},
			},
			in: presentation.CreateNewsTagsAssoc{
				NewsID:    1,
				NewsTagID: []int{1, 2, 3},
			},
			mustErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			err := uc.AssignNewsWithNewsTag(tc.in)
			if (tc.mustErr && err == nil) || (!tc.mustErr && err != nil) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("mustErr %v, err %v", tc.mustErr, err),
				}.Error())
			}
		})
	}
}
