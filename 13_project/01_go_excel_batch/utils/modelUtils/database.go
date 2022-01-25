package modelUtils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB
var globalDB *gorm.DB

func InitDB(conf *viper.Viper) error {
	var err error
	db, err = gorm.Open("mysql", conf.Get("Mysql.DataSource"))
	if err != nil {
		log.Println("db connection err: ", err)
		return err
	}

	db.DB().SetMaxOpenConns(conf.GetInt("Mysql.MaxOpenConn"))
	db.DB().SetMaxIdleConns(conf.GetInt("Mysql.MaxIdleConn"))

	SetDB(db)

	log.Println("db connection success!")
	return nil
}

func SetDB(db *gorm.DB) {
	globalDB = db
}

func GetDB() *gorm.DB {
	return globalDB
}
