package usecases

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mocksdb"
	"github.com/uiansol/product-follow-up/internal/application/entities"
)

func TestNewProductReadUseCase(t *testing.T) {
	productRepositoryMock := mocksdb.NewIProductRepository(t)

	t.Run("should return a Read user use case", func(t *testing.T) {
		ReadUserUseCase := NewProductReadUseCase(productRepositoryMock)

		assert.NotNil(t, ReadUserUseCase)
	})
}

func TestProductReadExecute(t *testing.T) {
	input := ProductReadInput{
		ID: "123",
	}

	t.Run("should read the product and return error nil", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productReadUseCase := NewProductReadUseCase(productRepositoryMock)

		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		productRepositoryMock.On("Read", mock.Anything).Return(
			&entities.Product{
				ID:        "123",
				Name:      "test name",
				Comments:  "test comments",
				Link:      "https://test.com",
				Price:     10.5,
				PriceDate: date,
			}, nil)

		output, err := productReadUseCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, "123", output.ID)
		assert.Equal(t, "test name", output.Name)
		assert.Equal(t, "test comments", output.Comments)
		assert.Equal(t, "https://test.com", output.Link)
		assert.Equal(t, 10.5, output.Price)
		assert.Equal(t, date, output.PriceDate)
	})

	t.Run("should return error read has error", func(t *testing.T) {
		productRepositoryMock := mocksdb.NewIProductRepository(t)
		productReadUseCase := NewProductReadUseCase(productRepositoryMock)

		productRepositoryMock.On("Read", mock.Anything).Return(&entities.Product{}, assert.AnError)

		output, err := productReadUseCase.Execute(input)
		assert.NotNil(t, err)
		assert.Equal(t, ProductReadOutput{}, output)
	})
}
