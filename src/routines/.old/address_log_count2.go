package _old

import (
	"strconv"

	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func addressLogCountRoutine2() {

	valueInt, _ := strconv.Atoi(valueString)
	address := &models.Address{
		Address:  addressString,
		LogCount: int64(valueInt),
	}

	crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "log_count"})
}
