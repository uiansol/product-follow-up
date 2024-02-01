package mysql

import (
	"time"

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
