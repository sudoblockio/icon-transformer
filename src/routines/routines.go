package routines

func StartRoutines() {

	// Address Balance
	go addressBalanceRoutine()

	// Address Type
	go addressTypeRoutine()

	// Address Count
	go addressCountRoutine()

	// Address Token Balance
	go tokenAddressBalanceRoutine()

	// Address Token Count
	go tokenAddressCountRoutine()

	// Transaction Create Score
	go transactionCreateScoreRoutine()

	// Redis Store Count Keys
	go redisStoreKeysRoutine()
}
