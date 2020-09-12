package registry

import (
	"github.com/garaujo/golang-clean-architecture/interface/controller"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

// Registry interface
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry function registers a DB
func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

// NewAppControler method for registry struct - Implementation of Registry interface
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
	}
}
