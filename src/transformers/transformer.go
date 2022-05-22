package transformers

func Start() {
	// Blocks Topic
	go startBlocks()

	// Contracts Topic
	go startContracts()

	// Dead Message Topic
	//go startDeadMessages()
}
