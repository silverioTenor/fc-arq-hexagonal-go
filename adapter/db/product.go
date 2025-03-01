package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silverioTenor/fc-arq-hexagonal-go/app"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (app.IProduct, error) {
	var product app.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	stmt.Close()

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) GetAll() ([]app.IProduct, error) {
	var products []app.IProduct
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product app.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Status)

		if (err != nil) {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductDb) Save(product app.IProduct) (app.IProduct, error) {
   _, err := p.Get(product.GetId())

   if err != nil {
      return p.create(product)
   }

   return p.update(product)
}

func (p *ProductDb) create(product app.IProduct) (app.IProduct, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
   
	if err != nil {
      return nil, err
	}

   err = stmt.Close()

   if err != nil {
      return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product app.IProduct) (app.IProduct, error) {
   stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")

   if err != nil {
      return nil, err
   }

   _, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())

   if err != nil {
      return nil, err
   }

   err = stmt.Close()

   if err != nil {
      return nil, err
   }

   return product, nil
}