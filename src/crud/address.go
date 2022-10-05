package crud

import (
	"sync"
	"time"

	"github.com/sudoblockio/icon-transformer/models"
)

var addressCrudOnce sync.Once

var AddressCrud *Crud[models.Address, models.AddressORM]
var AddressContractCrud *Crud[models.Address, models.AddressORM]

// GetAddressCrud - create and/or return the crud object
func GetAddressCrud() *Crud[models.Address, models.AddressORM] {
	addressCrudOnce.Do(func() {
		AddressCrud = GetCrud(models.Address{}, models.AddressORM{})

		// There is no loader channel for generic addresses -> only contracts
		AddressCrud.Migrate()

		// For main transformer - requires tuning due to many duplicate records
		AddressContractCrud = GetCrud(models.Address{}, models.AddressORM{})
		AddressContractCrud.columns = []string{"address", "is_contract"}
		//AddressContractCrud.dbBufferWait = 1 * time.Millisecond
		AddressContractCrud.MakeStartLoaderChannel()
	})

	return AddressCrud
}

var addressRoutinesCrudOnce sync.Once

type RoutineCrud struct {
	Name    string
	Columns []string
}

var AddressRoutineCruds = make(map[string]*Crud[models.Address, models.AddressORM])

// GetAddressRoutineCruds - create and/or return the crud object for routines
func GetAddressRoutineCruds() map[string]*Crud[models.Address, models.AddressORM] {
	addressRoutinesCrudOnce.Do(func() {
		for _, v := range []RoutineCrud{
			{
				Columns: []string{"address", "balance"},
				Name:    "address_balance",
			},
			{
				Columns: []string{"address", "is_prep"},
				Name:    "address_is_prep",
			},
			{
				Columns: []string{"address", "is_type"},
				Name:    "address_is_type",
			},
			{
				Columns: []string{"address", "token_contract_address", "balance"},
				Name:    "address_token_contract_balance",
			},
			{
				Columns: []string{
					"address",
					"transaction_count",
					"transaction_internal_count",
					"token_transfer_count",
					"log_count",
				},
				Name: "counts",
			},
		} {
			AddressRoutineCruds[v.Name] = GetCrud(models.Address{}, models.AddressORM{})
			AddressRoutineCruds[v.Name].dbBufferWait = 250 * time.Millisecond
			AddressRoutineCruds[v.Name].columns = v.Columns
			AddressRoutineCruds[v.Name].metrics.Name = v.Name + "_routine"
			AddressRoutineCruds[v.Name].MakeStartLoaderChannel()
		}
	})

	return AddressRoutineCruds
}

func InitAddressCrud() {
	GetAddressCrud()
}

//func (m *Crud[M, O]) addressBatchErrorHandler(b []*M, cols []string) error {
//	b = m.RemoveDuplicatePrimaryKeys(b)
//	return m.retryBatchLoader(b, m.UpsertMany, []string{""}, 0)
//}

//func (m *Crud[M, O]) RemoveDuplicatePrimaryKeys(batch []*M) []*M {
//	var Model *M
//	allKeys := make(map[*M]bool)
//	list := []*M{}
//
//	for _, item := range batch {
//		for _, pkey := range m.primaryKeys {
//			value := reflect.ValueOf(item).FieldByName(pkey.Name)
//			reflect.ValueOf(Model).Elem().FieldByName(pkey.Name).Set(reflect.ValueOf(value))
//		}
//
//		if _, value := allKeys[Model]; !value {
//			allKeys[item] = true
//			list = append(list, item)
//		}
//	}
//	return list
//}
