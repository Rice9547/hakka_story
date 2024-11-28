package hcategory

import (
	"github.com/gin-gonic/gin"

	dcategory "github.com/rice9547/hakka_story/domain/category"
	scategory "github.com/rice9547/hakka_story/service/category"
)

type (
	Category struct {
		service scategory.Service
	}

	UpsertRequest struct {
		Name string `json:"name"`
	}

	CategoryResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func New(service scategory.Service) *Category {
	return &Category{service: service}
}

func (r *UpsertRequest) bind(c *gin.Context) (*dcategory.Category, error) {
	if err := c.BindJSON(r); err != nil {
		return nil, err
	}

	return &dcategory.Category{
		Name: r.Name,
	}, nil
}

func toResponse(category dcategory.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
