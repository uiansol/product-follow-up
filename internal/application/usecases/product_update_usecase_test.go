package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
)

func TestNewProductUpdateUseCase(t *testing.T) {
	productRepositoryMock := mocksdb.NewIProductRepository(t)

	t.Run("should return a Update user use case", func(t *testing.T) {
		UpdateUserUseCase := NewProductUpdateUseCase(productRepositoryMock)

		assert.NotNil(t, UpdateUserUseCase)
	})
}

func TestProductUpdateExecute(t *testing.T) {
	t.Run("should update the product and return error nil", func(t *testing.T) {
		input := ProductUpdateInput{
			ID:       "123",
			Name:     "test product update",
			Comments: "test description",
			Link:     "https://test.com",
			Price:    10.5,
		}

		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productUpdateUseCase := NewProductUpdateUseCase(productRepositoryMock)

		productRepositoryMock.On("Update", mock.Anything).Return(nil)

		err := productUpdateUseCase.Execute(input)

		assert.Nil(t, err)
	})

	t.Run("should return error when couldn't update", func(t *testing.T) {
		input := ProductUpdateInput{
			ID:       "123",
			Name:     "test product update",
			Comments: "test description",
			Link:     "https://test.com",
			Price:    10.5,
		}

		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productUpdateUseCase := NewProductUpdateUseCase(productRepositoryMock)

		productRepositoryMock.On("Update", mock.Anything).Return(assert.AnError)

		err := productUpdateUseCase.Execute(input)

		assert.Equal(t, "error updating product", err.Error())
	})
}
