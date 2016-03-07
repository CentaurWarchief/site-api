package links

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage interface {
	FindAll() ([]*Entity, error)
}

func NewStorage(coll *mgo.Collection) Storage {
	return &storage{
		coll: coll,
	}
}

type storage struct {
	coll *mgo.Collection
}

func (s *storage) FindAll() ([]*Entity, error) {
	result := []*Entity{}
	if err := s.coll.Find(bson.M{}).Sort("order").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}
