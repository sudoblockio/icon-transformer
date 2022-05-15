package routines

import (
	"go.uber.org/zap"
)

func StartRedisRecovery() {

	// Missing Keys
	//go redisRecovery()
	zap.S().Info("Starting redis recovery...")
	// Block count
	go func() {
		err := blockCountExec()
		if err != nil {
			zap.S().Panic("Block count: ", err.Error())
		}
	}()

	// Transaction regular count
	go func() {
		err := transactionRegularCountExec()
		if err != nil {
			zap.S().Panic("Transaction count: ", err.Error())
		}
	}()

	// Token transfer count
	go func() {
		err := tokenTransferCountExec()
		if err != nil {
			zap.S().Panic("Token transfer count: ", err.Error())
		}
	}()

	// By address transaction
	go func() {
		err := byAddressCountExec()
		if err != nil {
			zap.S().Panic("By address counts: ", err.Error())
		}
	}()
}

//func redisRecovery() {
//
//	redisKeys, err := crud.GetRedisKeyCrud().SelectAll()
//	if err != nil {
//		zap.S().Fatal(err.Error())
//	}
//
//	for _, redisKey := range *redisKeys {
//		err = redis.GetRedisClient().SetValue(redisKey.Key, redisKey.Value)
//		if err != nil {
//			zap.S().Fatal(err.Error())
//		}
//	}
//
//	zap.S().Info("Routine=redisRecovery - Recovered redis")
//}
