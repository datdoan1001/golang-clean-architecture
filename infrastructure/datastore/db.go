package datastore

import (
	"log"

	"github.com/garaujo/golang-clean-architecture/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewDB function opens the database connection
func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySQLConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mySQLConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
