package utils

import (
	"antrian/internal/database"

	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func InitConnection() *Connection {

	return &Connection{
		DB: database.DBMySQLConnecting(),
	}

}
