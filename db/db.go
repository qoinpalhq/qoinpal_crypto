package db

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
)

var db *badger.DB

func Init() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("./badgerDB"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

// set value takes two argument key and value and return an error.
// The key,value is of type slice byte
func SetValue(key, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

// This guy takes a key argument and return slice of byte and error
// This will retrive the value using txn.Get
func GetValue(key, value []byte) ([]byte, error) {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		// This func with val would only be called if item.Value encounters no error.
		// according to the documentation
		return item.Value(func(val []byte) error {
			/*
			* @dev  I needed to log something
			* because the files inside the badgerdb is unreadable
			* i even use json marshal to encode it
			* The thing no work
			 */
			fmt.Printf("The answer is: %s\n", val)
			value = append([]byte{}, val...)
			return nil

		})
	})
	return value, err
}
