package haudio

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

// GenerateAudio godoc
// @Summary      Generate Audio
// @Description  Generate an audio from a text
// @Tags         admin audio
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token"
// @Param        request        body    GenerateRequest  true  "Text prompt"
// @Success      200            {object}  response.Response{data=GenerateResponse}
// @Failure      400            {object}  response.ResponseBase
// @Failure      401            {object}  response.ResponseBase
// @Failure      500            {object}  response.ResponseBase
// @Router       /admin/audio/generate [post]
func (h *Audio) Generate(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err, "Invalid input")
		return
	}

	data, err := h.generator.Text2Speech(c.Request.Context(), req.Prompt)
	if err != nil {
		response.InternalServerError(c, err, "Failed to generate audio")
		return
	}

	header := &multipart.FileHeader{
		Filename: "audio.mp3",
		Header:   make(textproto.MIMEHeader),
	}
	header.Header.Add("Content-Type", "audio/mpeg")
	url, err := h.uploader.UploadAudio(c.Request.Context(), bytes.NewReader(data), header)
	if err != nil {
		response.InternalServerError(c, err, "Failed to upload audio")
		return
	}

	response.Success(c, GenerateResponse{
		URL: url,
	})
}
