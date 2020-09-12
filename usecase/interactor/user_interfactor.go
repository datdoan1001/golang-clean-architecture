package interactor

import (
	"github.com/garaujo/golang-clean-architecture/domain/model"
	"github.com/garaujo/golang-clean-architecture/usecase/presenter"
	"github.com/garaujo/golang-clean-architecture/usecase/repository"
)

// userInteractor composes UserRepository and UserPresenter interfaces
type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

// UserInteractor interface
type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

// NewUserInteractor function
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

// Get method for userInteractor struct - Implementation of UserInteractor interface
func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
