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

	addressBalanceRoutine()
	addressIsPrep()
	addressTypeRoutine()
	addressCountRoutine()
	addressTransactionCountRoutine()
	addressTransactionInternalCountRoutine()
	addressLogCountRoutine()
	addressTokenTransferCountRoutine()
	tokenAddressBalanceRoutine()
	tokenAddressCountRoutine()
	// Transaction Create Score
	transactionCreateScoreRoutine()

	// Redis Store Count Keys
	redisStoreKeysRoutine()

	// Block count
	blockCountRoutine()

	// Transaction regular count
	transactionRegularCountRoutine()

	// Token transfer count
	tokenTransferCountRoutine()
}
