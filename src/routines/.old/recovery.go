package _old

//import (
//	"github.com/sudoblockio/icon-transformer/models"
//	"github.com/sudoblockio/icon-transformer/utils"
//)
//
//var addressRoutines = []func(a *models.Address){
//	setAddressBalances,
//	setAddressTxCounts,
//}
//
//var tokenAddressRoutines = []func(t *models.TokenAddress){
//	setTokenAddressBalances,
//	//setTokenAddressTxCounts,
//}
//
//func StartRecovery() {
//	// By address
//	utils.AddressGoRoutines(addressRoutines)
//	utils.TokenAddressGoRoutines(tokenAddressRoutines)
//
//	// One shot
//	addressTypeRoutine()
//	countAddressesToRedisRoutine()
//}
