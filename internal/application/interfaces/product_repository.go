package interfaces

import "github.com/uiansol/product-follow-up/internal/application/entities"

type IProductRepository interface {
	Save(user *entities.Product) (string, error)
	Read(id string) (*entities.Product, error)
	Update(user *entities.Product) error
	Delete(id string) error
}
