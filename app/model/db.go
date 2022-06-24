package model

import (
	"fmt"

	"github.com/0l1v3rr/todo/app/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() error {
	// creating the dsn from the environment variables
	dsn := fmt.Sprintf(
		"%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		data.Env["MYSQL_USERNAME"],
		data.Env["MYSQL_DOMAIN"],
		data.Env["MYSQL_PORT"],
		data.Env["MYSQL_DATABASE"],
	)

	// opening a gorm connection
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// migrating the models
	DB.AutoMigrate(&User{})

	return nil
}
