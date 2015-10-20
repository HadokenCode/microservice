package services

import (
	"github.com/bmizerany/pat"
	"github.com/scjudd/microservice/json"
	"github.com/scjudd/microservice/resources"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type CRUD struct {
	Resource resources.Interface
	Type     reflect.Type
}

func (svc *CRUD) Get(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := svc.Resource.Get(id)
		if err != nil {
			if err == resources.ErrDoesNotExist {
				return &json.Error{404, "entity does not exist", err}
			}
			return &json.Error{500, "problem getting requested entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *CRUD) Put(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		if err := svc.Resource.Put(id, entity); err != nil {
			return &json.Error{500, "problem storing entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *CRUD) Post(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		entity, err := unmarshalBody(r, svc.Type)
		if err != nil {
			return err
		}
		id, err := svc.Resource.Post(entity)
		if err != nil {
			return &json.Error{500, "problem storing entity", err}
		}
		return json.WriteResponse(w, id, entity)
	})
}

func (svc *CRUD) Delete(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, func() error {
		id, err := getID(r)
		if err != nil {
			return err
		}
		if err := svc.Resource.Delete(id); err != nil {
			return &json.Error{500, "problem deleting entity", err}
		}
		return json.WriteResponse(w, id, svc.Type)
	})
}

func (svc *CRUD) Handler() http.Handler {
	m := pat.New()
	m.Get("/:id", http.HandlerFunc(svc.Get))
	m.Put("/:id", http.HandlerFunc(svc.Put))
	m.Post("/", http.HandlerFunc(svc.Post))
	m.Del("/:id", http.HandlerFunc(svc.Delete))
	return m
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
