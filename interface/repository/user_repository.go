package repository

import (
	"github.com/garaujo/golang-clean-architecture/domain/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// UserRepository interface
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

// NewUserRepository function
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// FindAll method for userRepository struct - Implementation of UserRepository interface
func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
