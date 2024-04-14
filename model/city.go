package model

import "time"

type City struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	PostalCode string    `json:"postal_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type Cities []City

func (c City) HasID() bool { return c.Id > 0 }

func (c City) HasName() bool { return c.Name != "" }

func (c Cities) MakeMapById() map[uint]City {
	m := make(map[uint]City)
	for _, city := range c {
		m[city.Id] = city
	}

	return m
}
