package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, h func() error) {
	if err := h(); err != nil {
		serr, ok := err.(*Error)
		if !ok {
			serr = &Error{500, "internal server error", err}
		}
		data, err := json.Marshal(serr)
		if err != nil {
			serr = &Error{500, "problem marshalling error", err}
		}
		if serr.Err != nil {
			log.Println(serr.Error())
		}
		w.WriteHeader(serr.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
