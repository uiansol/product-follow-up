package mysql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/uiansol/product-follow-up/internal/application/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSave(t *testing.T) {
	t.Run("should save data and return error nil", func(t *testing.T) {
		mockDb, mock, _ := sqlmock.New()
		dialector := mysql.New(mysql.Config{
			Conn:       mockDb,
			DriverName: "mysql",
		})
		db, _ := gorm.Open(dialector, &gorm.Config{})

		productRepository := NewProductRepository(db)

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `product_d_bs`").WillReturnResult(sqlmock.NewResult(1, 1))

		product, _ := entities.NewProduct("test product", "test description", "https://test.com", 10.5)

		id, err := productRepository.Save(product)
		assert.Nil(t, err)
		assert.Equal(t, product.ID, id)
	})
}
