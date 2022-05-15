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
		zap.S().Warn(fmt.Sprintf("Found missing block %d", missingBlockNumber))
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
