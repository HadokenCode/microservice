package main

import (
	"github.com/boltdb/bolt"
	"github.com/scjudd/microservice/resource"
	"github.com/scjudd/microservice/service"
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

	svc := &service.Resource{
		Res: &resource.Bolt{
			DB:     db,
			Bucket: bucket,
			Type:   reflect.TypeOf(Bio{}),
		},
	}

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", service.Routes(svc)); err != nil {
		log.Fatal(err)
	}
}
