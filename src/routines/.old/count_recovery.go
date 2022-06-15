package _old

//
//import (
//	"errors"
//	"go.uber.org/zap"
//	"gorm.io/gorm"
//	"time"
//
//	"github.com/sudoblockio/icon-transformer/config"
//	"github.com/sudoblockio/icon-transformer/crud"
//	"github.com/sudoblockio/icon-transformer/service"
//	"github.com/sudoblockio/icon-transformer/utils"
//)
//
//// TODO: Old / rm
//func addressCountRecovery() {
//	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
//		go getAddressCounts(i)
//	}
//}
//
//func addressRecovery() {
//	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
//		go getAddressCounts(i)
//	}
//}
//
//func getAddressCounts(worker_id int) {
//	skip := worker_id * config.Config.RoutinesBatchSize
//	limit := config.Config.RoutinesBatchSize
//	for {
//		addresses, err := crud.GetAddressCrud().SelectMany(limit, skip)
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			// Sleep
//			break
//		} else if err != nil {
//			zap.S().Fatal(err.Error())
//		}
//		if len(*addresses) == 0 {
//			// Sleep
//			break
//		}
//
//		zap.S().Info("Routine=AddressCounts", " - Processing ", skip, " addresses...")
//		for _, address := range *addresses {
//			balance, err := service.IconNodeServiceGetBalance(address.Address)
//			if err != nil {
//				// Icon node error
//				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
//				continue
//			}
//
//			// Hex -> float64
//			address.Balance = utils.StringHexToFloat64(balance, 18)
//
//			////////////////////
//			// Staked Balance //
//			////////////////////
//			stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
//			if err != nil {
//				// Icon node error
//				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
//				continue
//			}
//
//			// Hex -> float64
//			address.Balance += utils.StringHexToFloat64(stakedBalance, 18)
//
//			crud.GetAddressCrud().UpsertOneCols(&address, []string{"address", "balance"})
//		}
//
//		skip += skip + config.Config.RoutinesBatchSize*config.Config.RoutinesNumWorkers
//	}
//	zap.S().Info("Routine=AddressBalance - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
//	time.Sleep(config.Config.RoutinesSleepDuration)
//}
