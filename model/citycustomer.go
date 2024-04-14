package model

import "time"

type CityCustomer struct {
	Id         uint      `json:"id"`
	CityID     uint      `json:"city_id"`
	CustomerID uint      `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type CityCustomers []CityCustomer

func (c CityCustomer) HasID() bool { return c.Id > 0 }

func (c CityCustomers) GetCityIDs() []uint {
	ids := make([]uint, 0, len(c))
	for _, v := range c {
		ids = append(ids, v.CityID)
	}
	return ids
}

func (c CityCustomers) MakeMapByCustomerId() map[uint]CityCustomer {
	m := make(map[uint]CityCustomer)
	for _, cityCustomer := range c {
		m[cityCustomer.CustomerID] = cityCustomer
	}

	return m
}
