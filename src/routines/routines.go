package routines

func StartRoutines() {

	// Address Balance
	go addressBalanceRoutine()

	// Governance add preps
	go addressIsPrep()

	// Address Type
	go addressTypeRoutine()

	// Address Count
	go addressCountRoutine()

	// Address Transaction Count
	go addressTransactionCountRoutine()

	// Address Transaction Internal Count
	go addressTransactionInternalCountRoutine()

	// Address Log Count
	go addressLogCountRoutine()

	// Address Token Transfer Count
	go addressTokenTransferCountRoutine()

	// Address Token Balance
	go tokenAddressBalanceRoutine()

	// Address Token Count
	go tokenAddressCountRoutine()

	// Transaction Create Score
	go transactionCreateScoreRoutine()

	// Redis Store Count Keys
	go redisStoreKeysRoutine()
}
