package resource

import (
	"encoding/binary"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/scjudd/microservice/gob"
	"reflect"
)

var (
	ErrBucketNotCreated = errors.New("bucket has not been created")
)

type Bolt struct {
	DB     *bolt.DB
	Bucket []byte
	Type   reflect.Type
}

func (res *Bolt) Get(id uint64) (interface{}, error) {
	var entity interface{}
	err := res.DB.View(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket(res.Bucket)
		if b == nil {
			return ErrBucketNotCreated
		}
		g := b.Get(parseKey(id))
		if g == nil {
			return ErrDoesNotExist
		}
		entity, err = gob.Unmarshal(g, res.Type)
		return err
	})
	return entity, err
}

func (res *Bolt) Put(id uint64, entity interface{}) error {
	return res.DB.Update(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket(res.Bucket)
		if b == nil {
			return ErrBucketNotCreated
		}
		g, err := gob.Marshal(entity)
		if err != nil {
			return err
		}
		return b.Put(parseKey(id), g)
	})
}

func (res *Bolt) Post(entity interface{}) (uint64, error) {
	var id uint64
	err := res.DB.Update(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket(res.Bucket)
		if b == nil {
			return ErrBucketNotCreated
		}
		for {
			var err error
			id, err = b.NextSequence()
			if err != nil {
				return err
			}
			if entity := b.Get(parseKey(id)); entity == nil {
				break
			}
		}
		g, err := gob.Marshal(entity)
		if err != nil {
			return err
		}
		return b.Put(parseKey(id), g)
	})
	return id, err
}

func (res *Bolt) Delete(id uint64) error {
	return res.DB.Update(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket(res.Bucket)
		if b == nil {
			return ErrBucketNotCreated
		}
		if entity := b.Get(parseKey(id)); entity == nil {
			return ErrDoesNotExist
		}
		return b.Delete(parseKey(id))
	})
}

func parseKey(u uint64) []byte {
	k := make([]byte, 8)
	binary.BigEndian.PutUint64(k, u)
	return k
}
