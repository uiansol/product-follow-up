package mappers

import (
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func ProductCreateRequestToProductCreateInput(p dto.ProductCreateRequest) usecases.ProductCreateInput {
	return usecases.ProductCreateInput{
		Name:        p.Name,
		Description: p.Description,
		Link:        p.Link,
		Price:       p.Price,
	}
}

func ProductCreateOutputToProductCreateResponse(p usecases.ProductCreateOutput) dto.ProductCreateResponse {
	return dto.ProductCreateResponse{
		ProductID: p.ProductID,
	}
}
