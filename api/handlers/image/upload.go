package himage

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

type (
	UploadImageRequest struct {
		Image string `form:"image" binding:"required"`
	}

	UploadImageResponse struct {
		URL string `json:"url"`
	}
)

// UploadImage godoc
// @Summary      Upload Image
// @Description  Upload Image
// @Tags         admin image
// @Accept       multipart/form-data
// @Param		 Authorization  header  string  true  "Bearer token"
// @Param        image  body      UploadImageRequest  true  "圖片"
// @Success      201    {object}  response.Response{data=UploadImageResponse}
// @Failure      400    {object}  response.ResponseBase
// @Failure      401    {object}  response.ResponseBase
// @Failure      500    {object}  response.ResponseBase
// @Router       /admin/image/upload [post]
func (h *Image) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to read file")
		return
	}
	defer file.Close()

	url, err := h.uploader.UploadImage(c.Request.Context(), file, header)
	if err != nil {
		if errors.Is(err, errors.ErrUnsupportedFileType) {
			response.Error(c, http.StatusBadRequest, "unsupported file type")
			return
		}
		response.Error(c, http.StatusInternalServerError, "failed to upload image")
		return
	}

	response.Success(c, UploadImageResponse{
		URL: url,
	})
}
