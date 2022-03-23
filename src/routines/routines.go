package routines

func StartRoutines() {

	// Address Balance
	go addressBalanceRoutine()

	// Address Type
	go addressTypeRoutine()

	// Address Token Balance
	go tokenAddressBalanceRoutine()

	// Address Token Count
	go tokenAddressCountRoutine()

	// Transaction Create Score
	go transactionCreateScoreRoutine()
}