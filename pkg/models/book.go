package models

import (
	"simple-rest/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	//Id          string `json:"id"`
	Name      string `gorm:" "`
	About     string
	Price     string

}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (b *Product) CreateProduct() *Product {
	db.NewRecord(b)
	db.LogMode(true).Create(&b)
	return b
}

func  GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func GetProductById(Id int64) (*Product , *gorm.DB){
	var getProduct Product
	db:=db.Where("ID = ?", Id).Find(&getProduct)
	return &getProduct, db
}

func DeleteBook(ID int64) Product {
	var product Product
	db.Where("ID = ?", ID).Delete(product)
	return product
}