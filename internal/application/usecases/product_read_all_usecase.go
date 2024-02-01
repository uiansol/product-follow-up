package usecases

import (
	"time"

	"github.com/uiansol/product-follow-up/internal/application/interfaces"
)

type IProductReadAllUseCase interface {
	Execute() ([]*ProductReadAllOutput, error)
}

type ProductReadAllUseCase struct {
	productRepository interfaces.IProductRepository
}

type ProductReadAllOutput struct {
	ID        string
	Name      string
	Comments  string
	Link      string
	Price     float64
	PriceDate time.Time
}

func NewProductReadAllUseCase(productRepository interfaces.IProductRepository) *ProductReadAllUseCase {
	return &ProductReadAllUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductReadAllUseCase) Execute() ([]*ProductReadAllOutput, error) {
	products, err := u.productRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	var output []*ProductReadAllOutput
	for _, product := range products {
		output = append(output, &ProductReadAllOutput{
			ID:        product.ID,
			Name:      product.Name,
			Comments:  product.Comments,
			Link:      product.Link,
			Price:     product.Price,
			PriceDate: product.PriceDate,
		})
	}

	return output, nil
}
