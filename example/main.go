package main

import (
	"github.com/boltdb/bolt"
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

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := &services.CRUD{
		Resource: &resources.Bolt{
			DB:     db,
			Bucket: bucket,
			Type:   reflect.TypeOf(Bio{}),
		},
		Type: reflect.TypeOf(Bio{}),
	}

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", svc.Handler()); err != nil {
		log.Fatal(err)
	}
}
