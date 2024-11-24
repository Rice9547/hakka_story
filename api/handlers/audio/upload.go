package haudio

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
	supload "github.com/rice9547/hakka_story/service/upload"
)

type (
	Audio struct {
		service supload.UploadService
	}

	UploadAudioRequest struct {
		Audio string `form:"audio" binding:"required"`
	}

	UploadAudioResponse struct {
		URL string `json:"url"`
	}
)

func New(service *supload.UploadService) *Audio {
	return &Audio{
		service: *service,
	}
}

// UploadAudio godoc
// @Summary      Upload Audio
// @Description  Upload Audio
// @Tags         admin audio
// @Accept       multipart/form-data
// @Param		 Authorization  header  string  true  "Bearer token"
// @Param        audio  body      UploadAudioRequest  true  "音檔"
// @Success      201    {object}  response.Response{data=UploadAudioResponse}
// @Failure      400    {object}  response.ResponseBase
// @Failure      401    {object}  response.ResponseBase
// @Failure      500    {object}  response.ResponseBase
// @Router       /admin/audio/upload [post]
func (h *Audio) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("audio")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to read file")
		return
	}
	defer file.Close()

	url, err := h.service.UploadAudio(c, file, header)
	if err != nil {
		if errors.Is(err, errors.ErrUnsupportedFileType) {
			response.Error(c, http.StatusBadRequest, "unsupported file type")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to upload audio")
		return
	}

	response.Success(c, UploadAudioResponse{
		URL: url,
	})
}
