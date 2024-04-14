package city

import (
	"fmt"
	"net/http"

	"github.com/MikelSot/repository"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type City struct {
	storage Storage
}

func New(s Storage) City {
	return City{s}
}

func (c City) Create(m model.City) error {
	if err := model.ValidateStructNil(m); err != nil {
		return err
	}

	err := c.validateRequest(m)
	if err != nil {
		return err
	}

	if err := c.storage.Create(m); err != nil {
		return err
	}

	return nil
}

func (c City) Update(m model.City) error {
	if err := model.ValidateStructNil(m); err != nil {
		return err
	}

	if !m.HasID() {
		return model.ErrInvalidID
	}

	err := c.validateRequest(m)
	if err != nil {
		return err
	}

	if err := c.storage.Update(m); err != nil {
		return err
	}

	return nil
}

func (c City) Delete(ID uint) error {
	err := c.storage.Delete(ID)
	if err != nil {
		return c.errorConstraint(err)
	}

	return nil
}

func (c City) GetByID(ID uint) (model.City, error) {
	if ID == 0 {
		return model.City{}, fmt.Errorf("city: %w", model.ErrInvalidID)
	}

	city, err := c.storage.GetWhere(repository.FieldsSpecification{Filters: repository.Fields{{Name: "id", Value: ID}}})
	if err != nil {
		return model.City{}, fmt.Errorf("city: %w", err)
	}

	return city, nil
}

func (c City) GetAllByIds(IDs []uint) (model.Cities, error) {
	ms, err := c.storage.GetAllWhere(repository.FieldsSpecification{Filters: repository.Fields{
		repository.Field{Name: "id", Value: IDs, Operator: repository.In},
	}})
	if err != nil {
		return nil, fmt.Errorf("city.storage.GetAllWhere(): %w", err)
	}

	return ms, nil

}

func (c City) GetAll() (model.Cities, error) {
	ms, err := c.storage.GetAllWhere(repository.FieldsSpecification{Filters: repository.Fields{
		repository.Field{Name: "deleted_at", Operator: repository.IsNull},
	}})
	if err != nil {
		return nil, fmt.Errorf("city.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (c City) validateRequest(m model.City) error {
	if !m.HasName() {
		customErr := model.NewError()
		customErr.SetStatusHTTP(http.StatusUnprocessableEntity)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "name",
			Issue:       model.IssueViolatedValidation,
			Description: fmt.Sprintf("name: %s invalid", m.Name),
		})
		customErr.SetAPIMessage("¡Upps! Error el nombre de la ciudad no debe estar vacío")

		return customErr
	}

	return nil
}

func (c City) errorConstraint(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
