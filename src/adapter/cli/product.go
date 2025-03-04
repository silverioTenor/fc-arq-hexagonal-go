package cli

import (
	"fmt"

	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app"
)

func Run(
	service app.IProductService,
	action string,
	productId string,
	productName string,
	price float64,
) (string, error) {
	var result = ""

	switch action {
		case "create":
			product, err := service.Create(productName, price)
			
			if (err != nil) {
				return result, err
			}

			result = fmt.Sprintf("Product ID %s with name %s has been created with the price %.2f and status %s",
				product.GetId(),
				product.GetName(),
				product.GetPrice(),
				product.GetStatus(),
			)
		case "enable":
			product, err := service.Get(productId)

			if (err != nil) {
				return result, err
			}

			response, err := service.Toggle(product)

			if (err != nil) {
				return result, err
			}

			result = fmt.Sprintf("Product %s has been enabled", response.GetName())
		case "disable":
			product, err := service.Get(productId)

			if (err != nil) {
				return result, err
			}

			response, err := service.Toggle(product)

			if (err != nil) {
				return result, err
			}

			result = fmt.Sprintf("Product %s has been disabled", response.GetName())
		default:
			response, err := service.Get(productId)

			if (err != nil) {
				return result, err
			}

			result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %.2f\nStatus: %s",
				response.GetId(),
				response.GetName(),
				response.GetPrice(),
				response.GetStatus(),
			)
	}

	return result, nil
}
