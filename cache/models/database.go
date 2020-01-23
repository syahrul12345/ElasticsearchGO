package models

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	db, _ = leveldb.OpenFile("./db", nil)
	defer db.Close()
}

// GetDB gets the levelDB instance
func GetDB() *leveldb.DB {
	return db
}
