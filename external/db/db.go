package db

import (
	"fmt"

	"github.com/jwilyandi19/simple-product/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLDatabase struct {
	Database *gorm.DB
}

func InitDBConnection(conf helper.DBConfig) (SQLDatabase, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.Username, conf.Password, conf.Host, conf.Port, conf.DB)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return SQLDatabase{}, err
	}
	return SQLDatabase{
		Database: db,
	}, nil
}
