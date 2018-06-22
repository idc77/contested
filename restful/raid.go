package restful

import (
	"dalu/contested/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) RaidList(cx *gin.Context) {
	keycloakId := cx.Query("keycloak_id")

	q := bson.M{}

	if keycloakId != "" {
		q["keycloak_id"] = keycloakId
	}

	q["date_end"] = bson.M{"$gt": bson.Now().Add(time.Hour * 2)}

	var m []*models.Raid
	var e error

	if len(q) == 0 {
		m, e = h.repo.RaidFindAll(nil)
	} else {
		m, e = h.repo.RaidFindAll(q)
	}

	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaidGet(cx *gin.Context) {
	q := bson.M{"_id": bson.ObjectIdHex(cx.Param("id"))}
	m, e := h.repo.RaidFind(q)
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, m)
}

func (h *Handler) RaidPost(cx *gin.Context) {
	m := new(models.Raid)
	if e := cx.BindJSON(m); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m.Id = bson.NewObjectId()
	m.KeycloakId = cx.GetHeader("X-Auth-Subject")
	m.DateMeet = m.DateStart.Add(time.Minute * -15)

	if m.DateEnd.Sub(m.DateStart) < 0 {
		m.DateEnd = m.DateEnd.Add(24 * time.Hour)
	}

	if e := h.repo.RaidInsert(m); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(201, m)
}

func (h *Handler) RaidPut(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	n := new(models.Raid)
	if e := cx.BindJSON(n); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	m, e := h.repo.RaidFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if n.KeycloakId != m.KeycloakId {
		// MT
		if n.Setup.MT.Rules.Tank != m.Setup.MT.Rules.Tank {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.MT.Rules.SwashOrHatetransfer != m.Setup.MT.Rules.SwashOrHatetransfer {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.MT.Rules.DirgeOrBard != m.Setup.MT.Rules.DirgeOrBard {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.MT.Rules.CoercerOrEnchanter != m.Setup.MT.Rules.CoercerOrEnchanter {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.MT.Rules.DefilerOrWarder != m.Setup.MT.Rules.DefilerOrWarder {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.MT.Rules.TemplarOrHealer != m.Setup.MT.Rules.TemplarOrHealer {
			cx.AbortWithStatus(403)
			return
		}
		// OT
		if n.Setup.OT.Rules.Tank != m.Setup.OT.Rules.Tank {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.OT.Rules.CoercerOrEnchanter != m.Setup.OT.Rules.CoercerOrEnchanter {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.OT.Rules.DirgeOrBard != m.Setup.OT.Rules.DirgeOrBard {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.OT.Rules.ShamanOrHealer != m.Setup.OT.Rules.ShamanOrHealer {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.OT.Rules.ClericOrHealer != m.Setup.OT.Rules.ClericOrHealer {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.OT.Rules.DPSOrHatetransfer != m.Setup.OT.Rules.DPSOrHatetransfer {
			cx.AbortWithStatus(403)
			return
		}

		// DPS1
		if n.Setup.DPS1.Rules.DPS != m.Setup.DPS1.Rules.DPS {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS1.Rules.Bard != m.Setup.DPS1.Rules.Bard {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS1.Rules.DPSOrTank != m.Setup.DPS1.Rules.DPSOrTank {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS1.Rules.Enchanter != m.Setup.DPS1.Rules.Enchanter {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS1.Rules.HealerOrDPS != m.Setup.DPS1.Rules.HealerOrDPS {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS1.Rules.Healer != m.Setup.DPS1.Rules.Healer {
			cx.AbortWithStatus(403)
			return
		}

		// DPS2
		if n.Setup.DPS2.Rules.DPS != m.Setup.DPS2.Rules.DPS {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS2.Rules.Bard != m.Setup.DPS2.Rules.Bard {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS2.Rules.DPSOrTank != m.Setup.DPS2.Rules.DPSOrTank {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS2.Rules.Enchanter != m.Setup.DPS2.Rules.Enchanter {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS2.Rules.HealerOrDPS != m.Setup.DPS2.Rules.HealerOrDPS {
			cx.AbortWithStatus(403)
			return
		}
		if n.Setup.DPS2.Rules.Healer != m.Setup.DPS2.Rules.Healer {
			cx.AbortWithStatus(403)
			return
		}
	}

	if e := h.repo.RaidUpdate(id, n); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, n)
}

func (h *Handler) RaidDelete(cx *gin.Context) {
	id := cx.Param("id")
	if id == "" {
		cx.AbortWithStatus(400)
		return
	}
	kcid := cx.GetHeader("X-Auth-Subject")

	m, e := h.repo.RaidFind(bson.M{"_id": bson.ObjectIdHex(id)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	if kcid != m.KeycloakId {
		cx.AbortWithStatus(403)
		return
	}
	if e := h.repo.RaidRemove(id); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, map[string]interface{}{})
}
