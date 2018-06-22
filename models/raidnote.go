package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type RaidNote struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	KeycloakId string        `bson:"keycloak_id" json:"keycloak_id"`
	Username   string        `bson:"username" json:"username"`
	RaidId     string        `bson:"raid_id" json:"raid_id"`
	Date       time.Time     `bson:"date" json:"date"`
	Note       string        `bson:"note" json:"note"`
}
