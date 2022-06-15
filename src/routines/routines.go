package routines

import (
	"errors"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var addressRoutines = []func(a *models.Address){
	setAddressBalances,
	setAddressTxCounts,
}

var tokenAddressRoutines = []func(t *models.TokenAddress){
	setTokenAddressBalances,
}

func StartRecovery() {
	zap.S().Warn("Init recovery...")

	// By address
	AddressGoRoutines(addressRoutines)
	TokenAddressGoRoutines(tokenAddressRoutines)

	// One shot
	addressTypeRoutine()
	countAddressesToRedisRoutine()
}

var cronRoutines = []func(){
	addressIsPrep,
	tokenAddressCountRoutine, // Isn't used - RM?
}

func CronStart() {

	zap.S().Warn("Init cron...")
	// Init - Jobs that run once on startup
	addressTypeRoutine()

	// Short
	RoutinesCron(cronRoutines, config.Config.RoutinesSleepDuration)

	// Long
	AddressRoutinesCron(addressRoutines, 6*time.Hour)
}

// Wrapper for generic routines
func RoutinesCron(routines []func(), sleepDuration time.Duration) {
	for {
		zap.S().Warn("Starting cron...")
		for _, r := range routines {
			r()
		}
		zap.S().Info("Completed routine, sleeping...")
		time.Sleep(sleepDuration)
	}
}

// Wrapper for generic address routines
func AddressRoutinesCron(routines []func(address *models.Address), sleepDuration time.Duration) {
	for {
		AddressGoRoutines(routines)
		zap.S().Info("Completed routine, sleeping...")
		time.Sleep(sleepDuration)
	}
}

// Takes a crud count method, calls it, takes the count and puts it into a redis countKey.
func CrudCountSetRedis(c func() (int64, error), countKey string) error {
	count, err := c()
	if err != nil {
		// Postgres error
		zap.S().Warn(err)
		return err
	}
	err = redis.GetRedisClient().SetCount(countKey, count)
	if err != nil {
		// Redis error
		zap.S().Warn(err)
		return err
	}
	return nil
}

func AddressGoRoutines(routines []func(address *models.Address)) {
	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
		go AddressSetRoutine(routines, i)
	}
}

func AddressSetRoutine(routines []func(address *models.Address), worker_id int) {
	// Loop through all addresses
	skip := worker_id * config.Config.RoutinesBatchSize
	limit := config.Config.RoutinesBatchSize

	// Run loop until addresses have all been iterated over
	for {
		addresses, err := crud.GetAddressCrud().SelectMany(limit, skip)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Sleep
			break
		} else if err != nil {
			zap.S().Fatal(err.Error())
		}
		if len(*addresses) == 0 {
			// Sleep
			break
		}

		zap.S().Info("Routine=AddressBalance", " - Processing ", len(*addresses), " addresses...")
		for _, address := range *addresses {
			for _, r := range routines {
				r(&address)
			}
		}

		skip += skip + config.Config.RoutinesBatchSize*config.Config.RoutinesNumWorkers
	}
}

func TokenAddressGoRoutines(routines []func(address *models.TokenAddress)) {
	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
		go TokenAddressSetRoutine(routines, i)
	}
}

func TokenAddressSetRoutine(routines []func(tokenAddress *models.TokenAddress), worker_id int) {
	// Loop through all addresses
	skip := worker_id * config.Config.RoutinesBatchSize
	limit := config.Config.RoutinesBatchSize

	// Run loop until addresses have all been iterated over
	for {
		tokenAddresses, err := crud.GetTokenAddressCrud().SelectMany(limit, skip)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Sleep
			break
		} else if err != nil {
			zap.S().Fatal(err.Error())
		}
		if len(*tokenAddresses) == 0 {
			// Sleep
			break
		}

		zap.S().Info("Routine=AddressBalance", " - Processing ", len(*tokenAddresses), " addresses...")
		for _, tokenAddress := range *tokenAddresses {
			for _, r := range routines {
				r(&tokenAddress)
			}
		}

		skip += skip + config.Config.RoutinesBatchSize*config.Config.RoutinesNumWorkers
	}
}
