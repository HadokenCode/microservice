package service

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func Routes(svc Interface) http.Handler {
	m := pat.New()
	m.Get("/:id", http.HandlerFunc(svc.Get))
	m.Put("/:id", http.HandlerFunc(svc.Put))
	m.Post("/", http.HandlerFunc(svc.Post))
	m.Del("/:id", http.HandlerFunc(svc.Delete))
	return m
}
