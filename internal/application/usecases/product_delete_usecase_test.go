package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
)

func TestNewProductDeleteUseCase(t *testing.T) {
	productRepositoryMock := mocksdb.NewIProductRepository(t)

	t.Run("should return a Delete user use case", func(t *testing.T) {
		DeleteUserUseCase := NewProductDeleteUseCase(productRepositoryMock)

		assert.NotNil(t, DeleteUserUseCase)
	})
}

func TestProductDeleteExecute(t *testing.T) {
	input := ProductDeleteInput{
		ID: "123",
	}

	t.Run("should delete the product and return error nil", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productDeleteUseCase := NewProductDeleteUseCase(productRepositoryMock)

		productRepositoryMock.On("Delete", mock.Anything).Return(nil)

		err := productDeleteUseCase.Execute(input)

		assert.Nil(t, err)
	})

	t.Run("should return error delete has error", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productDeleteUseCase := NewProductDeleteUseCase(productRepositoryMock)

		productRepositoryMock.On("Delete", mock.Anything).Return(assert.AnError)

		err := productDeleteUseCase.Execute(input)

		assert.NotNil(t, err)
	})
}
