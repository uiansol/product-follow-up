package entities

import (
	"errors"
	"time"
)

type Product struct {
	ID               string
	Name             string
	Comments         string
	Link             string
	CurrentPrice     float64
	CurrentPriceDate time.Time
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
		Name:             name,
		Comments:         comments,
		Link:             link,
		CurrentPrice:     price,
		CurrentPriceDate: time.Now(),
	}

	return product, nil
}
