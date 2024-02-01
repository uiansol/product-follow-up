package dto

import "time"

type ProductCreateRequest struct {
	Name     string  `json:"name"`
	Comments string  `json:"description"`
	Link     string  `json:"link"`
	Price    float64 `json:"price"`
}

type ProductCreateResponse struct {
	ID string `json:"id"`
}

type ProductIDRequest struct {
	ID string `param:"id"`
}

type ProductReadResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Comments  string    `json:"comments"`
	Link      string    `json:"link"`
	Price     float64   `json:"price"`
	PriceDate time.Time `json:"price_date"`
}

type ProductUpdateRequest struct {
	ID       string  `param:"id"`
	Name     string  `json:"name"`
	Comments string  `json:"comments"`
	Link     string  `json:"link"`
	Price    float64 `json:"price"`
}
