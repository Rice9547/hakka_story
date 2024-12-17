package hcategory

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

// UpdateCategory godoc
// @Summary      Update category
// @Description  Update category by id
// @Tags         admin category
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Param        category  body  UpsertRequest  true  "Category data"
// @Success      200  {object}  response.Response{data=CategoryResponse}
// @Failure      400  {object}  response.ResponseBase
// @Failure      404  {object}  response.ResponseBase
// @Failure      500  {object}  response.ResponseBase
// @Router       /category/{id} [put]
func (h *Category) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Invalid category ID")
		return
	}

	req := new(UpsertRequest)
	category, err := req.bind(c)
	if err != nil {
		response.BadRequest(c, err, "Invalid input")
		return
	}

	category, err = h.service.Update(c.Request.Context(), id, category.Name)
	if err != nil {
		if errors.Is(err, errors.ErrCategoryNotFound) {
			response.NotFound(c, "Category not found")
			return
		}
		response.InternalServerError(c, err, "Failed to update category")
		return
	}

	response.Success(c, toResponse(*category))
}
