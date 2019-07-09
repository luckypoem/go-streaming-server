package orm

import (
	"go-streaming-server/conf"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Gorm *gorm.DB

func init() {
	var err error

	config, err := conf.LoadConfigFromFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	Gorm, err = gorm.Open("mysql", config.MySQLDSN)

	if err != nil {
		log.Fatal(err)
	}
}
