package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
	"time"
)

func TestAddressTxCounts(t *testing.T) {
	config.ReadEnvironment()
	address := &models.Address{Address: "cx216dc90e0bfd732b0e70108ac664aa50907d5cab"}
	GetAddressTxCounts(address)
	crud.GetAddressRoutineCruds()["counts"].LoaderChannel <- address
	time.Sleep(1 * time.Second)
}
