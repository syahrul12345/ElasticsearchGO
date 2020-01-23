package models

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	fmt.Println("INIT")
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		fmt.Println("failed to get cache db")
	}
	defer db.Close()
}

// GetDB gets the levelDB instance
func GetDB() *leveldb.DB {
	return db
}
