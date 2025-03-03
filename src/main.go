package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/adapter/db"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app/service"
)

func main() {
	dbConn, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db.NewProductDb(dbConn)
	productService := service.NewProductService(productDbAdapter)
	productService.Create("Product 1", 10)
}