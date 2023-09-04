package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() (*gorm.DB, error) {

	Db_user_name := os.Getenv("DB_USERNAME")
	Db_password := os.Getenv("DB_PASSWORD")
	Db_name := os.Getenv("DB_NAME")
	//Db_port := os.Getenv("DB_PORT")

	dsn := Db_user_name + ":" + Db_password + "@tcp(localhost:3306)/" + Db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
