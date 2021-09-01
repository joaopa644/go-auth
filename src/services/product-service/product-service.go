package productService

import (
	productEntity "go-auth-module/src/data/entity"
	"go-auth-module/src/data/repositories/product"
)

func GetProduct() productEntity.Product {
	product.CreateProduct()
	return product.FindProduct()
}
