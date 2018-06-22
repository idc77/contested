package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Raid struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title"`
	KeycloakId  string        `bson:"keycloak_id" json:"keycloak_id"`
	World       string        `bson:"world" json:"world"`
	DateMeet    time.Time     `bson:"date_meet" json:"date_meet"`
	DateStart   time.Time     `bson:"date_start" json:"date_start"`
	DateEnd     time.Time     `bson:"date_end" json:"date_end"`
	Description string        `bson:"description" json:"description"`
	Setup       *RaidSetup    `bson:"setup" json:"setup"`
}
