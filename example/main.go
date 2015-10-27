package main

import (
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/scjudd/microservice/resources"
	"github.com/scjudd/microservice/services"
	"log"
	"net/http"
	"reflect"
)

type Bio struct {
	Name    string `json:"name"`
	Bio     string `json:"bio"`
	Website string `json:"website"`
}

func main() {
	db, err := bolt.Open("bios.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	bucket := []byte("bios")

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		return nil
	})

	svc := &services.CRUD{
		Resource: &resources.Bolt{
			DB:     db,
			Bucket: bucket,
			Type:   reflect.TypeOf(Bio{}),
		},
		Type: reflect.TypeOf(Bio{}),
	}

	r := mux.NewRouter()
	r.HandleFunc("/bios/{id:[0-9]+}", svc.Get).Methods("GET")
	r.HandleFunc("/bios/{id:[0-9]+}", svc.Put).Methods("PUT")
	r.HandleFunc("/bios/", svc.Post).Methods("POST")
	r.HandleFunc("/bios/{id:[0-9]+}", svc.Delete).Methods("DELETE")

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
