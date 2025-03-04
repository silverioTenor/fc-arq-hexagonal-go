package service

import (
	"errors"

	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app"
)

type ProductService struct {
	ProductPersistence app.IProductPersistence
}

func NewProductService(productPersistence app.IProductPersistence) *ProductService {
	return &ProductService{
		ProductPersistence: productPersistence,
	}
}

func (s *ProductService) Get(id string) (app.IProduct, error) {
	return s.ProductPersistence.Get(id)
}

func (s *ProductService) Create(name string, price float64) (app.IProduct, error) {
	product := app.NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()

	if err != nil {
		return &app.Product{}, err
	}

	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return &app.Product{}, err
	}

	return result, nil
}

func (s *ProductService) Toggle(product app.IProduct) (app.IProduct, error) {
	err := error(nil)

	switch product.GetStatus() {
		case app.ENABLED:
			err = product.Enable()
		case app.DISABLED:
			err = product.Disable()
		default:
			return nil, errors.New("the status must be enabled or disabled")
	}

	if err != nil {
		return nil, err
	}

	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
