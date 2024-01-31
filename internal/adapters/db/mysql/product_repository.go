package mysql

import "github.com/uiansol/product-follow-up/internal/application/entities"

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) Save(user entities.Product) (string, error) {
	return "TODO: implement", nil
}
