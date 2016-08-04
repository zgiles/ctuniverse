package universestore

import (
	"github.com/zgiles/ctuniverse"
)

type StoreI interface {
	PutObject(o ctuniverse.UniverseObject) (error)
	GetObject(key string) (ctuniverse.UniverseObject, error)
}

type StoreDBI interface {
	PutObject(o ctuniverse.UniverseObject) (error)
	GetObject(key string) (ctuniverse.UniverseObject, error)
}

type store struct {
 	db StoreDBI
}

func (local store) PutObject(o ctuniverse.UniverseObject) (error) {
	return local.db.PutObject(o)
}

func (local store) GetObject(key string) (ctuniverse.UniverseObject, error) {
	return local.db.GetObject(key)
}

func New(db1 StoreDBI) StoreI {
  return &store{db1}
}
