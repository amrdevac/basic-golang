package database

import (
	viperconfig "antrian/internal/viperConfig"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conf(val string) string {
	return viperconfig.GetConfig(val)
}
func DBMySQLConnecting() *gorm.DB {
	str_conn := "{username}:{password}@tcp({host}:{port})/{name}?charset=utf8&parseTime=True&loc=Local"

	str_conn = strings.ReplaceAll(str_conn, "{username}", conf("db.username"))
	str_conn = strings.ReplaceAll(str_conn, "{password}", conf("db.password"))
	str_conn = strings.ReplaceAll(str_conn, "{host}", conf("db.host"))
	str_conn = strings.ReplaceAll(str_conn, "{port}", conf("db.port"))
	str_conn = strings.ReplaceAll(str_conn, "{name}", conf("db.name"))
	fmt.Println(str_conn)

	db, err := gorm.Open(mysql.Open(str_conn))

	if err != nil {
		log.Fatal(err)
	}
	return db

}
