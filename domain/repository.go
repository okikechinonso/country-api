package domain

import "countries-api/entity"

type DbInterface interface {
	Create(country entity.Country) (*entity.Country, error)
	Find(name string) (*entity.Country, error)
	FindMany() ([]entity.Country, error)
	Update(country entity.Country, id string) (interface{}, error)
}
