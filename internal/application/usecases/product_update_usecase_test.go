package usecases

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
	"github.com/uiansol/product-follow-up/internal/application/entities"
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
		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

		input := ProductUpdateInput{
			ID:        "123",
			Name:      "test product update",
			Comments:  "test description",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		product := entities.Product{
			ID:        "123",
			Name:      "test product",
			Comments:  "test comments",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productUpdateUseCase := NewProductUpdateUseCase(productRepositoryMock)

		productRepositoryMock.On("Read", mock.Anything).Return(&product, nil)
		productRepositoryMock.On("Update", mock.Anything).Return(nil)

		err := productUpdateUseCase.Execute(input)

		assert.Nil(t, err)
	})

	t.Run("should return error when couldn't read", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		input := ProductUpdateInput{
			ID:        "123",
			Name:      "test product update",
			Comments:  "test description",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productUpdateUseCase := NewProductUpdateUseCase(productRepositoryMock)

		productRepositoryMock.On("Read", mock.Anything).Return(nil, assert.AnError)

		err := productUpdateUseCase.Execute(input)

		assert.Equal(t, "error reading product", err.Error())
	})

	t.Run("should return error when couldn't update", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		input := ProductUpdateInput{
			ID:        "123",
			Name:      "test product update",
			Comments:  "test description",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		product := entities.Product{
			ID:        "123",
			Name:      "test product",
			Comments:  "test comments",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productUpdateUseCase := NewProductUpdateUseCase(productRepositoryMock)

		productRepositoryMock.On("Read", mock.Anything).Return(&product, nil)
		productRepositoryMock.On("Update", mock.Anything).Return(assert.AnError)

		err := productUpdateUseCase.Execute(input)

		assert.Equal(t, "error updating product", err.Error())
	})
}
