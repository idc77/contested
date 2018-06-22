package repo

import (
	"dalu/contested/models"

	"github.com/globalsign/mgo/bson"
)

func (r *Repo) RaiderFindAll(q bson.M) ([]*models.Raider, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raiderCollection)
	var m []*models.Raider
	if e := c.Find(q).All(&m); e != nil {
		return m, e
	}
	return m, nil
}

func (r *Repo) RaiderFind(q bson.M) (*models.Raider, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raiderCollection)
	m := new(models.Raider)
	if e := c.Find(q).One(m); e != nil {
		return nil, e
	}
	return m, nil
}

func (r *Repo) RaiderInsert(m *models.Raider) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raiderCollection)
	return c.Insert(m)
}

func (r *Repo) RaiderUpdate(id string, m *models.Raider) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raiderCollection)
	return c.UpdateId(bson.ObjectIdHex(id), m)
}

func (r *Repo) RaiderRemove(id string) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.raiderCollection)
	return c.RemoveId(bson.ObjectIdHex(id))
}
