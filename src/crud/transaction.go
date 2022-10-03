package crud

import (
	"sync"
	"time"

	"github.com/sudoblockio/icon-transformer/models"
)

var transactionCrudOnce sync.Once
var TransactionCrud *Crud[models.Transaction, models.TransactionORM]
var TransactionTypeCrud *Crud[models.Transaction, models.TransactionORM]

// GetTransactionCrud - create and/or return the transactions table model
func GetTransactionCrud() *Crud[models.Transaction, models.TransactionORM] {
	transactionCrudOnce.Do(func() {
		TransactionCrud = GetCrud(models.Transaction{}, models.TransactionORM{})

		TransactionCrud.Migrate()
		TransactionCrud.CreateIndexes("CREATE INDEX IF NOT EXISTS transaction_idx_block_number_type_hash ON public.transactions USING btree (block_number, type, hash)")
		TransactionCrud.dbBufferWait = 10 * time.Millisecond

		TransactionCrud.columns = removeColumnNames(TransactionCrud.columns, []string{"transaction_type"})
		TransactionCrud.MakeStartLoaderChannel()

		// Used in transactionCreateScores transformer
		TransactionTypeCrud = GetCrud(models.Transaction{}, models.TransactionORM{})
		TransactionTypeCrud.columns = []string{"hash", "transaction_type", "score_address", "log_index"}
		TransactionTypeCrud.metrics.Name = TransactionCrud.TableName + "_types"
		TransactionTypeCrud.MakeStartLoaderChannel()
	})

	return TransactionCrud
}

func InitTransactionCrud() {
	GetTransactionCrud()
}
