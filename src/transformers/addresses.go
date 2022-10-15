package transformers

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
	"time"
)

var allAddresses = make(map[string]bool)

func enrichContractsMeta(address *models.Address) {
	result, err := service.IconNodeServiceGetScoreStatus(address.Address)
	if err != nil {
		zap.S().Warn("Could not get tx result for hash: ", err.Error(), ",Address=", address.Address)
		return
	}

	auditTxHash, ok := result["current"].(map[string]interface{})["auditTxHash"].(string)
	codeHash, ok := result["current"].(map[string]interface{})["codeHash"].(string)
	deployTxHash, ok := result["current"].(map[string]interface{})["deployTxHash"].(string)
	contractType, ok := result["current"].(map[string]interface{})["type"].(string)
	status, ok := result["current"].(map[string]interface{})["status"].(string)
	owner, ok := result["owner"].(string)

	if ok == false {
		zap.S().Info("not able to parse IconNodeServiceGetScoreStatus for address: ", address.Address)
		return
	}

	address.AuditTxHash = auditTxHash
	address.CodeHash = codeHash
	address.DeployTxHash = deployTxHash
	address.ContractType = contractType
	address.Status = status
	address.Owner = owner

	contractName, _ := service.IconNodeServiceCallContractMethod(address.Address, "name")
	address.Name = contractName
	contractSymbol, _ := service.IconNodeServiceCallContractMethod(address.Address, "symbol")
	address.Symbol = contractSymbol
}

// Check if the address has been seen by the transformer instance, enrich its metadata (expensive) and load it into DB
func loadAddressCheckDuplicate(modelAddress *models.Address) {
	if _, ok := allAddresses[modelAddress.Address]; !ok {

		if modelAddress.IsContract || modelAddress.IsToken {
			enrichContractsMeta(modelAddress)
		}

		allAddresses[modelAddress.Address] = true
		crud.AddressCrud.LoaderChannel <- modelAddress

		// Metrics
		metricsBlockTransformer.addressesSeen.Inc()
		return
	}
	metricsBlockTransformer.addressesIgnored.Inc()
}

func appendAddress(addresses []*models.Address, address *models.Address) []*models.Address {
	if config.Config.ProcessCounts {
		return append(addresses, address)
	}
	return nil
}

func addresses(blockETL *models.BlockETL) {

	addresses := []*models.Address{}
	for _, transactionETL := range blockETL.Transactions {

		// From address
		if transactionETL.FromAddress != "" {

			// Is contract
			isContract := transactionETL.FromAddress[:2] == "cx"

			address := &models.Address{
				Address:    transactionETL.FromAddress,
				IsContract: isContract,
			}
			addresses = appendAddress(addresses, address)
			loadAddressCheckDuplicate(address)
		}

		// To address
		if transactionETL.ToAddress != "" {

			// Is contract
			isContract := transactionETL.ToAddress[:2] == "cx"

			address := &models.Address{
				Address:    transactionETL.ToAddress,
				IsContract: isContract,
			}
			addresses = appendAddress(addresses, address)
			loadAddressCheckDuplicate(address)
		}
	}

	// Internal Transactions
	for _, transactionETL := range blockETL.Transactions {
		for _, logETL := range transactionETL.Logs {

			method := extractMethodFromLogETL(logETL)

			// NOTE 'ICXTransfer' is a protected name in Icon
			if method == "ICXTransfer" {

				// From Address
				fromAddress := logETL.Indexed[1]
				if fromAddress != "" {

					// Is contract
					isContract := fromAddress[:2] == "cx"

					address := &models.Address{
						Address:    fromAddress,
						IsContract: isContract,
					}
					addresses = appendAddress(addresses, address)
					loadAddressCheckDuplicate(address)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					// Is contract
					isContract := toAddress[:2] == "cx"

					address := &models.Address{
						Address:    toAddress,
						IsContract: isContract,
					}

					addresses = appendAddress(addresses, address)
					loadAddressCheckDuplicate(address)
				}
			} else if logETL.Indexed[0] == "Transfer(Address,Address,int,bytes)" && len(logETL.Indexed) == 4 {
				// From Address
				fromAddress := logETL.Indexed[1]
				if fromAddress != "" {

					// Is contract
					isContract := fromAddress[:2] == "cx"

					address := &models.Address{
						Address:    fromAddress,
						IsContract: isContract,
					}

					addresses = appendAddress(addresses, address)
					loadAddressCheckDuplicate(address)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					// Is contract
					isContract := toAddress[:2] == "cx"

					address := &models.Address{
						Address:    toAddress,
						IsContract: isContract,
					}

					addresses = appendAddress(addresses, address)
					loadAddressCheckDuplicate(address)
				}
			}
		}
	}
	if config.Config.ProcessCounts {
		addressBalances(blockETL, addresses)
	}
}

// addressBalances - Update address balance. Expensive process normally just run at head.
func addressBalances(blockETL *models.BlockETL, addresses []*models.Address) {
	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	// block is recent enough, calculate balances
	if time.Since(blockTimestamp) <= config.Config.TransformerServiceCallThreshold {
		for _, address := range addresses {

			// Node call
			balance, err := service.IconNodeServiceGetBalance(address.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Hex -> float64
			address.Balance = utils.StringHexToFloat64(balance, 18)

			////////////////////
			// Staked Balance //
			////////////////////
			stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Hex -> float64
			address.Balance += utils.StringHexToFloat64(stakedBalance, 18)

			// Copy struct for pointer conflicts
			//addressCopy := &models.Address{}
			//copier.Copy(addressCopy, &address)

			// Insert to database
			err = crud.GetAddressCrud().UpsertOneColumns(address, []string{"address", "balance"})
			if err != nil {
				zap.S().Fatal("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
			}
		}
	}
}
