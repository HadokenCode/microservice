package resources

import (
	"errors"
)

var (
	ErrDoesNotExist = errors.New("entity does not exist")
)

type Interface interface {
	Get(id uint64) (interface{}, error)
	Put(id uint64, entity interface{}) error
	Post(entity interface{}) (uint64, error)
	Delete(id uint64) error
}
