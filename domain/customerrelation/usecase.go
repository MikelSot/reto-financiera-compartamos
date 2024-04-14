package customerrelation

import (
	"fmt"
	"net/http"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type CustomerRelation struct {
	customer     CustomerUseCase
	city         CityUseCase
	cityCustomer CityCustomerUseCase
}

func New(c CustomerUseCase, ci CityUseCase, cc CityCustomerUseCase) CustomerRelation {
	return CustomerRelation{c, ci, cc}
}

func (c CustomerRelation) Create(m model.CustomerRelation) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("customerrelation: %w", err)
	}

	if err := c.customer.Create(m.Customer); err != nil {
		return err
	}

	city, err := c.city.GetByID(m.City.Id)
	if err != nil {
		return err
	}

	if !city.HasID() {
		customErr := model.NewError()
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "city_id",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("city_id: %s invalid", m.City.Id),
		})
		customErr.SetAPIMessage("¡Upps! Error la ciudad no existe")

		return customErr
	}

	customerCity := model.CityCustomer{
		CustomerID: m.Customer.Id,
		CityID:     m.City.Id,
	}

	if err := c.cityCustomer.Create(customerCity); err != nil {
		return err
	}

	return nil
}

func (c CustomerRelation) Update(m model.CustomerRelation) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("customerrelation: %w", err)
	}

	if err := c.customer.Update(m.Customer); err != nil {
		return err
	}

	city, err := c.city.GetByID(m.City.Id)
	if err != nil {
		return err
	}

	if !city.HasID() {
		customErr := model.NewError()
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "city_id",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("city_id: %s invalid", m.City.Id),
		})
		customErr.SetAPIMessage("¡Upps! Error la ciudad no existe")

		return customErr
	}

	customerCity := model.CityCustomer{
		CustomerID: m.Customer.Id,
		CityID:     m.City.Id,
	}

	if err := c.cityCustomer.Update(customerCity); err != nil {
		return err
	}

	return nil
}

func (c CustomerRelation) Delete(customerID uint, cityID uint) error {
	if err := c.customer.CreateDelete(customerID); err != nil {
		return err
	}

	cityCustomer, err := c.cityCustomer.GetByCustomerIDAndCityID(customerID, cityID)
	if err != nil {
		return err
	}
	if !cityCustomer.HasID() {
		return model.ErrInvalidID
	}

	if err := c.cityCustomer.Delete(cityCustomer.Id); err != nil {
		return err
	}

	return nil
}

func (c CustomerRelation) GetAllCustomers() (model.CustomerRelations, error) {
	customers, err := c.customer.GetAll()
	if err != nil {
		return model.CustomerRelations{}, fmt.Errorf("customerrelation: %w", err)
	}

	cityCustomers, err := c.cityCustomer.GetAllByCustomerIDs(customers.GetUniqueIds())
	if err != nil {
		return model.CustomerRelations{}, fmt.Errorf("customerrelation: %w", err)
	}

	cities, err := c.city.GetAllByIds(cityCustomers.GetCityIDs())
	if err != nil {
		return model.CustomerRelations{}, fmt.Errorf("customerrelation: %w", err)
	}

	customerMap := customers.MakeMapById()
	cityMap := cities.MakeMapById()

	relations := model.CustomerRelations{}
	for _, cityCustomer := range cityCustomers {
		customer, ok := customerMap[cityCustomer.CustomerID]
		if !ok {
			continue
		}

		city, _ := cityMap[cityCustomer.CityID]

		relation := model.CustomerRelation{
			Customer: customer,
			City:     city,
		}

		relations = append(relations, relation)
	}

	return relations, nil
}
