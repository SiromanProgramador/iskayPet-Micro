package repository

import (
	"iskayPetMicro/model"
)

type RepositoryInterface interface {
	Create(teacher *model.Pet) error
	GetAll(filter model.QueryFilters) ([]model.Pet, error)
	Count(filter model.QueryFilters) (int, error)
}
