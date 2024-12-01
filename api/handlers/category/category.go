package hcategory

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/entities"
)

type (
	Service interface {
		Create(ctx context.Context, c *entities.Category) (*entities.Category, error)
		ListByName(ctx context.Context, name string) ([]entities.Category, error)
		Update(ctx context.Context, id uint64, name string) (*entities.Category, error)
		DeleteByID(ctx context.Context, id uint64) error
	}

	Category struct {
		service Service
	}

	UpsertRequest struct {
		Name string `json:"name"`
	}

	CategoryResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func New(service Service) *Category {
	return &Category{service: service}
}

func (r *UpsertRequest) bind(c *gin.Context) (*entities.Category, error) {
	if err := c.BindJSON(r); err != nil {
		return nil, err
	}

	return &entities.Category{
		Name: r.Name,
	}, nil
}

func toResponse(category entities.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
