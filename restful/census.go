package restful

import (
	"dalu/contested/models"

	"strings"

	"dalu/contested/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CensusSearch(cx *gin.Context) {
	type Input struct {
		Name       string `json:"name,omitempty"`
		World      string `json:"world,omitempty"`
		URL        string `json:"url,omitempty"`
		SearchType string `json:"search_type"`
	}

	in := new(Input)

	if e := cx.BindJSON(in); e != nil {
		cx.AbortWithError(400, e)
		return
	}

	params := new(models.CensusParams)
	search_name_world := false
	search_id := false

	if in.URL != "" {
		search_name_world = true
		params.Id = in.URL[strings.LastIndex(in.URL, "/")+1:]
	}

	if in.Name != "" && in.World != "" {
		search_id = true
		switch in.SearchType {
		case "prefix":
			params.NameFirstLower = "^" + strings.ToLower(in.Name)
		case "all":
			params.NameFirstLower = "*" + strings.ToLower(in.Name)
		default:
			params.NameFirstLower = strings.ToLower(in.Name)
		}

		params.World = in.World
	}

	if search_id || search_name_world {
		m, e := util.CensusRequest(params)
		if e != nil {
			cx.AbortWithError(500, e)
			return
		}
		cx.SecureJSON(200, m)
		return
	}
	cx.AbortWithStatus(400)
}
