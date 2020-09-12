package presenter

import "github.com/garaujo/golang-clean-architecture/domain/model"

// UserPresenter interface
type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
