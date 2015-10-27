package services

import (
	"github.com/scjudd/microservice/errors"
	"github.com/scjudd/microservice/json"
	"github.com/scjudd/microservice/resources"
	"net/http"
	"path"
	"reflect"
)

type CRUD struct {
	Resource resources.Interface
	Path     string
	Type     reflect.Type
}

type Entity struct {
	Id   uint64      `json:"id"`
	Href string      `json:"href"`
	Data interface{} `json:"data"`
}

}

func (svc *CRUD) Get(w http.ResponseWriter, r *http.Request) {
	errors.Handler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := svc.Resource.Get(id)
		if err != nil {
			if err == resources.ErrDoesNotExist {
				return &errors.Error{404, "entity does not exist", err}
			}
			return &errors.Error{500, "problem getting requested entity", err}
		}
		data, err := json.Marshal(Entity{id, path.Join(svc.Path, uint64toa(id)), entity})
		if err != nil {
			return &errors.Error{500, "problem marshalling requested entity", err}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	})
}

func (svc *CRUD) Put(w http.ResponseWriter, r *http.Request) {
	errors.Handler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		if err := svc.Resource.Put(id, entity); err != nil {
			return &errors.Error{500, "problem storing entity", err}
		}
		data, err := json.Marshal(Entity{id, path.Join(svc.Path, uint64toa(id)), entity})
		if err != nil {
			return &errors.Error{500, "problem marshalling requested entity", err}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	})
}

func (svc *CRUD) Post(w http.ResponseWriter, r *http.Request) {
	errors.Handler(w, func() error {
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		id, err := svc.Resource.Post(entity)
		if err != nil {
			return &errors.Error{500, "problem storing entity", err}
		}
		data, err := json.Marshal(Entity{id, path.Join(svc.Path, uint64toa(id)), entity})
		if err != nil {
			return &errors.Error{500, "problem marshalling requested entity", err}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	})
}

func (svc *CRUD) Delete(w http.ResponseWriter, r *http.Request) {
	errors.Handler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		if err := svc.Resource.Delete(id); err != nil {
			return &errors.Error{500, "problem deleting entity", err}
		}
		data, err := json.Marshal(svc.Type)
		if err != nil {
			return &errors.Error{500, "problem marshalling requested entity", err}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	})
}
