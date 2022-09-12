package crud

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
	"time"
)

func TestGetAddressCruds(t *testing.T) {
	config.ReadEnvironment()

	address := &models.Address{
		Address:          "hxe694b4744cb1f5c13f381f3c4c94e05e74759e2c",
		TransactionCount: 1,
	}
	GetAddressCrud().LoaderChannel <- address
	time.Sleep(1 * time.Second)
}

func TestGetAddressRoutineCruds(t *testing.T) {
	config.ReadEnvironment()

	address := &models.Address{
		Address:          "hxe694b4744cb1f5c13f381f3c4c94e05e74759e2c",
		TransactionCount: 1,
	}
	GetAddressRoutineCruds()

	AddressRoutineCruds["counts"].LoaderChannel <- address

	//GetAddressRoutineCruds()["counts"].LoaderChannel <- address
	time.Sleep(10 * time.Second)
}
