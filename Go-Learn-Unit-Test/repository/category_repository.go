package repository

import "Go-learn-Unit-Test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
