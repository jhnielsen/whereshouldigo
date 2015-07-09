package main

import (
	"fmt"
	"net/http"
	"github.com/boltdb/bolt"
)

//todo: see http://stackoverflow.com/a/17384204
var db *bolt.DB

type Message struct {
	Name string
	Body string
	Time int64
}

func handler(w http.ResponseWriter, r *http.Request) {
    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("MyBucket"))
        err := b.Put([]byte("answer"), []byte(r.URL.Path))
        return err
    })

}

func loadEntry(w http.ResponseWriter, r *http.Request) {
//    db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte("MyBucket"))
//		v := b.Get([]byte("answer"))
//		fmt.Printf("The answer is: %s\n", v)
//		return nil
//	})
	db.Close()
}

func main() {
	db,_ := bolt.Open("my.db", 0600, nil)

    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucket([]byte("MyBucket"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
    defer db.Close()

	http.HandleFunc("/save", handler)
	http.HandleFunc("/view", loadEntry)
	http.ListenAndServe(":8080", nil)

}
