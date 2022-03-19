package usecase

import "github.com/Mufidzz/bareksa-test/presentation"

type MockAssignNewsAssocRepository struct {
	createBulkNewsTopicsAssoc createBulkNewsTopicsAssoc
	createBulkNewsTagsAssoc   createBulkNewsTagsAssoc
}

type createBulkNewsTopicsAssoc struct {
	err error
}

type createBulkNewsTagsAssoc struct {
	err error
}

func (manar *MockAssignNewsAssocRepository) CreateBulkNewsTopicsAssoc(in []presentation.CreateNewsTopicsAssoc) (err error) {
	return manar.createBulkNewsTopicsAssoc.err
}
func (manar *MockAssignNewsAssocRepository) CreateBulkNewsTagsAssoc(in []presentation.CreateNewsTagsAssoc) (err error) {
	return manar.createBulkNewsTagsAssoc.err
}
