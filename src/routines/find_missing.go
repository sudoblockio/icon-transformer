package routines

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
)

func StartFindMissing() {

	// Missing Blocks
	go findMissingBlocks()

}

func getStartBlock() int64 {
	if config.Config.FindMissingStartBlock == 0 {
		//startBlock, err := crud.GetBlockIndexCrud().SelectLowestNumber()
		startBlock, err := crud.SelectLowestNumber(crud.GetBlockIndexCrud())

		if err != nil {
			zap.S().Fatal(err.Error())
		}
		return startBlock
	} else {
		return config.Config.FindMissingStartBlock
	}
}

func getEndBlock() int64 {
	if config.Config.FindMissingEndBlock == 0 {
		//endBlock, err := crud.GetBlockIndexCrud().SelectHighestNumber()
		endBlock, err := crud.SelectHighestNumber(crud.GetBlockIndexCrud())
		if err != nil {
			zap.S().Fatal(err.Error())
		}
		return endBlock
	} else {
		return config.Config.FindMissingEndBlock
	}
}

func findMissingBlocks() {

	zap.S().Info("Starting finding missing...")

	startBlock := getStartBlock()
	endBlock := getEndBlock()

	//missingBlockNumbers, err := crud.GetBlockIndexCrud().FindMissing(startBlock, endBlock)
	missingBlockNumbers, err := crud.FindMissingBlocks(crud.GetBlockIndexCrud(), startBlock, endBlock)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	zap.S().Info("Found ", len(missingBlockNumbers), " missing blocks.")

	if len(missingBlockNumbers) == 0 {
		zap.S().Info("No missing blocks. Exiting.")
		return
	}

	_, tableCheck := crud.GetMissingBlockCrud().SelectMany(1, 1)
	if tableCheck != nil {
		// Delete old rows
		zap.S().Info("Deleting missing blocks rows in DB.")
		err = crud.GetMissingBlockCrud().DeleteMissing("number > 0")
		if err != nil {
			zap.S().Fatal(err.Error())
		}
		zap.S().Info("Deleted rows")
	}

	// Insert new rows
	for _, missingBlockNumber := range missingBlockNumbers {
		zap.S().Info(fmt.Sprintf("Found missing block %d", missingBlockNumber))

		missingBlock := &models.MissingBlock{
			Number: missingBlockNumber,
		}
		crud.MissingBlockCrud.LoaderChannel <- missingBlock
		//err = crud.GetMissingBlockCrud().UpsertOne(&models.MissingBlock{
		//	Number: missingBlockNumber,
		//})
		//if err != nil {
		//	zap.S().Fatal(err.Error())
		//}
	}

	zap.S().Info("Creating new jobs.")
	// Create extractor jobs
	for _, missingBlockNumber := range missingBlockNumbers {
		body := strings.NewReader(fmt.Sprintf(`{
  		"start_block_number": %d,
  		"end_block_number": %d
		}`, missingBlockNumber, missingBlockNumber+1))

		resp, err := http.Post(config.Config.FindMissingExtractorAPILocation+"/create-job", "application/json", body)
		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["json"])
	}

	zap.S().Info("Routine=FindMissingBlocks - Completed routine")
}
