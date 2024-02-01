package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string
	Name      string
	Comments  string
	Link      string
	Price     float64
	PriceDate time.Time
}

func NewProduct(name, comments, link string, price float64) (Product, error) {

	if name == "" {
		return Product{}, errors.New("product name is required")
	}

	if link == "" {
		return Product{}, errors.New("product link is required")
	}

	if price <= 0 {
		return Product{}, errors.New("product price is required")
	}

	product := Product{
		ID:        uuid.New().String(),
		Name:      name,
		Comments:  comments,
		Link:      link,
		Price:     price,
		PriceDate: time.Now(),
	}

	return product, nil
}

func (p *Product) UpdatePriceDate() {
	p.PriceDate = time.Now()
}
