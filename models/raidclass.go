package models

type RaidClass struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}
