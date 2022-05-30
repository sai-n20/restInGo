package database

import (
	"fmt"
)

type InMemDB interface {
	Create(key string, val interface{})

	Read(key string) interface{}

	RetrieveAllData() map[string]interface{}
}

type memDB struct {
	keyValStore map[string]interface{}
}

func CreateInMemDB() *memDB {
	fmt.Println("Memorystore reset and initialized")
	return &memDB{make(map[string]interface{})}
}

func (db *memDB) Create(key string, val interface{}) {
	db.keyValStore[key] = val
	fmt.Println("Insertion operation complete, key for the entire payload is- ", key)
}

func (db *memDB) RetrieveAllData() map[string]interface{} {

	return db.keyValStore
}

func (db *memDB) Read(key string) interface{} {
	if val, ok := db.keyValStore[key]; ok {
		return val
	}
	return nil
}
