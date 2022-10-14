package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:visioN1995@/ikeadb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		fmt.Println("Error connecting to the DB")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
