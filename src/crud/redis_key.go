package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TODO: RM

// RedisKeyCrud - type for redisKey table Model
type RedisKeyCrud struct {
	db            *gorm.DB
	model         *models.RedisKey
	modelORM      *models.RedisKeyORM
	LoaderChannel chan *models.RedisKey
}

var redisKeyCrud *RedisKeyCrud
var redisKeyCrudOnce sync.Once

// GetRedisKeyCrud - create and/or return the redisKeys table Model
func GetRedisKeyCrud() *RedisKeyCrud {
	redisKeyCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		redisKeyCrud = &RedisKeyCrud{
			db:            dbConn,
			model:         &models.RedisKey{},
			modelORM:      &models.RedisKeyORM{},
			LoaderChannel: make(chan *models.RedisKey, 1),
		}

		err := redisKeyCrud.Migrate()
		if err != nil {
			zap.S().Fatal("RedisKeyCrud: Unable migrate postgres table: ", err.Error())
		}

		StartRedisKeyLoader()
	})

	return redisKeyCrud
}

// Migrate - migrate redisKeys table
func (m *RedisKeyCrud) Migrate() error {
	// Only using RedisKeyRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *RedisKeyCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *RedisKeyCrud) SelectAll() (*[]models.RedisKey, error) {
	db := m.db

	// Set table
	db = db.Model(&models.RedisKey{})

	redisKeys := &[]models.RedisKey{}
	db = db.Find(redisKeys)

	return redisKeys, db.Error
}

func (m *RedisKeyCrud) UpsertOne(
	redisKey *models.RedisKey,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*redisKey),
		reflect.TypeOf(*redisKey),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(redisKey)

	return db.Error
}

// StartRedisKeyLoader starts loader
func StartRedisKeyLoader() {
	go func() {

		for {
			// Read redisKey
			newRedisKey := <-GetRedisKeyCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetRedisKeyCrud().UpsertOne(newRedisKey)
			zap.S().Debug(
				"Loader=RedisKey",
				" Key=", newRedisKey.Key,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=RedisKey",
					" Key=", newRedisKey.Key,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
