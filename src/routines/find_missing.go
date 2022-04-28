package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
)

func StartFindMissing() {

	// Missing Blocks
	go findMissingBlocks()

}

func findMissingBlocks() {

	missingBlockNumbers, err := crud.GetBlockCrud().FindMissing()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	for _, missingBlockNumber := range missingBlockNumbers {
		crud.GetMissingBlockCrud().LoaderChannel <- &models.MissingBlock{
			Number: missingBlockNumber,
		}
	}
}
