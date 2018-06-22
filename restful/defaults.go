package restful

import (
	"dalu/contested/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DefaultSetup(cx *gin.Context) {
	cx.SecureJSON(200, util.DefaultSetup())
}

func (h *Handler) DefaultRaid(cx *gin.Context) {
	setup := util.DefaultSetup()
	raid := util.DefaultRaid()
	raid.Setup = setup
	cx.SecureJSON(200, raid)
}

func (h *Handler) RaidClassList(cx *gin.Context) {
	c := util.RaidClasses()
	cx.SecureJSON(200, c)
}
