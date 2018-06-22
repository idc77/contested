package util

import (
	"dalu/contested/models"
	"encoding/json"
	"os"
	"testing"
)

func TestCensusRequest(t *testing.T) {
	p := new(models.CensusParams)
	p.NameFirstLower = "loriega"
	p.World = "Fallen Gate"

	m, e := CensusRequest(p)
	if e != nil {
		t.Error(e)
		return
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(m.CharacterList[0])
}
