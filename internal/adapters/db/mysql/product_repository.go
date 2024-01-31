package mysql

import (
	"database/sql"

	"github.com/uiansol/product-follow-up/internal/application/entities"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Save(user entities.Product) (string, error) {
	return "TODO: implement", nil
}
