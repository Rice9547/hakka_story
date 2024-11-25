package himage

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

type (
	GenerateRequest struct {
		Prompt string `json:"prompt"`
	}

	GenerateResponse struct {
		URL string `json:"url"`
	}
)

// GenerateImage godoc
// @Summary      Generate Image
// @Description  Generate an image from a text prompt
// @Tags         admin image
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token"
// @Param        request        body    GenerateRequest  true  "Text prompt"
// @Success      200            {object}  response.Response{data=GenerateResponse}
// @Failure      400            {object}  response.ResponseBase
// @Failure      401            {object}  response.ResponseBase
// @Failure      500            {object}  response.ResponseBase
// @Router       /admin/image/generate [post]
func (h *Image) Generate(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, err.Error()))
		return
	}

	url, err := h.generator.Text2Image(req.Prompt)
	if err != nil {
		errors.ErrorHandler(c, err)
		return
	}

	response.Success(c, GenerateResponse{
		URL: url,
	})
}
