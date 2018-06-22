package repo

import "github.com/globalsign/mgo"

type Repo struct {
	ms                 *mgo.Session
	database           string
	profileCollection  string
	raidCollection     string
	raiderCollection   string
	raidNoteCollection string
}

func NewRepo(ms *mgo.Session) *Repo {
	return &Repo{
		ms:                 ms.Clone(),
		database:           "contested",
		profileCollection:  "profiles",
		raidCollection:     "raids",
		raiderCollection:   "raiders",
		raidNoteCollection: "raidnotes",
	}
}

func (r *Repo) Close() {
	r.ms.Close()
}
