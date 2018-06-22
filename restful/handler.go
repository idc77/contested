package restful

import (
	"dalu/contested/restful/repo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repo.Repo
}

func NewHandler(r *repo.Repo) *Handler {
	return &Handler{
		repo: r,
	}
}

func (h *Handler) TestHandler(cx *gin.Context) {
	cx.JSON(200, cx.Request.Header)
}
