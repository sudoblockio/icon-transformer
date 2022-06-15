package service

import (
	"errors"
	"fmt"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/redis"
	"github.com/sudoblockio/icon-transformer/utils"
)

func IconNodeServiceGetBlockTransactionHashes(height int) (*[]string, error) {

	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "method": "icx_getBlockByHeight",
    "id": 1,
    "params": {
        "height": "0x%x"
    }
	}`, height)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return nil, err
	}

	// Extract result
	result, ok := body["result"].(map[string]interface{})
	if ok == false {
		return nil, errors.New("Cannot read result")
	}

	// Extract transactions
	transactions, ok := result["confirmed_transaction_list"].([]interface{})
	if ok == false {
		return nil, errors.New("Cannot read confirmed_transaction_list")
	}

	// Extract transaciton hashes
	transactionHashes := []string{}
	for _, t := range transactions {
		tx, ok := t.(map[string]interface{})
		if ok == false {
			return nil, errors.New("1 Cannot read transaction hash from block #" + fmt.Sprintf("%x", height))
		}

		// V1
		hash, ok := tx["tx_hash"].(string)
		if ok == true {
			transactionHashes = append(transactionHashes, "0x"+hash)
			continue
		}

		// V3
		hash, ok = tx["txHash"].(string)
		if ok == true {
			transactionHashes = append(transactionHashes, hash)
			continue
		}

		if ok == false {
			return nil, errors.New("2 Cannot read transaction hash from block #" + fmt.Sprintf("%x", height))
		}
	}

	return &transactionHashes, nil
}

func IconNodeServiceGetTokenDecimalBase(tokenContractAddress string) (int, error) {
	redisCacheKey := config.Config.RedisKeyPrefix + "token_contract_decimals_" + tokenContractAddress
	decimals, err := redis.GetRedisClient().GetCount(redisCacheKey)
	if err != nil {
		return 0, err
	} else if decimals != -1 {
		return int(decimals), nil
	}

	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "%s",
        "dataType": "call",
        "data": {
            "method": "decimals",
            "params": {}
        }
    }
	}`, tokenContractAddress)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return 0, err
	}

	// Extract balance
	decimalsHex, ok := body["result"].(string)
	if ok == false {
		return 0, errors.New("Invalid response")
	}
	decimals = int64(utils.StringHexToFloat64(decimalsHex, 0))

	// Redis cache
	err = redis.GetRedisClient().SetCount(redisCacheKey, decimals)
	if err != nil {
		// Redis error
		return 0, err
	}

	return int(decimals), nil
}

func IconNodeServiceGetTokenContractName(tokenContractAddress string) (string, error) {
	// Redis cache
	redisCacheKey := config.Config.RedisKeyPrefix + "token_contract_name_" + tokenContractAddress
	tokenContractName, err := redis.GetRedisClient().GetValue(redisCacheKey)
	if err != nil {
		return "", err
	} else if tokenContractName != "" {
		return tokenContractName, nil
	}

	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "%s",
        "dataType": "call",
        "data": {
            "method": "name",
            "params": {}
        }
    }
	}`, tokenContractAddress)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return "", err
	}

	// Extract balance
	tokenContractName, ok := body["result"].(string)
	if ok == false {
		return "", errors.New("Invalid response")
	}

	// Redis cache
	err = redis.GetRedisClient().SetValue(redisCacheKey, tokenContractName)
	if err != nil {
		// Redis error
		return "", err
	}

	return tokenContractName, nil
}

func IconNodeServiceGetTokenContractSymbol(tokenContractAddress string) (string, error) {
	// Redis cache
	redisCacheKey := config.Config.RedisKeyPrefix + "token_contract_symbol_" + tokenContractAddress
	tokenContractSymbol, err := redis.GetRedisClient().GetValue(redisCacheKey)
	if err != nil {
		return "", err
	} else if tokenContractSymbol != "" {
		return tokenContractSymbol, nil
	}

	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "%s",
        "dataType": "call",
        "data": {
            "method": "symbol",
            "params": {}
        }
    }
	}`, tokenContractAddress)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return "", err
	}

	// Extract balance
	tokenContractSymbol, ok := body["result"].(string)
	if ok == false {
		return "", errors.New("Invalid response")
	}

	// Redis cache
	err = redis.GetRedisClient().SetValue(redisCacheKey, tokenContractSymbol)
	if err != nil {
		// Redis error
		return "", err
	}

	return tokenContractSymbol, nil
}

