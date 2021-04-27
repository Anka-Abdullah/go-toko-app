package app

import "github.com/Anka-Abdullah/Go-toko-1/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
	}

}
