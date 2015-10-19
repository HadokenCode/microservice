package json

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func Marshal(entity interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(entity)
	return buf.Bytes(), err
}

func Unmarshal(b []byte, typ reflect.Type) (interface{}, error) {
	entity := reflect.New(typ).Interface()
	dec := json.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(entity)
	return entity, err
}
