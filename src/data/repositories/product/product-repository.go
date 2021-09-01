package product

import (
	"go-auth-module/src/data/database"
	productEntity "go-auth-module/src/data/entity"

	"gorm.io/gorm"
)

var Context *gorm.DB

func LoadDB() {
	Context = database.OpenDatabase()
}

func CreateProduct() {
	LoadDB()
	Context.AutoMigrate(&productEntity.Product{})
	Context.Create(&productEntity.Product{Code: "D42", Price: 100})
}

func FindProduct() productEntity.Product {
	var product productEntity.Product
	Context.First(&product, "code = ?", "D42")

	return product
}
