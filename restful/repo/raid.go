package repo

import (
	"dalu/contested/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (r *Repo) RaidFindAll(q bson.M) ([]*models.Raid, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	var m []*models.Raid
	if e := c.Find(q).Sort("date_start").All(&m); e != nil {
		return m, e
	}
	return m, nil
}

func (r *Repo) RaidFind(q bson.M) (*models.Raid, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	m := new(models.Raid)
	if e := c.Find(q).One(m); e != nil {
		return nil, e
	}
	return m, nil
}

func (r *Repo) RaidInsert(m *models.Raid) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	return c.Insert(m)
}

func (r *Repo) RaidUpdate(id string, m *models.Raid) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	return c.UpdateId(bson.ObjectIdHex(id), m)
}

func (r *Repo) RaidRemove(id string) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	return c.RemoveId(bson.ObjectIdHex(id))
}

func (r *Repo) RaidWatch() (*mgo.ChangeStream, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raidCollection)
	pipeline := []bson.M{}
	return c.Watch(pipeline, mgo.ChangeStreamOptions{})
}

func (r *Repo) RaidWatchClose(changeStream *mgo.ChangeStream) error {
	if e := changeStream.Close(); e != nil {
		return e
	}
	return nil
}
