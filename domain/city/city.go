package city

import (
	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type UseCase interface {
	Create(m model.City) error
	Update(m model.City) error
	Delete(ID uint) error

	GetByID(ID uint) (model.City, error)
	GetAllByIds(IDs []uint) (model.Cities, error)
	GetAll() (model.Cities, error)
}

type Storage interface {
	Create(m model.City) error
	Update(m model.City) error
	Delete(ID uint) error

	GetWhere(specification repository.FieldsSpecification) (model.City, error)
	GetAllWhere(specification repository.FieldsSpecification) (model.Cities, error)
}
