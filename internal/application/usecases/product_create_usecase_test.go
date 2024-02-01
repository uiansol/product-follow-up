package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
)

func TestNewProductCreateUseCase(t *testing.T) {
	productRepositoryMock := mocksdb.NewIProductRepository(t)

	t.Run("should return a create user use case", func(t *testing.T) {
		createUserUseCase := NewProductCreateUseCase(productRepositoryMock)

		assert.NotNil(t, createUserUseCase)
	})
}

func TestProductCreateExecute(t *testing.T) {
	input := ProductCreateInput{
		Name:        "test product",
		Description: "test description",
		Link:        "https://test.com",
		Price:       10.5,
	}

	t.Run("should create the product and return error nil", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productCreateUseCase := NewProductCreateUseCase(productRepositoryMock)

		productRepositoryMock.On("Save", mock.Anything).Return("123", nil)

		output, err := productCreateUseCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, "123", output.ProductID)
	})

	t.Run("should return error when name is invalid", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productCreateUseCase := NewProductCreateUseCase(productRepositoryMock)

		inputInvalidName := ProductCreateInput{
			Name:        "",
			Description: "test description",
			Link:        "https://test.com",
			Price:       10.5,
		}

		output, err := productCreateUseCase.Execute(inputInvalidName)
		assert.Equal(t, "product name is required", err.Error())
		assert.Equal(t, ProductCreateOutput{}, output)
	})

	t.Run("should return error when couldn't save", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productCreateUseCase := NewProductCreateUseCase(productRepositoryMock)

		productRepositoryMock.On("Save", mock.Anything).Return("", assert.AnError)

		output, err := productCreateUseCase.Execute(input)

		assert.Equal(t, assert.AnError, err)
		assert.Equal(t, ProductCreateOutput{}, output)
	})
}
