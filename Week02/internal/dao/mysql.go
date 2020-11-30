package dao

import (
	"bytes"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB gorm.DB
var DB *sql.DB

var (
	Username = "root"
	Password = "root"
	Host     = "127.0.0.1"
	Port     = "3310"
	Database = "sf"
)

func init() {
	DB = newDB()
}

func newDB() (orm *sql.DB) {
	// var orm *gorm.DB
	var err error

	mysqlLink := bytes.NewBufferString("")

	mysqlLink.WriteString(Username)
	mysqlLink.WriteString(":" + Password + "@tcp")
	mysqlLink.WriteString("(" + Host)
	mysqlLink.WriteString(":" + Port + ")")
	mysqlLink.WriteString("/" + Database)
	mysqlLink.WriteString("?charset=utf8mb4&parseTime=True&loc=Local&timeout=100ms")

	orm, err = sql.Open("mysql", mysqlLink.String())
	if err != nil {
		log.Printf("database err: %v", err.Error())
	}

	orm.SetMaxOpenConns(100)
	orm.SetMaxIdleConns(10)
	orm.SetConnMaxLifetime(60 * time.Second)

	return
}
