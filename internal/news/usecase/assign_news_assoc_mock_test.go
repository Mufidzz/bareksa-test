package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type MockAssignNewsAssocRepository struct {
	createBulkNewsTopicsAssoc createBulkNewsTopicsAssoc
	createBulkNewsTagsAssoc   createBulkNewsTagsAssoc
	cleanNewsTopicsAssoc      cleanNewsTopicsAssoc
	cleanNewsTagsAssoc        cleanNewsTagsAssoc
}

type createBulkNewsTopicsAssoc struct {
	err error
}

type createBulkNewsTagsAssoc struct {
	err error
}

type cleanNewsTopicsAssoc struct {
	err error
}

type cleanNewsTagsAssoc struct {
	err error
}

func (manar *MockAssignNewsAssocRepository) CreateBulkNewsTopicsAssoc(in []presentation.CreateNewsTopicsAssoc) (err error) {
	return manar.createBulkNewsTopicsAssoc.err
}
func (manar *MockAssignNewsAssocRepository) CreateBulkNewsTagsAssoc(in []presentation.CreateNewsTagsAssoc) (err error) {
	return manar.createBulkNewsTagsAssoc.err
}

func (manar *MockAssignNewsAssocRepository) CleanNewsTopicsAssoc(newsID []int) (err error) {
	return manar.cleanNewsTopicsAssoc.err
}
func (manar *MockAssignNewsAssocRepository) CleanNewsTagAssoc(newsID []int) (err error) {
	return manar.cleanNewsTagsAssoc.err
}
