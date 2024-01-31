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
	Name        string
	Description string
	Link        string
	Price       float64
}

type ProductCreateOutput struct {
	ProductID string
}

func NewProductCreateUseCase(productRepository interfaces.IProductRepository) ProductCreateUseCase {
	return ProductCreateUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductCreateUseCase) Execute(productCreateInput ProductCreateInput) (ProductCreateOutput, error) {
	product, err := entities.NewProduct(productCreateInput.Name, productCreateInput.Description, productCreateInput.Link, productCreateInput.Price)
	if err != nil {
		return ProductCreateOutput{}, err
	}

	id, err := u.productRepository.Save(product)
	if err != nil {
		return ProductCreateOutput{}, err
	}

	output := ProductCreateOutput{
		ProductID: id,
	}

	return output, nil
}
