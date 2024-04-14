package citycustomer

import (
	"fmt"

	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type CityCustomer struct {
	storage Storage
}

func New(storage Storage) CityCustomer {
	return CityCustomer{storage}
}

func (c CityCustomer) Create(m model.CityCustomer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("citycustomer: %w", err)
	}

	if err := c.storage.Create(m); err != nil {
		return err
	}

	return nil
}

func (c CityCustomer) Update(m model.CityCustomer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("citycustomer: %w", err)
	}

	if !m.HasID() {
		return model.ErrInvalidID
	}

	if err := c.storage.Update(m); err != nil {
		return err
	}

	return nil
}

func (c CityCustomer) Delete(ID uint) error {
	err := c.storage.Delete(ID)
	if err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c CityCustomer) GetByCustomerIDAndCityID(customerID, cityID uint) (model.CityCustomer, error) {
	if customerID == 0 || cityID == 0 {
		return model.CityCustomer{}, model.ErrInvalidID
	}

	customerCity, err := c.storage.GetWhere(
		repository.FieldsSpecification{
			Filters: repository.Fields{
				{Name: "customer_id", Value: customerID},
				{Name: "city_id", Value: cityID},
			},
		},
	)
	if err != nil {
		return model.CityCustomer{}, fmt.Errorf("citycustomer: %w", err)
	}

	return customerCity, nil
}

func (c CityCustomer) GetAllByCustomerIDs(customerIDs []uint) (model.CityCustomers, error) {
	ms, err := c.storage.GetAllWhere(repository.FieldsSpecification{Filters: repository.Fields{
		repository.Field{Name: "customer_id", Value: customerIDs, Operator: repository.In},
	}})
	if err != nil {
		return nil, fmt.Errorf("citycustomer.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (c CityCustomer) errorConstraint(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
