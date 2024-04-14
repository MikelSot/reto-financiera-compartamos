package citycustomer

import (
	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type UseCase interface {
	Create(m model.CityCustomer) error
	Update(m model.CityCustomer) error
	Delete(ID uint) error

	GetByCustomerIDAndCityID(customerID, cityID uint) (model.CityCustomer, error)
	GetAllByCustomerIDs(customerIDs []uint) (model.CityCustomers, error)
}

type Storage interface {
	Create(m model.CityCustomer) error
	Update(m model.CityCustomer) error
	Delete(ID uint) error

	GetWhere(specification repository.FieldsSpecification) (model.CityCustomer, error)
	GetAllWhere(specification repository.FieldsSpecification) (model.CityCustomers, error)
}
