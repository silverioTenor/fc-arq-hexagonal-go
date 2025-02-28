package app

import "errors"

type ProductService struct {
	ProductPersistence IProductPersistence
}

func (s *ProductService) Get(id string) (IProduct, error) {
	return s.ProductPersistence.Get(id)
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Toggle(product IProduct) (IProduct, error) {
	err := error(nil)

	switch product.GetStatus() {
		case ENABLED:
			err = product.Disable()
		case DISABLED:
			err = product.Enable()
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

// func (s *ProductService) Enable(product IProduct) (IProduct, error) {
// 	err := product.Enable()
// 	if err != nil {
// 		return &Product{}, err
// 	}

// 	result, err := s.ProductPersistence.Save(product)
// 	if err != nil {
// 		return &Product{}, err
// 	}

// 	return result, nil
// }

// func (s *ProductService) Disable(product IProduct) (IProduct, error) {
// 	err := product.Enable()
// 	if err != nil {
// 		return &Product{}, err
// 	}

// 	result, err := s.ProductPersistence.Save(product)
// 	if err != nil {
// 		return &Product{}, err
// 	}

// 	return result, nil
// }
