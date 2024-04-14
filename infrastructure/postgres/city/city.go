package city

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MikelSot/repository"
	"github.com/jackc/pgx/v5"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

var _fieldInserts = []string{
	"name",
	"postal_code",
}

var _fieldsSelect = []string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
}

const _table = "cities"

var (
	_psqlInsert = repository.BuildSQLInsertNoID(_table, _fieldInserts)
	_psqlUpdate = repository.BuildSQLUpdateByID(_table, _fieldInserts)
	_psqlDelete = "DELETE FROM " + _table + " WHERE id = $1"

	_psqlGetAll = repository.BuildSQLSelectFields(_table, append(_fieldInserts, _fieldsSelect...))
)

type City struct {
	db model.PgxPool
}

func New(db model.PgxPool) City {
	return City{db: db}
}

func (c City) Create(m model.City) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlInsert,
		m.Name,
		repository.StringToNull(m.PostalCode),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c City) Update(m model.City) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlUpdate,
		m.Name,
		repository.StringToNull(m.PostalCode),
		m.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c City) Delete(ID uint) error {
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

func (c City) GetWhere(specification repository.FieldsSpecification) (model.City, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryAndArgs(_psqlGetAll, specification)

	m, err := c.scanRow(c.db.QueryRow(ctx, query, args...))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.City{}, nil
	}
	if err != nil {
		return model.City{}, err
	}

	return m, nil
}

func (c City) GetAllWhere(specification repository.FieldsSpecification) (model.Cities, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAll, specification)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	ms := model.Cities{}
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (c City) scanRow(row pgx.Row) (model.City, error) {
	m := model.City{}

	postalCodeNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	deletedAtNull := sql.NullTime{}

	err := row.Scan(
		&m.Name,
		&postalCodeNull,
		&m.Id,
		&m.CreatedAt,
		&updatedAtNull,
		&deletedAtNull,
	)
	if err != nil {
		return model.City{}, err
	}

	m.PostalCode = postalCodeNull.String
	m.UpdatedAt = updatedAtNull.Time
	m.DeletedAt = deletedAtNull.Time

	return m, nil
}
