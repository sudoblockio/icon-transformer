package crud

//// TransactionCreateScoreCrud - type for transactionCreateScore table Model
//type TransactionCreateScoreCrud struct {
//	db            *gorm.DB
//	Model         *models.TransactionCreateScore
//	ModelORM      *models.TransactionCreateScoreORM
//	LoaderChannel chan *models.TransactionCreateScore
//}
//
//var transactionCreateScoreCrud *TransactionCreateScoreCrud
//var transactionCreateScoreCrudOnce sync.Once
//
//// GetTransactionCreateScoreCrud - create and/or return the transactionCreateScores table Model
//func GetTransactionCreateScoreCrud() *TransactionCreateScoreCrud {
//	transactionCreateScoreCrudOnce.Do(func() {
//		dbConn := getPostgresConn()
//		if dbConn == nil {
//			zap.S().Fatal("Cannot connect to postgres database")
//		}
//
//		transactionCreateScoreCrud = &TransactionCreateScoreCrud{
//			db:            dbConn,
//			Model:         &models.TransactionCreateScore{},
//			ModelORM:      &models.TransactionCreateScoreORM{},
//			LoaderChannel: make(chan *models.TransactionCreateScore, 1),
//		}
//
//		err := transactionCreateScoreCrud.Migrate()
//		if err != nil {
//			zap.S().Fatal("TransactionCreateScoreCrud: Unable migrate postgres table: ", err.Error())
//		}
//
//		StartTransactionCreateScoreLoader()
//	})
//
//	return transactionCreateScoreCrud
//}
//
//// Migrate - migrate transactionCreateScores table
//func (m *TransactionCreateScoreCrud) Migrate() error {
//	// Only using TransactionCreateScoreRawORM (ORM version of the proto generated struct) to create the TABLE
//	err := m.db.AutoMigrate(m.ModelORM) // Migration and Index creation
//	return err
//}
//func (m *TransactionCreateScoreCrud) TableName() string {
//	return m.ModelORM.TableName()
//}
//
//// SelectMany - select many from addreses table
//func (m *TransactionCreateScoreCrud) SelectMany(
//	limit int,
//	skip int,
//) (*[]models.TransactionCreateScore, error) {
//	db := m.db
//
//	// Set table
//	db = db.Model(&models.TransactionCreateScore{})
//
//	// Limit
//	db = db.Limit(limit)
//
//	// Skip
//	if skip != 0 {
//		db = db.Offset(skip)
//	}
//
//	transactionCreateScores := &[]models.TransactionCreateScore{}
//	db = db.Find(transactionCreateScores)
//
//	return transactionCreateScores, db.Error
//}
