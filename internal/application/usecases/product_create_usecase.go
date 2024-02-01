package usecases

import (
	"github.com/uiansol/product-follow-up/internal/application/entities"
	"github.com/uiansol/product-follow-up/internal/application/interfaces"
)

type IProductCreateUseCase interface {
	Execute(productCreateInput ProductCreateInput) (ProductCreateOutput, error)
}

type ProductCreateUseCase struct {
	productRepository interfaces.IProductRepository
}

type ProductCreateInput struct {
	Name     string
	Comments string
	Link     string
	Price    float64
}

type ProductCreateOutput struct {
	ID string
}

func NewProductCreateUseCase(productRepository interfaces.IProductRepository) *ProductCreateUseCase {
	return &ProductCreateUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductCreateUseCase) Execute(productCreateInput ProductCreateInput) (ProductCreateOutput, error) {
	product, err := entities.NewProduct(productCreateInput.Name, productCreateInput.Comments, productCreateInput.Link, productCreateInput.Price)
	if err != nil {
		return ProductCreateOutput{}, err
	}

	id, err := u.productRepository.Save(&product)
	if err != nil {
		return ProductCreateOutput{}, err
	}

	output := ProductCreateOutput{
		ID: id,
	}

	return output, nil
}
