package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app"
	mock_app "github.com/silverioTenor/fc-arq-hexagonal-go/src/app/mocks"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app/service"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}

func TestProductService_ToggleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	product.EXPECT().GetStatus().Return("").AnyTimes()

	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Toggle(product)
	require.Error(t, err, "the status must be enabled or disabled")
	require.Nil(t, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	product.EXPECT().GetStatus().Return(app.ENABLED).AnyTimes()

	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Toggle(product)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()
	product.EXPECT().GetStatus().Return(app.DISABLED).AnyTimes()

	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		ProductPersistence: persistence,
	}

	result, err := service.Toggle(product)
	require.Nil(t, err)
	require.NotNil(t, result)
	require.Equal(t, product, result)
}
