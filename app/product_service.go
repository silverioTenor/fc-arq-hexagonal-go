package app

type ProductService struct {
	ProductPersistence IProductPersistence
}

func (s *ProductService) Get(id string) (IProduct, error) {
	return s.ProductPersistence.Get(id)
}