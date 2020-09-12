package registry

import (
	"github.com/garaujo/golang-clean-architecture/interface/controller"
	ip "github.com/garaujo/golang-clean-architecture/interface/presenter"
	ir "github.com/garaujo/golang-clean-architecture/interface/repository"
	"github.com/garaujo/golang-clean-architecture/usecase/interactor"
	up "github.com/garaujo/golang-clean-architecture/usecase/presenter"
	ur "github.com/garaujo/golang-clean-architecture/usecase/repository"
)

// NewUserController method for registry struct
func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

// NewUserInteractor method for registry struct
func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

// NewUserRepository method for registry struct
func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

// NewUserPresenter method for registry struct
func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
