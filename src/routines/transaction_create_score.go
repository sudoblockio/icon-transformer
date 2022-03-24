package routines

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func transactionCreateScoreRoutine() {

	// Loop every duration
	for {

		// Loop through all addresses
		skip := 0
		limit := config.Config.RoutinesBatchSize
		for {
			transactionCreateScores, err := crud.GetTransactionCreateScoreCrud().SelectMany(limit, skip)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Sleep
				break
			} else if err != nil {
				zap.S().Fatal(err.Error())
			}
			if len(*transactionCreateScores) == 0 {
				// Sleep
				break
			}

			zap.S().Info("Routine=TransactionCreateScoresRoutine", " - Processing ", len(*transactionCreateScores), " transactionCreateScores...")
			for _, transactionCreateScore := range *transactionCreateScores {
				acceptTransactionHash := transactionCreateScore.AcceptTransactionHash
				rejectTransactionHash := transactionCreateScore.RejectTransactionHash

				createScoreTransaction, err := crud.GetTransactionCrud().SelectOne(transactionCreateScore.CreationTransactionHash, -1)

				// Update accept/reject transactions
				updateScoreTransaction := &models.Transaction{}
				if acceptTransactionHash != "" {
					updateScoreTransaction, err = crud.GetTransactionCrud().SelectOne(acceptTransactionHash, -1)
				} else if rejectTransactionHash != "" {
					updateScoreTransaction, err = crud.GetTransactionCrud().SelectOne(rejectTransactionHash, -1)
				}

				updateScoreTransaction.ToAddress = createScoreTransaction.ToAddress

				if err == nil {
					crud.GetTransactionCrud().UpsertOne(updateScoreTransaction)
				}
			}

			skip += limit
		}
		zap.S().Info("Routine=TransactionCreateScoreRoutine - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
