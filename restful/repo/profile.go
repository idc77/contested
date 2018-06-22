package repo

import (
	"dalu/contested/models"

	"github.com/globalsign/mgo/bson"
)

func (r *Repo) ProfileFindAll(q bson.M) ([]*models.Profile, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.profileCollection)
	var m []*models.Profile
	if e := c.Find(q).All(&m); e != nil {
		return m, e
	}
	return m, nil
}

func (r *Repo) ProfileFind(q bson.M) (*models.Profile, error) {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.profileCollection)
	m := new(models.Profile)
	if e := c.Find(q).One(m); e != nil {
		return nil, e
	}
	return m, nil

}

func (r *Repo) ProfileInsert(m *models.Profile) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.profileCollection)
	return c.Insert(m)

}

func (r *Repo) ProfileUpdate(id string, m *models.Profile) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.profileCollection)
	return c.UpdateId(bson.ObjectIdHex(id), m)
}

func (r *Repo) ProfileRemove(id string) error {
	ms := r.ms.Copy()
	defer ms.Close()
	c := ms.DB(r.database).C(r.profileCollection)
	return c.RemoveId(bson.ObjectIdHex(id))
}
