package controller

import (
	"net/http"

	"github.com/garaujo/golang-clean-architecture/domain/model"
	"github.com/garaujo/golang-clean-architecture/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

// UserController interface
type UserController interface {
	GetUsers(c Context) error
}

// NewUserController function
func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
