package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/satori/go.uuid"
	"net/http"
	"retro-memo/server/httputil"
)

// @Summary create a new topic
// @Accept json
// @Produce json
// @Param topic body topic.addTopicCommand true "Add topic"
// @Success 202 {string} string
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /topics [post]
func (h *handler) AddTopic(context *gin.Context) {
	var command *addTopicCommand
	if err := context.ShouldBindWith(&command, binding.JSON); err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	if err := h.useCase.add(command); err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.Status(http.StatusAccepted)
}

// @Summary list sprint topics
// @Accept json
// @Produce json
// @Param sprint_id path string true "Sprint ID"
// @Success 200 {array} topic.topicModel
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /topics/{sprint_id} [get]
func (h *handler) ListTopic(context *gin.Context) {

	sprintId := context.Params.ByName("sprintId")
	sprintUUID, err := uuid.FromString(sprintId)
	if err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	result, err := h.useCase.list(&listTopicQuery{sprintUUID})
	if err != nil {
		httputil.NewError(context, http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, result)
}

type handler struct {
	useCase *useCase
}

func NewHandler() *handler {
	repository := &inMemoryRepository{map[uuid.UUID]sprintTopics{}}
	useCase := &useCase{repository}
	return &handler{useCase}
}
