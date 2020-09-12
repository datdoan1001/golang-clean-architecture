package presenter

import "github.com/garaujo/golang-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
