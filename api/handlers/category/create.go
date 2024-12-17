package hcategory

import (
	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags admin category
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
		response.BadRequest(c, err, "Invalid input")
	}

	category, err = h.service.Create(c.Request.Context(), category)
	if err != nil {
		response.InternalServerError(c, err, "Failed to create category")
		return
	}

	response.Success(c, toResponse(*category))
}
