package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tbpathirana/go-ikea/pkg/config"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name         string `json:"name"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	Location     string `json:"location"`
	Availability string `json:"availability"`
}

//https://v1.gorm.io/docs/query.html

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetProduct() []Product {
	var product []Product
	db.Find(&product)
	return product
}

func GetProductById(i int64) (*Product, *gorm.DB) {
	var product Product
	db.Where("ID = ?", i).First(&product)
	return &product, db
}

func DeleteProductById(id int64) Product {
	var product Product
	db.Where("ID = ?", id).Delete(product)
	return product
}
