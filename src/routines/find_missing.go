package routines

import (
	"fmt"
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

func findMissingBlocks() {

	missingBlockNumbers, err := crud.GetBlockIndexCrud().FindMissing()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	// Delete old rows
	err = crud.GetMissingBlockCrud().DeleteAll()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	// Insert new rows
	for _, missingBlockNumber := range missingBlockNumbers {
		crud.GetMissingBlockCrud().LoaderChannel <- &models.MissingBlock{
			Number: missingBlockNumber,
		}
	}

	// Create extractor jobs
	for _, missingBlockNumber := range missingBlockNumbers {
		body := strings.NewReader(fmt.Sprintf(`{
  		"start_block_number": %d,
  		"end_block_number": %d
		}`, missingBlockNumber, missingBlockNumber+1))
		req, err := http.NewRequest("POST", config.Config.FindMissingExtractorAPILocation+"/create-job", body)
		if err != nil {
			zap.S().Fatal("Routine=FindMissingBlocks - Error making job ", err.Error())
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "*/*")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			zap.S().Fatal("Routine=FindMissingBlocks - Error making job ", err.Error())
		}
		defer resp.Body.Close()
	}

	zap.S().Info("Routine=FindMissingBlocks - Completed routine")
}
