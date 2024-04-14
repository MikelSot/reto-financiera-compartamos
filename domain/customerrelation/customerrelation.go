package customerrelation

import "github.com/MikelSot/reto-financiera-compartamos/model"

type UseCase interface {
	Create(m model.CustomerRelation) error
	Update(m model.CustomerRelation) error
	Delete(customerID, cityID uint) error

	GetAllCustomers() (model.CustomerRelations, error)
}

type CustomerUseCase interface {
	Create(m model.Customer) error
	Update(m model.Customer) error
	CreateDelete(ID uint) error
	Delete(ID uint) error

	GetAll() (model.Customers, error)
}

type CityUseCase interface {
	GetByID(ID uint) (model.City, error)
	GetAllByIds(IDs []uint) (model.Cities, error)
}

type CityCustomerUseCase interface {
	Create(m model.CityCustomer) error
	Update(m model.CityCustomer) error
	Delete(ID uint) error

	GetByCustomerIDAndCityID(customerID, cityID uint) (model.CityCustomer, error)
	GetAllByCustomerIDs(customerIDs []uint) (model.CityCustomers, error)
}
