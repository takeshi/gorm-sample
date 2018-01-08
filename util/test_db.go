package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func getDb() *gorm.DB {
	return db
}

func CreateTestDb() *gorm.DB {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("err is %s", err)
		panic("fail to open db ")
	}

	return db
}
