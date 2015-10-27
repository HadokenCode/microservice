package services

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scjudd/microservice/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func uint64toa(i uint64) string {
	return strconv.Itoa(int(i))
}

func getID(r *http.Request) (uint64, error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, &errors.Error{400, "missing 'id' parameter", nil}
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return id, &errors.Error{400, "missing or bad 'id' parameter", err}
	}
	return id, nil
}

func unmarshalBody(r *http.Request, typ reflect.Type) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &errors.Error{500, "problem reading request body", err}
	}
	defer r.Body.Close()
	entity := reflect.New(typ).Interface()
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(entity); err != nil {
		return nil, &errors.Error{500, "problem unmarshalling entity", err}
	}
	return entity, nil
}
