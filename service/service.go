package service

import (
	"net/http"
)

type Interface interface {
	Get(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
