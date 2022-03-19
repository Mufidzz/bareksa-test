package usecase

import (
	"fmt"
	"github.com/Mufidzz/bareksa-test/pkg/response"
	"github.com/Mufidzz/bareksa-test/presentation"
	"reflect"
	"testing"
)

func Test_CreateTags(t *testing.T) {
	type inputParam struct {
		in []presentation.CreateNewsTagsRequest
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
		mustReturn []int
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					createBulkNewsTags: createBulkNewsTags{
						err: fmt.Errorf("ASD"),
					},
				},
			},
			in: inputParam{
				in: []presentation.CreateNewsTagsRequest{
					{
						Name: "AAA",
					},
					{
						Name: "BBB",
					},
				},
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Repo return no error",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					createBulkNewsTags: createBulkNewsTags{
						insertedID: []int{1, 2},
					},
				},
			},
			in: inputParam{
				in: []presentation.CreateNewsTagsRequest{
					{
						Name: "AAA",
					},
					{
						Name: "BBB",
					},
				},
			},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			got, err := uc.CreateNewsTags(tc.in.in)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(got, tc.mustReturn) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_DeleteNewsTags(t *testing.T) {
	type inputParam struct {
		in []int
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
		mustReturn []int
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					deleteBulkNewsTags: deleteBulkNewsTags{
						err: fmt.Errorf("ASD"),
					},
				},
			},
			in: inputParam{
				in: []int{1, 2, 3, 4},
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Repo return no error, all deleted",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					deleteBulkNewsTags: deleteBulkNewsTags{
						deletedID: []int{1, 2, 3, 4},
					},
				},
			},
			in: inputParam{
				in: []int{1, 2, 3, 4},
			},
			mustReturn: []int{1, 2, 3, 4},
			mustErr:    false,
		},
		{
			name: "Success - Repo return no error, partially deleted",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					deleteBulkNewsTags: deleteBulkNewsTags{
						deletedID: []int{1, 2},
					},
				},
			},
			in: inputParam{
				in: []int{1, 2, 3, 4},
			},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			got, err := uc.DeleteNewsTags(tc.in.in)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(got, tc.mustReturn) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_UpdateNewsTags(t *testing.T) {
	type inputParam struct {
		in []presentation.UpdateNewsTagsRequest
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
		mustReturn []int
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					updateBulkNewsTags: updateBulkNewsTags{
						err: fmt.Errorf("ASD"),
					},
				},
			},
			in: inputParam{
				in: []presentation.UpdateNewsTagsRequest{
					{
						ID:   1,
						Name: "A",
					},
					{
						ID:   2,
						Name: "B",
					},
					{
						ID:   3,
						Name: "C",
					}, {
						ID:   4,
						Name: "D",
					},
				},
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - Repo return no error, all updated",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					updateBulkNewsTags: updateBulkNewsTags{
						updatedID: []int{1, 2, 3, 4},
					},
				},
			},
			in: inputParam{
				in: []presentation.UpdateNewsTagsRequest{
					{
						ID:   1,
						Name: "A",
					},
					{
						ID:   2,
						Name: "B",
					},
					{
						ID:   3,
						Name: "C",
					}, {
						ID:   4,
						Name: "D",
					},
				},
			},
			mustReturn: []int{1, 2, 3, 4},
			mustErr:    false,
		},
		{
			name: "Success - Repo return no error, partially updated",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					updateBulkNewsTags: updateBulkNewsTags{
						updatedID: []int{1, 2},
					},
				},
			},
			in: inputParam{
				in: []presentation.UpdateNewsTagsRequest{
					{
						ID:   1,
						Name: "A",
					},
					{
						ID:   2,
						Name: "B",
					},
					{
						ID:   3,
						Name: "C",
					}, {
						ID:   4,
						Name: "D",
					},
				},
			},
			mustReturn: []int{1, 2},
			mustErr:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			got, err := uc.UpdateNewsTags(tc.in.in)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(got, tc.mustReturn) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}

func Test_GetNewsTags(t *testing.T) {
	type inputParam struct {
		paginationString string
		filterString     string
	}

	testcases := []struct {
		name       string
		repository *Repositories
		in         inputParam
		mustErr    bool
		mustReturn []presentation.GetNewsTagsResponse
	}{
		{
			name: "Failed - Repo return error",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						err: fmt.Errorf("ASD"),
					},
				},
			},
			in:         inputParam{},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Success - No Pagination, No Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{},
			mustReturn: []presentation.GetNewsTagsResponse{
				{
					ID:   1,
					Name: "A",
				},
				{
					ID:   2,
					Name: "B",
				},
				{
					ID:   3,
					Name: "C",
				},
			},
			mustErr: false,
		},
		{
			name: "Success - With Pagination, No Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
			},
			mustReturn: []presentation.GetNewsTagsResponse{
				{
					ID:   1,
					Name: "A",
				},
				{
					ID:   2,
					Name: "B",
				},
				{
					ID:   3,
					Name: "C",
				},
			},
			mustErr: false,
		},
		{
			name: "Success - No Pagination, With Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				filterString: "eyJuYW1lIjoiQSJ9",
			},
			mustReturn: []presentation.GetNewsTagsResponse{
				{
					ID:   1,
					Name: "A",
				},
				{
					ID:   2,
					Name: "B",
				},
				{
					ID:   3,
					Name: "C",
				},
			},
			mustErr: false,
		},
		{
			name: "Success - With Pagination, With Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "eyJuYW1lIjoiQSJ9",
			},
			mustReturn: []presentation.GetNewsTagsResponse{
				{
					ID:   1,
					Name: "A",
				},
				{
					ID:   2,
					Name: "B",
				},
				{
					ID:   3,
					Name: "C",
				},
			},
			mustErr: false,
		},
		{
			name: "Failed - Not Recognized Pagination, Not Recognized Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				paginationString: "eyJvZ12312mZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "eyJuYW1lI321231joiQSJ9",
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Failed - Not Recognized Pagination, With Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				paginationString: "eyJvZ12312mZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "eyJuYW1lIjoiQSJ9",
			},
			mustReturn: nil,
			mustErr:    true,
		},
		{
			name: "Failed - With Pagination, Not Recognized Filter",
			repository: &Repositories{
				NewsTagDataRepository: &MockNewsTagDataRepository{
					getBulkNewsTags: getBulkNewsTags{
						res: []presentation.GetNewsTagsResponse{
							{
								ID:   1,
								Name: "A",
							},
							{
								ID:   2,
								Name: "B",
							},
							{
								ID:   3,
								Name: "C",
							},
						},
					},
				},
			},

			in: inputParam{
				paginationString: "eyJvZmZzZXQiIDogMSwgImNvdW50IiA6IDd9",
				filterString:     "eyJuYW1lIwdafafda213joiQSJ9",
			},
			mustReturn: nil,
			mustErr:    true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			uc := Usecase{
				repositories: tc.repository,
			}

			got, err := uc.GetNewsTags(tc.in.paginationString, tc.in.filterString)

			if ((tc.mustErr && err == nil) || (!tc.mustErr && err != nil)) || !reflect.DeepEqual(got, tc.mustReturn) {
				tt.Error(response.InternalTestError{
					Name:         tt.Name(),
					FunctionName: "Test_CreateSingleNews",
					Description:  "Testcase run unsuccessfully",
					Trace:        fmt.Sprintf("got %v, expected %v, mustErr %v, err %v", got, tc.mustReturn, tc.mustErr, err),
				}.Error())
			}
		})
	}
}
