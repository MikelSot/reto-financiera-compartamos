package customer

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

const (
	_dniRexExp = `^\d{8}$`

	_minimumAge = 18
	_maximumAge = 80
)

type Customer struct {
	storage Storage
}

func New(s Storage) Customer {
	return Customer{s}
}

func (c Customer) Create(m *model.Customer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("customer: %w", err)
	}

	err := c.validateRequest(*m)
	if err != nil {
		return err
	}

	if err := c.validateUniqueDni(*m); err != nil {
		return err
	}

	if err := c.storage.Create(m); err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c Customer) Update(m model.Customer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("customer: %w", err)
	}

	if !m.HasID() {
		return model.ErrInvalidID
	}

	err := c.validateRequest(m)
	if err != nil {
		return err
	}

	customer, err := c.GetByID(m.Id)
	if err != nil {
		return err
	}

	if customer.Dni != m.Dni {
		err = c.validateUniqueDni(m)
		if err != nil {
			return err
		}
	}

	if err := c.storage.Update(m); err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c Customer) updateDeletedAt(Id uint) error {
	if err := c.storage.UpdateDeletedAt(Id); err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c Customer) CreateDelete(ID uint) error {
	err := c.validateDelete(ID)
	if err != nil {
		return err
	}

	if err := c.updateDeletedAt(ID); err != nil {
		return err
	}

	return nil
}

func (c Customer) Delete(ID uint) error {
	err := c.validateDelete(ID)
	if err != nil {
		return err
	}

	err = c.storage.Delete(ID)
	if err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c Customer) GetByID(ID uint) (model.Customer, error) {
	if ID == 0 {
		return model.Customer{}, model.ErrInvalidID
	}

	customer, err := c.storage.GetWhere(repository.FieldsSpecification{Filters: repository.Fields{{Name: "id", Value: ID}}})
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer: %w", err)
	}

	return customer, nil
}

func (c Customer) getByDni(dni string) (model.Customer, error) {
	customer, err := c.storage.GetWhere(repository.FieldsSpecification{Filters: repository.Fields{{Name: "dni", Value: dni}}})
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer: %w", err)
	}

	return customer, nil
}

func (c Customer) GetAll() (model.Customers, error) {
	ms, err := c.storage.GetAllWhere(repository.FieldsSpecification{Filters: repository.Fields{
		repository.Field{Name: "deleted_at", Operator: repository.IsNull},
	}})
	if err != nil {
		return nil, fmt.Errorf("customer.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (c Customer) validateRequest(m model.Customer) error {
	if err := c.validateDni(m.Dni); err != nil {
		return err
	}

	if m.GetAge() <= _minimumAge {
		customErr := model.NewError()
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "birth_date",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("birth_date: %s invalid", m.BirthDate),
		})
		customErr.SetAPIMessage(fmt.Sprintf("¡Upps! Error el cliente debe ser mayor de %d años", _minimumAge))

		return customErr
	}

	return nil
}

func (c Customer) validateDni(dni string) error {
	expReg, err := regexp.Compile(_dniRexExp)

	customErr := model.NewError()
	if err != nil {
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "dni",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("dni: %s invalid", dni),
		})
		customErr.SetAPIMessage("¡Upps! Error formato de dni invalido")

		return customErr
	}

	if !expReg.MatchString(dni) {
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "dni",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("dni: %s invalid", dni),
		})
		customErr.SetAPIMessage("¡Upps! Error el dni ingresado no es valido")

		return customErr
	}

	return nil
}

func (c Customer) validateDelete(ID uint) error {
	customer, err := c.GetByID(ID)
	if err != nil {
		return err
	}

	if customer.GetAge() <= _maximumAge {
		customErr := model.NewError()
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "birth_date",
			Issue:       model.IssueRequestBackendFailed,
			Description: fmt.Sprintf("Age: %s invalid", customer.BirthDate),
		})
		customErr.SetAPIMessage(fmt.Sprintf("¡Upps! Error el cliente debe ser mayor de %d años", _maximumAge))

		return customErr
	}

	return nil
}

func (c Customer) validateUniqueDni(customer model.Customer) error {
	otherCustomer, err := c.getByDni(customer.Dni)
	if err != nil {
		return err
	}

	customErr := model.NewError()
	customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
	customErr.Fields.Add(model.ErrorDetail{
		Field:       "dni",
		Issue:       model.IssueViolatedValidation,
		Description: fmt.Sprintf("dni: %s invalid", customer.Dni),
	})
	customErr.SetAPIMessage("¡Upps! Error el dni ya existe")

	// We validate that the user who has the ID I want to use is not the same
	if customer.HasID() && otherCustomer.HasID() && customer.Id != otherCustomer.Id {
		return customErr
	}

	if otherCustomer.HasID() {
		return customErr
	}

	return nil
}

func (c Customer) errorConstraint(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
