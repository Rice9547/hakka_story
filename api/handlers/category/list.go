package hcategory

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

// ListCategory godoc
// @Summary      List categories
// @Description  Get list of categories by name
// @Tags         categories
// @Produce      json
// @Param        name  query     string  false  "Category name"
// @Success      200   {array}   response.Response{data=CategoryResponse}
// @Failure      500   {object}  response.ResponseBase
// @Router       /category [get]
func (h *Category) List(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	categories, err := h.service.ListByName(c.Request.Context(), name)
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to retrieve categories, err: %v", err))
		return
	}

	resp := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, toResponse(category))
	}

	response.Success(c, resp)
}
