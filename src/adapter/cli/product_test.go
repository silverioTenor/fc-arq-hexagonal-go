package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/adapter/cli"
	mock_app "github.com/silverioTenor/fc-arq-hexagonal-go/src/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 19.90
	productStatus := "enabled"
	productId := "abc123"

	productMock := mock_app.NewMockIProduct(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_app.NewMockIProductService(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Toggle(gomock.Any()).Return(productMock, nil).AnyTimes()

	/**
	* =====================================================
	* ================= TEST CLI - CREATE =================
	* =====================================================
	*/
	resultExpected := fmt.Sprintf("Product ID %s with name %s has been created with the price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	/**
	* =====================================================
	* ================= TEST CLI - ENABLE =================
	* =====================================================
	*/
	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(serviceMock, "enable", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	/**
	* =====================================================
	* ================= TEST CLI - DISABLE ================
	* =====================================================
	*/
	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(serviceMock, "disable", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	/**
	* =====================================================
	* ================= TEST CLI - DEFAULT ================
	* =====================================================
	*/
	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(serviceMock, "abc", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}