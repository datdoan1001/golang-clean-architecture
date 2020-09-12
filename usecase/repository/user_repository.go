package repository

import "github.com/garaujo/golang-clean-architecture/domain/model"

// UserRepository interface
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
