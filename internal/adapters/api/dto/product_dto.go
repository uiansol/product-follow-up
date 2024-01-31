package dto

type (
	ProductCreateRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Link        string  `json:"link"`
		Price       float64 `json:"price"`
	}
)

type ProductCreateResponse struct {
	ProductID string `json:"product_id"`
}
