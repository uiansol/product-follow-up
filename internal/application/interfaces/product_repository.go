package interfaces

import "github.com/uiansol/product-follow-up/internal/application/entities"

type IProductRepository interface {
	Save(user entities.Product) (string, error)
}
