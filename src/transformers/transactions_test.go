package transformers

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
)

func TestTransactions(t *testing.T) {
	config.ReadEnvironment()

	cases := []TestCase[models.Transaction]{
		{
			"32136447.json",
			&models.Transaction{
				Hash:               "0x970ee35852877013872a049975d0e525d01e0e4de87e6ad22714651e5246ae02",
				LogIndex:           -1,
				Type:               "transaction",
				Method:             "",
				FromAddress:        "",
				ToAddress:          "",
				BlockNumber:        32136447,
				LogCount:           2,
				Version:            "0x3",
				Value:              "",
				ValueDecimal:       0,
				StepLimit:          "",
				Timestamp:          1616479654021402,
				BlockTimestamp:     1616479654021402,
				Nid:                "",
				Nonce:              "",
				TransactionIndex:   0,
				BlockHash:          "ce2d533dd2e0301c49f623444371208abd6de8307893855edafb6e90be8b9b6e",
				TransactionFee:     "0x0",
				Signature:          "",
				DataType:           "base",
				CumulativeStepUsed: "0x0",
				StepUsed:           "0x0",
				StepPrice:          "0x0",
				ScoreAddress:       "hx1000000000000000000000000000000000000000",
				Status:             "0x1",
				TransactionType:    0,
			},
		},
	}

	for _, v := range cases {
		block := testReadBlock(t, v)

		crud.TransactionCrud = &crud.Crud[models.Transaction, models.TransactionORM]{}
		crud.TransactionCrud.LoaderChannel = make(chan *models.Transaction, 2)

		transactions(block)
		channelOutput := <-crud.TransactionCrud.LoaderChannel

		testCompareStructValues(t, v.Expected, channelOutput)
	}
}

func TestTransactionByAddress(t *testing.T) {
	config.ReadEnvironment()

	cases := []TestCase[models.TransactionByAddress]{
		{
			"32136447.json",
			&models.TransactionByAddress{
				TransactionHash: "0x00fe47fbdb47603e7637dad21f952bb6a144c8a21fd0ced93a56e489ad8dca97",
				Address:         "hxa2a3ad042ce6f1d2d41469115b597a7a0e20f11c",
				BlockNumber:     32136447,
			},
		},
	}

	for _, v := range cases {
		block := testReadBlock(t, v)

		crud.TransactionByAddressCrud = crud.GetCrud(models.TransactionByAddress{}, models.TransactionByAddressORM{})
		crud.TransactionByAddressCrud.LoaderChannel = make(chan *models.TransactionByAddress, 2)
		transactionByAddresses(block)

		assert.Equal(t, <-crud.TransactionByAddressCrud.LoaderChannel, v.Expected)
	}
}
