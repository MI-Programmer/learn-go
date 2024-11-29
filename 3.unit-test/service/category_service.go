package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

var categoryNotFound = errors.New("Category Not Found")

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (s CategoryService) Get(id string) (*entity.Category, error) {
	category := s.Repository.FindById(id)

	if category == nil {
		return nil, categoryNotFound
	} else {
		return category, nil
	}
}
