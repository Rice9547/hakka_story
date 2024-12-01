package htranslate

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

type (
	Translate struct {
		generator textGenerator
	}

	textGenerator interface {
		Text2Text(ctx context.Context, text string) (string, error)
	}

	TranslateRequest struct {
		Text string `json:"text"`
	}

	TranslateResponse struct {
		Hakka string `json:"hakka"`
	}
)

func New(generator textGenerator) *Translate {
	return &Translate{generator: generator}
}

// TranslateHakka godoc
// @Summary      Translate Text to Hakka
// @Description  Translate a given text to Hakka language
// @Tags         admin translate
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer token"
// @Param        request        body    TranslateRequest  true  "Text to translate"
// @Success      200            {object}  response.Response{data=TranslateResponse}
// @Failure      400            {object}  response.ResponseBase
// @Failure      500            {object}  response.ResponseBase
// @Router       /admin/translate/hakka [post]
func (t *Translate) TranslateHakka(c *gin.Context) {
	var request TranslateRequest
	if err := c.BindJSON(&request); err != nil {
		response.Error(c, 400, "Invalid input")
		return
	}

	hakka, err := t.generator.Text2Text(c.Request.Context(), request.Text)
	if err != nil {
		response.Error(c, 500, "Failed to translate")
		return
	}

	response.Success(c, TranslateResponse{Hakka: hakka})
}
