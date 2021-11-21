package logic

import "account.oscto.icu/repository"

type Logic struct {
	repository *repository.Repository
}

func Register() *Logic {

	return &Logic{
		repository: repository.Register(),
	}
}
