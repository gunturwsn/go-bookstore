package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect establishes a connection to a remote server.
// note: using your username & password and create your own database to connect
func Connect() {
	d, err := gorm.Open("mysql", "{username}:{password}@/{database}?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB retrieves a database connection.
func GetDB() *gorm.DB {
	return db
}