func IconNodeServiceGetBalance(publicKey string) (string, error) {
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "method": "icx_getBalance",
    "id": 1234,
    "params": {
        "address": "%s"
    }
	}`, publicKey)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return "", err
	}

	// Extract balance
	balance, ok := body["result"].(string)
	if ok == false {
		return "0x0", errors.New("Invalid response")
	}

	return balance, nil
}

func IconNodeServiceGetStakedBalance(publicKey string) (string, error) {
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "cx0000000000000000000000000000000000000000",
        "dataType": "call",
        "data": {
            "method": "getStake",
            "params": {
                "address": "%s"
            }
        }
    }
	}`, publicKey)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return "", err
	}

	// Extract balance
	resultMap, ok := body["result"].(map[string]interface{})
	if ok == false {
		return "0x0", errors.New("Invalid response")
	}

	balance, ok := resultMap["stake"].(string)
	if ok == false {
		return "0x0", errors.New("Invalid response")
	}

	return balance, nil
}

func IconNodeServiceGetTokenBalance(tokenContractAddress string, tokenHolderAddress string) (string, error) {
	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "%s",
        "dataType": "call",
        "data": {
            "method": "balanceOf",
						"params": {"_owner": "%s"}
        }
    }
	}`, tokenContractAddress, tokenHolderAddress)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return "", err
	}

	// Extract balance
	tokenBalance, ok := body["result"].(string)
	if ok == false {
		return "", errors.New("Invalid response")
	}

	return tokenBalance, nil
}

func IconNodeServiceGetPreps() ([]string, error) {
	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "id": 1234,
    "method": "icx_call",
    "params": {
        "to": "cx0000000000000000000000000000000000000000",
        "dataType": "call",
        "data": {
            "method": "getPReps",
            "params": {
                "startRanking" : "0x1",
                "endRanking": "0xff"
            }
        }
    }
}`)

	body, err := JsonRpcRequestWithRetry(payload)
	pRepNames := []string{}

	if err != nil {
		return nil, err
	}

	var pReps []interface{}
	result, ok := body["result"].(map[string]interface{})
	if ok == false {
		return pRepNames, errors.New("Invalid response")
	}

	pReps, ok = result["preps"].([]interface{})
	if ok == false {
		return pRepNames, errors.New("Invalid response")
	}

	for _, pRepInterface := range pReps {
		pRep, ok := pRepInterface.(map[string]interface{})
		if ok == false {
			return pRepNames, errors.New("Invalid p-rep")
		}

		pRepName, ok := pRep["address"].(string)
		if ok == false {
			return pRepNames, errors.New("Invalid p-rep address")
		}

		pRepNames = append(pRepNames, pRepName)
	}

	return pRepNames, nil
}

func IconNodeServiceGetTransactionResult(hash string) (map[string]interface{}, error) {
	// Request icon contract
	payload := fmt.Sprintf(`{
    "jsonrpc": "2.0",
    "method": "icx_getTransactionResult",
    "id": 1234,
    "params": {
        "txHash": "%s"
    }
	}`, hash)

	body, err := JsonRpcRequestWithRetry(payload)
	if err != nil {
		return nil, err
	}

	result, ok := body["result"].(map[string]interface{})
	if ok == false {
		return nil, errors.New("Invalid response")
	}

	return result, nil
}
