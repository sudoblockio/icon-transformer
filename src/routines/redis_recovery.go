package routines

import (
	"go.uber.org/zap"
)

func StartRedisRecovery() {
	zap.S().Warn("Starting redis recovery...")
	zap.S().Info("Running block count...")
	err := blockCountExec()
	if err != nil {
		zap.S().Panic("Block count: ", err.Error())
	}

	zap.S().Info("Running tx count...")
	err = transactionRegularCountExec()
	if err != nil {
		zap.S().Panic("Transaction count: ", err.Error())
	}

	zap.S().Info("Running token transfer count...")
	err = tokenTransferCountExec()
	if err != nil {
		zap.S().Panic("Token transfer count: ", err.Error())
	}

	zap.S().Info("Running address count...")
	err = byAddressCountExec()
	if err != nil {
		zap.S().Panic("By address counts: ", err.Error())
	}
}

//func StartRedisRecovery() {
//
//	// Missing Keys
//	//go redisRecovery()
//	zap.S().Warn("Starting redis recovery...")
//	// Block count
//	go func() {
//		zap.S().Info("Running block count...")
//		err := blockCountExec()
//		if err != nil {
//			zap.S().Panic("Block count: ", err.Error())
//		}
//	}()
//
//	// Transaction regular count
//	go func() {
//		zap.S().Info("Running tx count...")
//		err := transactionRegularCountExec()
//		if err != nil {
//			zap.S().Panic("Transaction count: ", err.Error())
//		}
//	}()
//
//	// Token transfer count
//	go func() {
//		zap.S().Info("Running token transfer count...")
//		err := tokenTransferCountExec()
//		if err != nil {
//			zap.S().Panic("Token transfer count: ", err.Error())
//		}
//	}()
//
//	// By address transaction
//	go func() {
//		zap.S().Info("Running address count...")
//		err := byAddressCountExec()
//		if err != nil {
//			zap.S().Panic("By address counts: ", err.Error())
//		}
//	}()
//}

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
