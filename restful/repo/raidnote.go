package repo

import (
	"dalu/contested/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (r *Repo) RaidNoteFindAll(q bson.M) ([]*models.RaidNote, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	var m []*models.RaidNote
	if e := c.Find(q).Sort("-date").All(&m); e != nil {
		return m, e
	}
	return m, nil
}

func (r *Repo) RaidNoteFind(q bson.M) (*models.RaidNote, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	m := new(models.RaidNote)
	if e := c.Find(q).One(m); e != nil {
		return nil, e
	}
	return m, nil
}

func (r *Repo) RaidNoteInsert(m *models.RaidNote) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	return c.Insert(m)
}

func (r *Repo) RaidNoteUpdate(id string, m *models.RaidNote) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	return c.UpdateId(bson.ObjectIdHex(id), m)
}

func (r *Repo) RaidNoteRemove(id string) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	return c.RemoveId(bson.ObjectIdHex(id))
}

func (r *Repo) RaidNotesRemove(q bson.M) (*mgo.ChangeInfo, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidNoteCollection)
	return c.RemoveAll(q)
}
