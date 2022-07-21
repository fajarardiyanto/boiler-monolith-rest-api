package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"privy/internal/response"
	"privy/models"
)

type PostServiceControllers struct {
	model models.PostRepository
}

func NewPostControllers() PostControllerRepository {
	return &PostServiceControllers{
		model: models.NewModelPost(),
	}
}

func (c *PostServiceControllers) GetPost(w http.ResponseWriter, r *http.Request) {
	res, err := c.model.GetPost()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	logrus.Info(res)
	response.JSON(w, http.StatusOK, res)
}
