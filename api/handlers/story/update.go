package hstory

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

// UpdateStory godoc
// @Summary      Update Story
// @Description  Update Story by ID
// @Tags         admin story
// @Accept       json
// @Produce      json
// @Param		 Authorization  header  string  true  "Bearer token"
// @Param        story  body      UpsertStoryRequest  true  "故事內容"
// @Success      200    {object}  response.Response{data=FullStoryResponse}
// @Failure      400    {object}  response.ResponseBase
// @Failure      401    {object}  response.ResponseBase
// @Failure      500    {object}  response.ResponseBase
// @Router       /admin/story/{id} [put]
func (h *Story) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Invalid story ID")
		return
	}

	request := new(UpsertStoryRequest)
	story, err := request.bind(c)
	if err != nil {
		response.BadRequest(c, err, "Invalid input")
		return
	}

	if err = h.service.UpdateByID(c.Request.Context(), id, story); err != nil {
		if errors.Is(err, errors.ErrStoryNotFound) {
			response.NotFound(c, "Story not found")
			return
		}
		response.InternalServerError(c, err, "Failed to update story")
		return
	}

	response.Success(c, toFullyResponse(*story))
}
