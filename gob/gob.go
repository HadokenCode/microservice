package gob

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

func Marshal(entity interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(entity)
	return buf.Bytes(), err
}

func Unmarshal(b []byte, typ reflect.Type) (interface{}, error) {
	entity := reflect.New(typ).Interface()
	dec := gob.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(entity)
	return entity, err
}
