package usecase

type Repositories struct {
	NewsDataRepository
	NewsTopicDataRepository
	NewsTagDataRepository
}

type Usecase struct {
	repositories *Repositories
}

func New(repositories *Repositories) *Usecase {
	return &Usecase{
		repositories: repositories,
	}
}
