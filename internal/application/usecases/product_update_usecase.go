package usecases

import (
	"errors"
	"time"

	"github.com/uiansol/product-follow-up/internal/application/interfaces"
)

type IProductUpdateUseCase interface {
	Execute(productUpdateInput ProductUpdateInput) error
}

type ProductUpdateUseCase struct {
	productRepository interfaces.IProductRepository
}

type ProductUpdateInput struct {
	ID        string
	Name      string
	Comments  string
	Link      string
	Price     float64
	PriceDate time.Time
}

func NewProductUpdateUseCase(productRepository interfaces.IProductRepository) *ProductUpdateUseCase {
	return &ProductUpdateUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductUpdateUseCase) Execute(productUpdateInput ProductUpdateInput) error {
	product, err := u.productRepository.Read(productUpdateInput.ID)
	if err != nil {
		return errors.New("error reading product")
	}

	product.Name = productUpdateInput.Name
	product.Comments = productUpdateInput.Comments
	product.Link = productUpdateInput.Link
	product.Price = productUpdateInput.Price
	product.PriceDate = productUpdateInput.PriceDate

	err = u.productRepository.Update(product)
	if err != nil {
		return errors.New("error updating product")
	}

	return nil
}
