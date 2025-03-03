package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/adapter/db"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app/service"
)

func main() {
	dbConn, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(dbConn)
	productService := service.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10)
	// product, _ := productService.Get("c1e5fc2d-f660-4f0b-a2a9-58dd9318a81a")

	_, err := productService.Toggle(product)

	if (err != nil) {
		fmt.Println(err)
	}
}