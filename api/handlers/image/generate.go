package himage

import (
	"bytes"
	"mime/multipart"
	"net/textproto"

	"github.com/gin-gonic/gin"

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
		response.BadRequest(c, err, "Invalid input")
		return
	}

	_, data, err := h.generator.Text2Image(c.Request.Context(), req.Prompt)
	if err != nil {
		response.InternalServerError(c, err, "Failed to generate image")
		return
	}

	header := &multipart.FileHeader{
		Filename: "image.png",
		Header:   make(textproto.MIMEHeader),
	}
	header.Header.Add("Content-Type", "image/png")
	url, err := h.uploader.UploadImage(c.Request.Context(), bytes.NewReader(data), header)
	if err != nil {
		response.InternalServerError(c, err, "Failed to upload image")
		return
	}

	response.Success(c, GenerateResponse{
		URL: url,
	})
}
