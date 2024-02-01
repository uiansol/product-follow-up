package mysql

import (
	"errors"
	"time"

	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type ProductDB struct {
	gorm.Model
	UUID      string `gorm:"primaryKey"`
	Name      string
	Comments  string
	Link      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Save(product *entities.Product) (string, error) {
	productDB := ProductDB{
		UUID:      product.ID,
		Name:      product.Name,
		Comments:  product.Comments,
		Link:      product.Link,
		Price:     product.Price,
		CreatedAt: product.PriceDate,
		UpdatedAt: product.PriceDate,
	}

	result := p.db.Create(&productDB)
	if result.Error != nil {
		return "", result.Error
	}

	return productDB.UUID, nil
}

func (p *ProductRepository) Read(id string) (*entities.Product, error) {
	productDB := ProductDB{}

	result := p.db.First(&productDB, "uuid = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &entities.Product{}, errors.New(apperr.ErrNotFound)
		}
		return &entities.Product{}, result.Error
	}

	product := entities.Product{
		ID:        productDB.UUID,
		Name:      productDB.Name,
		Comments:  productDB.Comments,
		Link:      productDB.Link,
		Price:     productDB.Price,
		PriceDate: productDB.UpdatedAt,
	}

	return &product, nil
}

func (p *ProductRepository) Update(product *entities.Product) error {
	productDB := ProductDB{}

	result := p.db.First(&productDB, "uuid = ?", product.ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New(apperr.ErrNotFound)
		}
		return result.Error
	}
	productDB.Name = product.Name
	productDB.Comments = product.Comments
	productDB.Link = product.Link
	productDB.Price = product.Price
	productDB.UpdatedAt = product.PriceDate

	result = p.db.Save(&productDB)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepository) Delete(id string) error {
	result := p.db.Unscoped().Delete(&ProductDB{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New(apperr.ErrNotFound)
	}

	return nil
}
