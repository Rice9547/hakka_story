package hcategory

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags admin categories
// @Accept json
// @Produce json
// @Param category body UpsertRequest true "Category to create"
// @Success 200 {object} CategoryResponse
// @Failure 400 {object} response.ResponseBase
// @Failure 500 {object} response.ResponseBase
// @Router /categories [post]
func (h *Category) Create(c *gin.Context) {
	request := new(UpsertRequest)
	category, err := request.bind(c)
	if err != nil {
		response.Error(c, 400, err.Error())
	}

	category, err = h.service.Create(c.Request.Context(), category)
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to create category, err: %v", err))
		return
	}

	response.Success(c, toResponse(*category))
}
