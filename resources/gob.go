package resources

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

func gobMarshal(entity interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(entity)
	return buf.Bytes(), err
}

func gobUnmarshal(b []byte, typ reflect.Type) (interface{}, error) {
	entity := reflect.New(typ).Interface()
	dec := gob.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(entity)
	return entity, err
}
