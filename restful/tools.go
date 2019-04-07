package restful

import (
	"dalu/contested/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (h *Handler) JoinRaid(cx *gin.Context) {
	kcid := cx.GetHeader("X-Auth-Subject")

	type Input struct {
		RaidId      string `json:"raid_id"`
		Spot        string `json:"spot"`
		Displayname string `json:"displayname"`
	}
	in := new(Input)
	if e := cx.BindJSON(in); e != nil {
		cx.AbortWithError(400, e)
		return
	}

	raid, e := h.repo.RaidFind(bson.M{"_id": bson.ObjectIdHex(in.RaidId)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}
	switch in.Spot {
	case "mt.tank":
		raid.Setup.MT.Slot1 = in.Displayname
	case "mt.dirge_or_bard":
		raid.Setup.MT.Slot2 = in.Displayname
	case "mt.coercer_or_enchanter":
		raid.Setup.MT.Slot3 = in.Displayname
	case "mt.templar_or_healer":
		raid.Setup.MT.Slot4 = in.Displayname
	case "mt.defiler_or_warder":
		raid.Setup.MT.Slot5 = in.Displayname
	case "mt.swash_or_hatetransfer":
		raid.Setup.MT.Slot6 = in.Displayname
	case "ot.tank":
		raid.Setup.OT.Slot1 = in.Displayname
	case "ot.dirge_or_bard":
		raid.Setup.OT.Slot2 = in.Displayname
	case "ot.coercer_or_enchanter":
		raid.Setup.OT.Slot3 = in.Displayname
	case "ot.shaman_or_healer":
		raid.Setup.OT.Slot4 = in.Displayname
	case "ot.cleric_or_healer":
		raid.Setup.OT.Slot5 = in.Displayname
	case "ot.dps_or_hatetransfer":
		raid.Setup.OT.Slot6 = in.Displayname
	case "dps_1.dps_or_tank":
		raid.Setup.DPS1.Slot1 = in.Displayname
	case "dps_1.dps":
		raid.Setup.DPS1.Slot2 = in.Displayname
	case "dps_1.bard":
		raid.Setup.DPS1.Slot3 = in.Displayname
	case "dps_1.enchanter":
		raid.Setup.DPS1.Slot4 = in.Displayname
	case "dps_1.healer_or_dps":
		raid.Setup.DPS1.Slot5 = in.Displayname
	case "dps_1.healer":
		raid.Setup.DPS1.Slot6 = in.Displayname
	case "dps_2.dps_or_tank":
		raid.Setup.DPS2.Slot1 = in.Displayname
	case "dps_2.dps":
		raid.Setup.DPS2.Slot2 = in.Displayname
	case "dps_2.bard":
		raid.Setup.DPS2.Slot3 = in.Displayname
	case "dps_2.enchanter":
		raid.Setup.DPS2.Slot4 = in.Displayname
	case "dps_2.healer_or_dps":
		raid.Setup.DPS2.Slot5 = in.Displayname
	case "dps_2.healer":
		raid.Setup.DPS2.Slot6 = in.Displayname
	}
	if e := h.repo.RaidUpdate(raid.Id.Hex(), raid); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	n := new(models.RaidNote)
	n.RaidId = in.RaidId
	n.KeycloakId = kcid
	n.Username = "SYSTEM"
	n.Date = time.Now()
	n.Note = in.Displayname + " signed up for the raid"
	if e := h.repo.RaidNoteInsert(n); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, raid)
}

func (h *Handler) LeaveRaid(cx *gin.Context) {
	kcid := cx.GetHeader("X-Auth-Subject")

	type Input struct {
		RaidId string `json:"raid_id"`
		Spot   string `json:"spot"`
	}

	in := new(Input)
	if e := cx.BindJSON(in); e != nil {
		cx.AbortWithError(400, e)
		return
	}

	raid, e := h.repo.RaidFind(bson.M{"_id": bson.ObjectIdHex(in.RaidId)})
	if e != nil {
		cx.AbortWithError(500, e)
		return
	}

	var raider *models.Raider

	switch in.Spot {
	case "mt.tank":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot1})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot1 = ""
	case "mt.dirge_or_bard":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot2})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot2 = ""
	case "mt.coercer_or_enchanter":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot3})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot3 = ""
	case "mt.templar_or_healer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot4})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot4 = ""
	case "mt.defiler_or_warder":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot5})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot5 = ""
	case "mt.swash_or_hatetransfer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.MT.Slot6})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.MT.Slot6 = ""
	case "ot.tank":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot1})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot1 = ""
	case "ot.dirge_or_bard":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot2})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot2 = ""
	case "ot.coercer_or_enchanter":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot3})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot3 = ""
	case "ot.shaman_or_healer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot4})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot4 = ""
	case "ot.cleric_or_healer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot5})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot5 = ""
	case "ot.dps_or_hatetransfer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.OT.Slot6})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.OT.Slot6 = ""
	case "dps_1.dps_or_tank":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot1})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot1 = ""
	case "dps_1.dps":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot2})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot2 = ""
	case "dps_1.bard":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot3})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot3 = ""
	case "dps_1.enchanter":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot4})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot4 = ""
	case "dps_1.healer_or_dps":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot5})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot5 = ""
	case "dps_1.healer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS1.Slot6})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS1.Slot6 = ""
	case "dps_2.dps_or_tank":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot1})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot1 = ""
	case "dps_2.dps":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot2})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot2 = ""
	case "dps_2.bard":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot3})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot3 = ""
	case "dps_2.enchanter":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot4})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot4 = ""
	case "dps_2.healer_or_dps":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot5})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot5 = ""
	case "dps_2.healer":
		raider, e = h.repo.RaiderFind(bson.M{"keycloak_id": kcid, "character.displayname": raid.Setup.DPS2.Slot6})
		if e != nil {
			cx.AbortWithError(400, e)
			return
		}
		raid.Setup.DPS2.Slot6 = ""
	}

	if e := h.repo.RaidUpdate(raid.Id.Hex(), raid); e != nil {
		cx.AbortWithError(400, e)
		return
	}
	n := new(models.RaidNote)
	n.RaidId = in.RaidId
	n.KeycloakId = kcid
	n.Username = "SYSTEM"
	n.Date = time.Now()
	n.Note = raider.Character.Displayname + " unsigned for the raid"
	if e := h.repo.RaidNoteInsert(n); e != nil {
		cx.AbortWithError(500, e)
		return
	}
	cx.SecureJSON(200, raid)
}
