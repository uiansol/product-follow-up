package usecases

import (
	"time"

	"github.com/uiansol/product-follow-up/internal/application/interfaces"
)

type IProductReadUseCase interface {
	Execute(productReadInput ProductReadInput) (ProductReadOutput, error)
}

type ProductReadUseCase struct {
	productRepository interfaces.IProductRepository
}

type ProductReadInput struct {
	ID string
}

type ProductReadOutput struct {
	ID        string
	Name      string
	Comments  string
	Link      string
	Price     float64
	PriceDate time.Time
}

func NewProductReadUseCase(productRepository interfaces.IProductRepository) *ProductReadUseCase {
	return &ProductReadUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductReadUseCase) Execute(productReadInput ProductReadInput) (ProductReadOutput, error) {
	product, err := u.productRepository.Read(productReadInput.ID)
	if err != nil {
		return ProductReadOutput{}, err
	}

	output := ProductReadOutput{
		ID:        product.ID,
		Name:      product.Name,
		Comments:  product.Comments,
		Link:      product.Link,
		Price:     product.Price,
		PriceDate: product.PriceDate,
	}

	return output, nil
}
