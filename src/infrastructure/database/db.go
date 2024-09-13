package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(user string, pass string, host string, port string, name string) {

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", host, user, pass, name, port)
	_db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = _db
}

func MigrateSchema() {

}
