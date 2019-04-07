package restful

import (
	"dalu/contested/models"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) RaiderList(cx *gin.Context) {
	kcid := cx.Query("keycloak_id")
	m, e := h.repo.RaiderFindAll(bson.M{"keycloak_id": kcid})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaiderGet(cx *gin.Context) {
	q := bson.M{"_id": bson.ObjectIdHex(cx.Param("id"))}
	m, e := h.repo.RaiderFind(q)
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaiderPost(cx *gin.Context) {
	m := new(models.Raider)
	if e := cx.BindJSON(m); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m.Id = bson.NewObjectId()
	m.KeycloakId = cx.GetHeader("X-Auth-Subject")

	if e := h.repo.RaiderInsert(m); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(201, m)
}

func (h *Handler) RaiderPut(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	n := new(models.Raider)
	if e := cx.BindJSON(n); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m, e := h.repo.RaiderFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if n.KeycloakId != m.KeycloakId {
		cx.AbortWithStatus(403)
		return
	}

	if e := h.repo.RaiderUpdate(id, n); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, n)
}

func (h *Handler) RaiderDelete(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	kcid := cx.GetHeader("X-Auth-Subject")

	m, e := h.repo.RaiderFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if kcid != m.KeycloakId {
		cx.AbortWithStatus(403)
		return
	}
	if e := h.repo.RaiderRemove(id); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, map[string]interface{}{})
}
