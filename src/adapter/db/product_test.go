package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/silverioTenor/fc-arq-hexagonal-go/adapter/db"
	"github.com/silverioTenor/fc-arq-hexagonal-go/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
		);
	`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("123", "Product 1", 10.0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("123")

	require.Nil(t, err)
	require.Equal(t, "123", product.GetId())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_SaveAndGetAll(t *testing.T) {
	setUp()
	defer Db.Close()

	// ==================
	// Should test CREATE
	// ==================
	productDb := db.NewProductDb(Db)
	newProduct := app.NewProduct()
	newProduct.Name = "Product 2"
	newProduct.Price = 20.0
	newProduct.Status = app.ENABLED

	result, err := productDb.Save(newProduct)
	require.Nil(t, err)

	products, err := productDb.GetAll()

	require.Nil(t, err)
	require.NotNil(t, result)
	require.Len(t, products, 2)
	require.Equal(t, newProduct.GetId(), result.GetId())
	require.Equal(t, newProduct.GetName(), result.GetName())
	require.Equal(t, newProduct.GetPrice(), result.GetPrice())
	require.Equal(t, newProduct.GetStatus(), result.GetStatus())

	// ==================
	// Should test UPDATE
	// ==================
	newProduct.Name = "Product 2 Updated"
	newProduct.Price = 30.0

	_, err = productDb.Save(newProduct)
	require.Nil(t, err)

	products, err = productDb.GetAll()

	require.Nil(t, err)
	require.Len(t, products, 2)
	require.Contains(t, products, newProduct)
	require.Equal(t, products[1].GetName(), "Product 2 Updated")
	require.Equal(t, products[1].GetPrice(), 30.0)
}
