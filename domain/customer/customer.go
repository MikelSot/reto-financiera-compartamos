package customer

import (
	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type UseCase interface {
	Create(m model.Customer) error
	Update(m model.Customer) error
	CreateDelete(ID uint) error
	Delete(ID uint) error

	GetByID(ID uint) (model.Customer, error)
	GetAll() (model.Customers, error)
}

type Storage interface {
	Create(m model.Customer) error
	Update(m model.Customer) error
	Delete(ID uint) error

	GetWhere(specification repository.FieldsSpecification) (model.Customer, error)
	GetAllWhere(specification repository.FieldsSpecification) (model.Customers, error)
}
