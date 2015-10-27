package services

import (
	"github.com/scjudd/microservice/errors"
	"github.com/scjudd/microservice/json"
	"github.com/scjudd/microservice/resources"
	"net/http"
	"path"
)

type Collection struct {
	Resource resources.Interface
	Path     string
}

func (svc *Collection) Index(w http.ResponseWriter, r *http.Request) {
	errors.Handler(w, func() error {
		entities := []interface{}{}
		for kv := range svc.Resource.Iter() {
			entity := Entity{kv.Key, path.Join(svc.Path, uint64toa(kv.Key)), kv.Value}
			entities = append(entities, entity)
		}
		data, err := json.Marshal(entities)
		if err != nil {
			return &errors.Error{500, "problem marshalling entities", err}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return nil
	})
}
