package application_test

import (
	"testing"

	"github.com/dgryski/trifles/uuid"
	"github.com/joaotavioos/hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduc_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())

}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()

	require.Equal(t, "The price must be zero in order to have product disabled", err.Error())

}

func Test_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.UUIDv4()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "Status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -25
	_, err = product.IsValid()
	require.Equal(t, "Price must be greater or equal to zero", err.Error())

}
