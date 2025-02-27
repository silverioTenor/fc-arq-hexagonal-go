package app_test

import (
	"githubcom/silverioTenor/fc-arq-hexagonal-go/app"
	mock_app "githubcom/silverioTenor/fc-arq-hexagonal-go/app/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)
func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}