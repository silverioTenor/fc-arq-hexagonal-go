package app_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"githubcom/silverioTenor/fc-arq-hexagonal-go/app"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "P1"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "P1"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	product.Status = app.ENABLED
	err = product.Disable()

	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "P1"
	product.Status = app.DISABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

	product.Id = ""
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the id is required", err.Error())

	product.Id = uuid.NewV4().String()
	product.Name = ""
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Error(t, err)

	product.Name = "P1"
	product.Status = "INVALID"
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.DISABLED
	product.Price = -10
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the price must be greater than or equal to zero", err.Error())
}
