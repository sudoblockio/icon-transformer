package crud

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-go-worker/models"
)

// KafkaJobCrud - type for kafkaJob table model
type KafkaJobCrud struct {
	db            *gorm.DB
	model         *models.KafkaJob
	modelORM      *models.KafkaJobORM
	LoaderChannel chan *models.KafkaJob
}

var kafkaJobCrud *KafkaJobCrud
var kafkaJobCrudOnce sync.Once

// GetKafkaJobCrud - create and/or return the kafkaJobs table model
func GetKafkaJobCrud() *KafkaJobCrud {
	kafkaJobCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		kafkaJobCrud = &KafkaJobCrud{
			db:            dbConn,
			model:         &models.KafkaJob{},
			modelORM:      &models.KafkaJobORM{},
			LoaderChannel: make(chan *models.KafkaJob, 1),
		}

		err := kafkaJobCrud.Migrate()
		if err != nil {
			zap.S().Fatal("KafkaJobCrud: Unable migrate postgres table: ", err.Error())
		}
	})

	return kafkaJobCrud
}

// Migrate - migrate kafkaJobs table
func (m *KafkaJobCrud) Migrate() error {
	// Only using KafkaJobRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

// SelectMany - select from kafkaJobs table
func (m *KafkaJobCrud) SelectMany(
	jobID string,
	workerGroup string,
) (*[]models.KafkaJob, error) {
	db := m.db

	// Job ID
	db = db.Where("job_id = ?", jobID)

	// Worker Group
	db = db.Where("worker_group = ?", workerGroup)

	kafkaJob := &[]models.KafkaJob{}
	db = db.Find(kafkaJob)

	return kafkaJob, db.Error
}
