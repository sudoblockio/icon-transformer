package _old

import "github.com/sudoblockio/icon-transformer/routines"

func StartRoutines() {

	////// Address Balance
	//go addressBalanceRoutine()
	//
	//// Governance add preps
	//go addressIsPrep()
	//
	//// Address Type
	//go addressTypeRoutine()
	//
	//// Address Count
	//go addressCountRoutine()
	//
	//// Address Transaction Count
	//go addressTransactionCountRoutine()
	//
	//// Address Transaction Internal Count
	//go addressTransactionInternalCountRoutine()
	//
	//// Address Log Count
	//go addressLogCountRoutine()
	//
	//// Address Token Transfer Count
	//go addressTokenTransferCountRoutine()
	//
	//// Address Token Balance
	//go tokenAddressBalanceRoutine()
	//
	//// Address Token Count
	//go tokenAddressCountRoutine()
	//
	//// Transaction Create Score
	//go transactionCreateScoreRoutine()
	//
	//// Redis Store Count Keys
	//go redisStoreKeysRoutine()
	//
	//// Block count
	//go blockCountRoutine()
	//
	//// Transaction regular count
	//go transactionRegularCountRoutine()
	//
	//// Token transfer count
	//go tokenTransferCountRoutine()

	//go addressIsPrep()
	//time.Sleep(30 * time.Second)
	//go addressTypeRoutine()
	//time.Sleep(30 * time.Second)
	//go addressCountRoutine()
	//time.Sleep(30 * time.Second)
	//go addressTransactionCountRoutine()
	//time.Sleep(30 * time.Second)
	//go addressTransactionInternalCountRoutine()
	//time.Sleep(30 * time.Second)
	//go addressLogCountRoutine()
	//time.Sleep(30 * time.Second)
	//go addressTokenTransferCountRoutine()
	//time.Sleep(30 * time.Second)
	//go tokenAddressBalanceRoutine()
	//time.Sleep(30 * time.Second)
	//go tokenAddressCountRoutine()
	//time.Sleep(30 * time.Second)
	//go transactionCreateScoreRoutine()
	//time.Sleep(30 * time.Second)
	//go redisStoreKeysRoutine()
	//time.Sleep(30 * time.Second)
	//go blockCountRoutine()
	//time.Sleep(30 * time.Second)
	//go transactionRegularCountRoutine()
	//time.Sleep(30 * time.Second)
	//go tokenTransferCountRoutine()
	//time.Sleep(30 * time.Second)
	//go addressBalanceRoutine()

	routines.addressIsPrep()
	routines.addressTypeRoutine()
	// TODO: Rewrite these to not read all redis keys and loop through addresses in chunks
	//  These could be only in recovery
	//addressCountRoutine()
	//addressTransactionCountRoutine()
	//addressTransactionInternalCountRoutine()
	//addressLogCountRoutine()
	//addressTokenTransferCountRoutine()
	//tokenAddressBalanceRoutine()
	routines.tokenAddressCountRoutine()
	transactionCreateScoreRoutine()
	// TODO: Same here -> will oom
	//redisStoreKeysRoutine()
	routines.blockCountRoutine()
	routines.transactionRegularCountRoutine()
	routines.tokenTransferCountRoutine()
	addressBalanceRoutine()
}
