package models

import "github.com/globalsign/mgo/bson"

type Profile struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	KeycloakId string        `bson:"keycloak_id" json:"keycloak_id"`
}
