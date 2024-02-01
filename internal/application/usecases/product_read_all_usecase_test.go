package usecases

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
	"github.com/uiansol/product-follow-up/internal/application/entities"
)

func TestNewProductReadAllUseCase(t *testing.T) {
	productRepositoryMock := mocksdb.NewIProductRepository(t)

	t.Run("should return a read all user use case", func(t *testing.T) {
		ReadAllUserUseCase := NewProductReadAllUseCase(productRepositoryMock)

		assert.NotNil(t, ReadAllUserUseCase)
	})
}

func TestProductReadAllExecute(t *testing.T) {
	t.Run("should read all the product and return error nil", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productReadAllUseCase := NewProductReadAllUseCase(productRepositoryMock)

		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

		var products []*entities.Product
		products = append(products, &entities.Product{
			ID:        "123",
			Name:      "test name",
			Comments:  "test comments",
			Link:      "https://test.com",
			Price:     10.5,
			PriceDate: date,
		})
		products = append(products, &entities.Product{
			ID:        "456",
			Name:      "test name 2",
			Comments:  "test comments 2",
			Link:      "https://test.com",
			Price:     13.5,
			PriceDate: date,
		})

		productRepositoryMock.On("ReadAll", mock.Anything).Return(products, nil)

		output, err := productReadAllUseCase.Execute()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(output))

		assert.Equal(t, "123", output[0].ID)
		assert.Equal(t, "test name", output[0].Name)
		assert.Equal(t, "test comments", output[0].Comments)
		assert.Equal(t, "https://test.com", output[0].Link)
		assert.Equal(t, 10.5, output[0].Price)
		assert.Equal(t, date, output[0].PriceDate)

		assert.Equal(t, "456", output[1].ID)
		assert.Equal(t, "test name 2", output[1].Name)
		assert.Equal(t, "test comments 2", output[1].Comments)
		assert.Equal(t, "https://test.com", output[1].Link)
		assert.Equal(t, 13.5, output[1].Price)
		assert.Equal(t, date, output[1].PriceDate)
	})

	t.Run("should return error when read all has error", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productReadAllUseCase := NewProductReadAllUseCase(productRepositoryMock)

		productRepositoryMock.On("ReadAll", mock.Anything).Return(nil, assert.AnError)

		_, err := productReadAllUseCase.Execute()

		assert.NotNil(t, err)
	})
}
