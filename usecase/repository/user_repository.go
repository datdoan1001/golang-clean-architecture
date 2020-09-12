package repository

import "github.com/garaujo/golang-clean-architecture/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
