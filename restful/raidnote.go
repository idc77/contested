package restful

import (
	"dalu/contested/models"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) RaidNoteList(cx *gin.Context) {
	q := bson.M{}
	kcid := cx.Query("keycloak_id")
	if kcid != "" {
		q["keycloak_id"] = kcid
	}
	raidId := cx.Query("raid_id")
	if raidId != "" {
		q["raid_id"] = raidId
	}
	var m []*models.RaidNote
	var e error

	if len(q) == 0 {
		m, e = h.repo.RaidNoteFindAll(nil)
	} else {
		m, e = h.repo.RaidNoteFindAll(q)
	}
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaidNoteGet(cx *gin.Context) {
	q := bson.M{"_id": bson.ObjectIdHex(cx.Param("id"))}
	m, e := h.repo.RaidNoteFind(q)
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaidNotePost(cx *gin.Context) {
	m := new(models.RaidNote)
	if e := cx.BindJSON(m); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m.Id = bson.NewObjectId()
	m.KeycloakId = cx.GetHeader("X-Auth-Subject")

	if e := h.repo.RaidNoteInsert(m); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(201, m)
}

func (h *Handler) RaidNotePut(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	n := new(models.RaidNote)
	if e := cx.BindJSON(n); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m, e := h.repo.RaidNoteFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if n.KeycloakId != m.KeycloakId {
		cx.AbortWithStatus(403)
		return
	}

	if e := h.repo.RaidNoteUpdate(id, n); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, n)
}

func (h *Handler) RaidNoteDelete(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	kcid := cx.GetHeader("X-Auth-Subject")

	m, e := h.repo.RaidNoteFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if kcid != m.KeycloakId {
		cx.AbortWithStatus(403)
		return
	}
	cx.SecureJSON(200, map[string]interface{}{})
}
