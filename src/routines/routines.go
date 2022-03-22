package routines

func StartRoutines() {

	// Address Balance
	go addressBalanceRoutine()

	// Address Token Balance
	go tokenAddressBalanceRoutine()

	// Address Token Count
	go tokenAddressCountRoutine()
}
