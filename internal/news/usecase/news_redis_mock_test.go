package usecase

type MockNewsRedisRepository struct {
	getObject  getObject
	saveObject saveObject
	flushAll   flushAll
}

type getObject struct {
	err error
}

type saveObject struct {
	err error
}

type flushAll struct {
	err error
}

func (mnrr *MockNewsRedisRepository) GetObject(key string, dest interface{}) error {
	return mnrr.getObject.err
}
func (mnrr *MockNewsRedisRepository) SaveObject(key string, value interface{}) error {
	return mnrr.saveObject.err
}
func (mnrr *MockNewsRedisRepository) FlushAll() error {
	return mnrr.flushAll.err
}
