package database

import (
	"fmt"
	"go-invoice-system/common/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GetEnv("DB_USERNAME"),
		helper.GetEnv("DB_PASSWORD"),
		helper.GetEnv("DB_HOST"),
		helper.GetEnv("DB_PORT"),
		helper.GetEnv("DB_DATABASE"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connect To Database :" + err.Error())
	}

	return DB, err
}
