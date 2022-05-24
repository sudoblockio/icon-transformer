package routines

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

	go addressBalanceRoutine()
	go addressIsPrep()
	go addressTypeRoutine()
	go addressCountRoutine()
	go addressTransactionCountRoutine()
	go addressTransactionInternalCountRoutine()
	go addressLogCountRoutine()
	go addressTokenTransferCountRoutine()
	go tokenAddressBalanceRoutine()
	go tokenAddressCountRoutine()
	go transactionCreateScoreRoutine()
	go redisStoreKeysRoutine()
	go blockCountRoutine()
	go transactionRegularCountRoutine()
	go tokenTransferCountRoutine()
}
