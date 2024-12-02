package hexercise

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

type (
	CountStoryExerciseResponse struct {
		StoryID    uint64 `json:"story_id"`
		StoryTitle string `json:"story_title"`
		Count      int64  `json:"count"`
	}
)

// CountStoriesExercise godoc
// @Summary Count exercises for stories
// @Description Retrieves the count of exercises associated with specific story IDs
// @Tags admin exercise
// @Accept json
// @Produce json
// @Param story_ids query []uint64 true "Story IDs"
// @Success 200 {object} response.Response{data=[]CountStoryExerciseResponse}
// @Failure 400 {object} response.ResponseBase "Invalid story id"
// @Failure 500 {object} response.ResponseBase "Internal server error"
// @Router /admin/story/exercise [get]
func (h *Exercise) CountStoriesExercise(c *gin.Context) {
	qryStoryIds := c.QueryArray("story_ids")
	storyIDs := make([]uint64, 0, len(qryStoryIds))
	for _, sid := range qryStoryIds {
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "Invalid story id")
			return
		}
		storyIDs = append(storyIDs, id)
	}

	counts, err := h.service.GetExerciseCountByStoryIDs(c, storyIDs)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := make([]CountStoryExerciseResponse, 0, len(counts))
	for _, counter := range counts {
		resp = append(resp, CountStoryExerciseResponse{
			StoryID:    counter.StoryID,
			StoryTitle: counter.StoryTitle,
			Count:      counter.Count,
		})
	}

	response.Success(c, resp)
}

// ListExerciseByStoryID godoc
// @Summary List exercises by story ID
// @Description Retrieves a list of exercises associated with a specific story ID
// @Tags admin exercise
// @Accept json
// @Produce json
// @Param id path uint64 true "Story ID"
// @Success 200 {object} response.Response{data=[]ExerciseResponse}
// @Failure 400 {object} response.ResponseBase "Invalid story id"
// @Failure 500 {object} response.ResponseBase "Internal server error"
// @Router /admin/story/{id}/exercise [get]
func (h *Exercise) ListExerciseByStoryID(c *gin.Context) {
	storyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid story id")
		return
	}

	exercises, err := h.service.ListExerciseByStoryID(c, storyID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := make([]ExerciseResponse, 0, len(exercises))
	for _, exercise := range exercises {
		resp = append(resp, toExerciseResponse(exercise, true))
	}

	response.Success(c, resp)
}
