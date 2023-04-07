package db

import (
	"fmt"

	"github.com/jwilyandi19/simple-product/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBConnection(conf helper.DBConfig) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Username, conf.Password, conf.Host, conf.DB)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
