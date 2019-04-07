package util

import (
	"dalu/contested/models"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func CensusRequest(params *models.CensusParams) (*models.CensusResult, error) {
	var addr *url.URL
	addr, e := url.Parse("http://census.daybreakgames.com")
	if e != nil {
		return nil, e
	}

	addr.Path += "/s:contested/get/eq2/character/"
	parameters := url.Values{}

	if params.Id != "" {
		parameters.Add("id", params.Id)
	} else {
		parameters.Add("c:limit", "20")
		if params.World != "" {
			parameters.Add("locationdata.world", params.World)
		}
		if params.NameFirst != "" {
			parameters.Add("name.first", params.NameFirst)
		}
		if params.NameFirstLower != "" {
			parameters.Add("name.first_lower", params.NameFirstLower)
		}
	}

	addr.RawQuery = parameters.Encode()

	qs := strings.Replace(addr.String(), "%2A", "*", -1)

	rsp, e := http.Get(qs)
	if e != nil {
		return nil, e
	}
	defer rsp.Body.Close()
	m := new(models.CensusResult)
	dec := json.NewDecoder(rsp.Body)
	if e := dec.Decode(m); e != nil {
		return nil, e
	}
	return m, nil
}

// https://census.daybreakgames.com/s:eq2wire/img/eq2/character/2611343626022/headshot
