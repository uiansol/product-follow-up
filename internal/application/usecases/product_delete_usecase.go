package usecases

import (
	"github.com/uiansol/product-follow-up/internal/application/interfaces"
)

type IProductDeleteUseCase interface {
	Execute(productDeleteInput ProductDeleteInput) error
}

type ProductDeleteUseCase struct {
	productRepository interfaces.IProductRepository
}

type ProductDeleteInput struct {
	ID string
}

func NewProductDeleteUseCase(productRepository interfaces.IProductRepository) *ProductDeleteUseCase {
	return &ProductDeleteUseCase{
		productRepository: productRepository,
	}
}

func (u *ProductDeleteUseCase) Execute(productDeleteInput ProductDeleteInput) error {
	err := u.productRepository.Delete(productDeleteInput.ID)
	if err != nil {
		return err
	}

	return nil
}
