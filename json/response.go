package json

import (
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	ID      uint64      `json:"id"`
	Data    interface{} `json:"data"`
}

func WriteResponse(w http.ResponseWriter, id uint64, entity interface{}) error {
	data, err := Marshal(Response{200, "ok", id, entity})
	if err != nil {
		return &Error{500, "problem marshalling response", err}
	}
	w.Write(data)
	return nil
}
