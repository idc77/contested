package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) ProfileList(cx *gin.Context) {
	m, e := h.repo.ProfileFindAll(nil)
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) ProfileGet(cx *gin.Context) {
	q := bson.M{"_id": bson.ObjectIdHex(cx.Param("id"))}
	m, e := h.repo.ProfileFind(q)
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) ProfilePost(cx *gin.Context) {

}

func (h *Handler) ProfilePut(cx *gin.Context) {

}

func (h *Handler) ProfileDelete(cx *gin.Context) {

}
