package citycustomer

import (
	"context"
	"database/sql"

	"github.com/MikelSot/repository"
	"github.com/jackc/pgx/v5"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

var _fieldInserts = []string{
	"customer_id",
	"city_id",
}

var _fieldsSelect = []string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
}

const _table = "city_customers"

var (
	_psqlInsert = repository.BuildSQLInsertNoID(_table, _fieldInserts)
	_psqlUpdate = repository.BuildSQLUpdateByID(_table, _fieldInserts)
	_psqlDelete = "DELETE FROM " + _table + " WHERE id = $1"

	_psqlGetAll = repository.BuildSQLSelectFields(_table, append(_fieldInserts, _fieldsSelect...))
)

type CityCustomer struct {
	db model.PgxPool
}

func New(db model.PgxPool) CityCustomer {
	return CityCustomer{db: db}
}

func (c CityCustomer) Create(m model.CityCustomer) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlInsert,
		m.CustomerID,
		m.CityID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c CityCustomer) Update(m model.CityCustomer) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlUpdate,
		m.CustomerID,
		m.CityID,
		m.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c CityCustomer) Delete(ID uint) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c CityCustomer) GetWhere(specification repository.FieldsSpecification) (model.CityCustomer, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryAndArgs(_psqlGetAll, specification)

	m, err := c.scanRow(c.db.QueryRow(ctx, query, args...))
	if err != nil {
		return model.CityCustomer{}, err
	}

	return m, nil
}

func (c CityCustomer) GetAllWhere(specification repository.FieldsSpecification) (model.CityCustomers, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAll, specification)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	ms := model.CityCustomers{}
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (c CityCustomer) scanRow(row pgx.Row) (model.CityCustomer, error) {
	m := model.CityCustomer{}

	updatedAtNull := sql.NullTime{}
	deletedAtNull := sql.NullTime{}

	err := row.Scan(
		&m.CustomerID,
		&m.CityID,
		&m.Id,
		&m.CreatedAt,
		&updatedAtNull,
		&deletedAtNull,
	)
	if err != nil {
		return model.CityCustomer{}, err
	}

	m.UpdatedAt = updatedAtNull.Time
	m.DeletedAt = deletedAtNull.Time

	return m, nil
}
