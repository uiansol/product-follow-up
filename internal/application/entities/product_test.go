package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("should create a product and return it with error nil", func(t *testing.T) {
		product, err := NewProduct("test product", "test product comments", "http://testproduct.com", 10.99)

		assert.Nil(t, err)
		assert.NotNil(t, product.ID)
		assert.Equal(t, "test product", product.Name)
		assert.Equal(t, "test product comments", product.Comments)
		assert.Equal(t, "http://testproduct.com", product.Link)
		assert.Equal(t, 10.99, product.Price)
		assert.NotNil(t, product.PriceDate)
	})

	t.Run("should return error for empty name", func(t *testing.T) {
		_, err := NewProduct("", "test product comments", "http://testproduct.com", 10.99)

		assert.Equal(t, "product name is required", err.Error())
	})

	t.Run("should return error for empty link", func(t *testing.T) {
		_, err := NewProduct("test product", "test product comments", "", 10.99)

		assert.Equal(t, "product link is required", err.Error())
	})

	t.Run("should return error for price 0", func(t *testing.T) {
		_, err := NewProduct("test product", "test product comments", "http://testproduct.com", 0)

		assert.Equal(t, "product price is required", err.Error())
	})
}
