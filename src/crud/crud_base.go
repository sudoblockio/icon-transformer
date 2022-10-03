package crud

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sudoblockio/icon-transformer/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type CrudMetrics struct {
	Name                         string
	loaderChannelLength          prometheus.Gauge
	loaderChannelDuplicateErrors prometheus.Counter
}

// Generic struct to hold common objects when doing crud
type Crud[M any, O any] struct {
	db                  *gorm.DB
	model               *M
	modelORM            *O
	columns             []string
	primaryKeys         []clause.Column
	TableName           string
	metrics             CrudMetrics
	LoaderChannel       chan *M
	loaderChannelBuffer int
	batchErrorHandler   func(error, []*M) error
	retryAttempts       int
	dbBufferWait        time.Duration
	batchSampleInterval int
}

type ModelOrm interface {
	TableName() string
}

// GetCrud - Builds generic crud struct
func GetCrud[M any, O ModelOrm](m M, o O) *Crud[M, O] {
	dbConn := getPostgresConn()
	if dbConn == nil {
		zap.S().Fatal("Cannot connect to postgres database")
	}

	return &Crud[M, O]{
		db:                  dbConn,
		model:               &m,
		modelORM:            &o,
		columns:             getModelColumnNames(m),
		primaryKeys:         getModelPrimaryKeys(o),
		TableName:           o.TableName(),
		dbBufferWait:        config.Config.DbBufferWait,
		loaderChannelBuffer: config.Config.DbLoaderChannelBuffer,
		batchSampleInterval: 100,
		retryAttempts:       5,
	}
}

// Migrate - migrate table
func (m *Crud[M, O]) Migrate() {
	if config.Config.DbSkipMigrations {
		return
	}

	err := m.db.AutoMigrate(m.modelORM)
	if err != nil {
		zap.S().Fatal("TokenTransferCrud: Unable migrate postgres table: ", err.Error())
	}
}

// TODO: Implement wrapper to pull indexes from gorm tags and code gen this statement

// CreateIndexes - Create indexes
func (m *Crud[Model, ModelOrm]) CreateIndexes(statement string) {
	if config.Config.DbSkipMigrations {
		return
	}

	db := m.db
	db = db.Exec(statement)
	if db.Error != nil {
		zap.S().Warn("Unable to create indexes: ", db.Error.Error())
	}
}

func (m *Crud[M, O]) DefaultRetryHandler(err error, values []*M) error {
	if err == nil {
		return nil
	}

	gormErr := getGormError(err)

	switch gormErr.Code {
	case "":
		zap.S().Info("Nil error on table: ", m.TableName)
		return nil
	case "0":
		return nil
	case "21000":
		m.metrics.loaderChannelDuplicateErrors.Inc()
		return m.LoopUpsertOne(values)
	default:
		zap.S().Info("Error on table: ", m.TableName)
		zap.S().Info("Exit code: ", gormErr.Code)
		zap.S().Info("Exit message: ", gormErr.Message)
		zap.S().Info("Exit on values: ", values)
		zap.S().Fatal(gormErr.Message)
		return err
	}
}
