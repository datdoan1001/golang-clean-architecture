package presenter

import "github.com/garaujo/golang-clean-architecture/domain/model"

type userPresenter struct {
}

// UserPresenter interface
type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

// NewUserPresenter function
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

// ResponseUsers method for userPresenter struct - Implementation of UserPresenter interface
func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
