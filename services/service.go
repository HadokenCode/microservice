package services

import (
	"net/http"
)

type Interface interface {
	Handler() http.Handler
}
