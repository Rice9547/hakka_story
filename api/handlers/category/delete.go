package hcategory

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Delete a category by its ID
// @Tags admin categories
// @Accept json
// @Produce json
// @Param id path uint64 true "Category ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ResponseBase
// @Failure 404 {object} response.ResponseBase
// @Failure 500 {object} response.ResponseBase
// @Router /admin/category/{id} [delete]
func (h *Category) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Invalid category ID")
		return
	}

	if err = h.service.DeleteByID(c.Request.Context(), id); err != nil {
		if errors.Is(err, errors.ErrStoryNotFound) {
			response.NotFound(c, "Category not found")
			return
		}
		response.InternalServerError(c, err, "Failed to delete category")
		return
	}

	response.Success(c, nil)
}
