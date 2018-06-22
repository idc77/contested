package models

import (
	"github.com/globalsign/mgo/bson"
)

type Raider struct {
	Id         bson.ObjectId    `bson:"_id,omitempty" json:"id"`
	KeycloakId string           `bson:"keycloak_id" json:"keycloak_id"`
	Active     bool             `bson:"active" json:"active"`
	Character  *CensusCharacter `bson:"character" json:"character"`
}

/*
func (r *Raider) WireURLToCensusId() {
	r.CensusId = r.WireURL[strings.LastIndex(r.WireURL, "/")+1:]
}
*/

// https://u.eq2wire.com/soe/character_detail/2611343626022
