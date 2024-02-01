package mappers

import (
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func ProductCreateRequestToProductCreateInput(p dto.ProductCreateRequest) usecases.ProductCreateInput {
	return usecases.ProductCreateInput{
		Name:     p.Name,
		Comments: p.Description,
		Link:     p.Link,
		Price:    p.Price,
	}
}

func ProductCreateOutputToProductCreateResponse(p usecases.ProductCreateOutput) dto.ProductCreateResponse {
	return dto.ProductCreateResponse{
		ProductID: p.ProductID,
	}
}

func ProductReadRequestToProductReadInput(p dto.ProductReadRequest) usecases.ProductReadInput {
	return usecases.ProductReadInput{
		ID: p.ID,
	}
}

func ProductReadOutputToProductReadResponse(p usecases.ProductReadOutput) dto.ProductReadResponse {
	return dto.ProductReadResponse{
		ID:        p.ID,
		Name:      p.Name,
		Comments:  p.Comments,
		Link:      p.Link,
		Price:     p.Price,
		PriceDate: p.PriceDate,
	}
}
