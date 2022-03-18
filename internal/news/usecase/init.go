package usecase

type Repositories struct {
	NewsDataRepository
	NewsTopicDataRepository
}

type Usecase struct {
	repositories *Repositories
}

func New(repositories *Repositories) *Usecase {
	return &Usecase{
		repositories: repositories,
	}
}
