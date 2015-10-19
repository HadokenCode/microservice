package service

import (
	"github.com/scjudd/microservice/json"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, h func() error) {
	if err := h(); err != nil {
		serr, ok := err.(*json.Error)
		if !ok {
			serr = &json.Error{500, "internal server error", err}
		}
		data, err := json.Marshal(serr)
		if err != nil {
			serr = &json.Error{500, "problem marshalling json error", err}
		}
		if serr.Err != nil {
			log.Println(serr.Error())
		}
		w.WriteHeader(serr.Status)
		w.Write(data)
	}
}
