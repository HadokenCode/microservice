package service

import (
	"github.com/scjudd/microservice/json"
	"github.com/scjudd/microservice/resource"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type Resource struct {
	Res  resource.Interface
	Type reflect.Type
}

func (svc *Resource) Get(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := svc.Res.Get(id)
		if err != nil {
			if err == resource.ErrDoesNotExist {
				return &json.Error{404, "entity does not exist", err}
			}
			return &json.Error{500, "problem getting requested entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *Resource) Put(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		if err := svc.Res.Put(id, entity); err != nil {
			return &json.Error{500, "problem storing entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *Resource) Post(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		id, err := svc.Res.Post(entity)
		if err != nil {
			return &json.Error{500, "problem storing entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *Resource) Delete(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		if err := svc.Res.Delete(id); err != nil {
			return &json.Error{500, "problem deleting entity", err}
		}
		return json.WriteResponse(w, id, svc.Type)
	})
}

func getID(r *http.Request) (uint64, error) {
	id, err := strconv.ParseUint(r.URL.Query().Get(":id"), 10, 64)
	if err != nil {
		return id, &json.Error{400, "missing or bad 'id' parameter", err}
	}
	return id, nil
}

func unmarshalBody(r *http.Request, typ reflect.Type) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &json.Error{500, "problem reading request body", err}
	}
	defer r.Body.Close()
	entity, err := json.Unmarshal(body, typ)
	if err != nil {
		return nil, &json.Error{500, "problem unmarshalling entity", err}
	}
	return entity, nil
}
