package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/gotoolkit/db/config"
)

var db *gorm.DB
var err error

func InitialDB(dbConfig *config.DBConfig) {

	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset,
	)
	db, err = gorm.Open(dbConfig.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("db is nil")
	}
	return db
}
