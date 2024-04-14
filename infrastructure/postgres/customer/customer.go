package customer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MikelSot/repository"
	"github.com/jackc/pgx/v5"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

var _fieldInserts = []string{
	"first_name",
	"last_name",
	"dni",
	"birth_date",
	"gender",
	"email",
	"is_staff",
	"picture",
	"nickname",
}

var _fieldsSelect = []string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
}

const _table = "customers"

var (
	_psqlInsert          = repository.BuildSQLInsertNoID(_table, _fieldInserts)
	_psqlUpdate          = repository.BuildSQLUpdateByID(_table, _fieldInserts)
	_psqlDelete          = "DELETE FROM " + _table + " WHERE id = $1"
	_psqlUpdateDeletedAt = "UPDATE " + _table + " SET deleted_at = now() WHERE id = $1"

	_psqlGetAll = repository.BuildSQLSelectFields(_table, append(_fieldInserts, _fieldsSelect...))
)

type Customer struct {
	db model.PgxPool
}

func New(db model.PgxPool) Customer {
	return Customer{db: db}
}

func (c Customer) Create(m *model.Customer) error {
	ctx := context.Background()

	err := c.db.QueryRow(
		ctx,
		_psqlInsert,
		m.FirstName,
		m.LastName,
		m.Dni,
		m.BirthDate,
		m.Gender,
		repository.StringToNull(m.Email),
		m.IsStaff,
		repository.StringToNull(m.Picture),
		repository.StringToNull(m.Nickname),
	).Scan(&m.Id, &m.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (c Customer) Update(m model.Customer) error {
	ctx := context.Background()

	_, err := c.db.Exec(
		ctx,
		_psqlUpdate,
		m.FirstName,
		m.LastName,
		m.Dni,
		m.BirthDate,
		m.Gender,
		repository.StringToNull(m.Email),
		m.IsStaff,
		repository.StringToNull(m.Picture),
		repository.StringToNull(m.Nickname),
		m.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c Customer) UpdateDeletedAt(Id uint) error {
	ctx := context.Background()

	fmt.Println(_psqlUpdateDeletedAt)

	_, err := c.db.Exec(
		ctx,
		_psqlUpdateDeletedAt,
		Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c Customer) Delete(ID uint) error {
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

func (c Customer) GetWhere(specification repository.FieldsSpecification) (model.Customer, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryAndArgs(_psqlGetAll, specification)

	m, err := c.scanRow(c.db.QueryRow(ctx, query, args...))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Customer{}, nil
	}
	if err != nil {
		return model.Customer{}, err
	}

	return m, nil
}

func (c Customer) GetAllWhere(specification repository.FieldsSpecification) (model.Customers, error) {
	ctx := context.Background()

	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAll, specification)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	ms := model.Customers{}
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (c Customer) scanRow(row pgx.Row) (model.Customer, error) {
	m := model.Customer{}

	birthDateNull := sql.NullTime{}
	emailNull := sql.NullString{}
	pictureNull := sql.NullString{}
	nicknameNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	deletedAtNull := sql.NullTime{}

	err := row.Scan(
		&m.FirstName,
		&m.LastName,
		&m.Dni,
		&birthDateNull,
		&m.Gender,
		&emailNull,
		&m.IsStaff,
		&pictureNull,
		&nicknameNull,
		&m.Id,
		&m.CreatedAt,
		&updatedAtNull,
		&deletedAtNull,
	)
	if err != nil {
		return model.Customer{}, err
	}

	m.BirthDate = birthDateNull.Time
	m.Email = emailNull.String
	m.Picture = pictureNull.String
	m.Nickname = nicknameNull.String
	m.UpdatedAt = updatedAtNull.Time
	m.DeletedAt = deletedAtNull.Time

	return m, nil
}
