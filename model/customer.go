package model

import (
	"errors"
	"time"
)

var (
	ErrCustomerDniUK = errors.New("customers_dni_key")
)

type Customer struct {
	Id        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Dni       string    `json:"dni"`
	BirthDate time.Time `json:"birth_date"`
	Gender    Gender    `json:"gender"`
	Email     string    `json:"email"`
	IsStaff   bool      `json:"is_staff"`
	Picture   string    `json:"picture"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Customers []Customer

func (c Customer) HasID() bool { return c.Id > 0 }

func (c Customer) GetAge() uint8 {
	currentDate := time.Now()
	age := uint8(currentDate.Year() - c.BirthDate.Year())

	hasDaysUntilBirthday := c.BirthDate.Month() == currentDate.Month() && c.BirthDate.Day() > currentDate.Day()

	if c.BirthDate.Month() > currentDate.Month() || hasDaysUntilBirthday {
		age--
	}

	return age
}

func (c Customers) GetUniqueIds() []uint {
	uniqueIds := make([]uint, 0)

	ids := make(map[uint]bool)

	for _, customer := range c {
		if _, value := ids[customer.Id]; !value {
			ids[customer.Id] = true
			uniqueIds = append(uniqueIds, customer.Id)
		}
	}

	return uniqueIds
}

func (c Customers) MakeMapById() map[uint]Customer {
	m := make(map[uint]Customer)
	for _, customer := range c {
		m[customer.Id] = customer
	}

	return m
}
