package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type TransactionCreateScoreORM struct {
	AcceptTransactionHash   string
	CreationTransactionHash string `gorm:"primary_key"`
	RejectTransactionHash   string
}

// TableName overrides the default tablename generated by GORM
func (TransactionCreateScoreORM) TableName() string {
	return "transaction_create_scores"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TransactionCreateScore) ToORM(ctx context.Context) (TransactionCreateScoreORM, error) {
	to := TransactionCreateScoreORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCreateScoreWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CreationTransactionHash = m.CreationTransactionHash
	to.AcceptTransactionHash = m.AcceptTransactionHash
	to.RejectTransactionHash = m.RejectTransactionHash
	if posthook, ok := interface{}(m).(TransactionCreateScoreWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionCreateScoreORM) ToPB(ctx context.Context) (TransactionCreateScore, error) {
	to := TransactionCreateScore{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCreateScoreWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CreationTransactionHash = m.CreationTransactionHash
	to.AcceptTransactionHash = m.AcceptTransactionHash
	to.RejectTransactionHash = m.RejectTransactionHash
	if posthook, ok := interface{}(m).(TransactionCreateScoreWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TransactionCreateScore the arg will be the target, the caller the one being converted from

// TransactionCreateScoreBeforeToORM called before default ToORM code
type TransactionCreateScoreWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionCreateScoreORM) error
}

// TransactionCreateScoreAfterToORM called after default ToORM code
type TransactionCreateScoreWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionCreateScoreORM) error
}

// TransactionCreateScoreBeforeToPB called before default ToPB code
type TransactionCreateScoreWithBeforeToPB interface {
	BeforeToPB(context.Context, *TransactionCreateScore) error
}

// TransactionCreateScoreAfterToPB called after default ToPB code
type TransactionCreateScoreWithAfterToPB interface {
	AfterToPB(context.Context, *TransactionCreateScore) error
}

// DefaultCreateTransactionCreateScore executes a basic gorm create call
func DefaultCreateTransactionCreateScore(ctx context.Context, in *TransactionCreateScore, db *gorm.DB) (*TransactionCreateScore, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionCreateScoreORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTransactionCreateScore(ctx context.Context, in *TransactionCreateScore, db *gorm.DB) (*TransactionCreateScore, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.CreationTransactionHash == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TransactionCreateScoreORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TransactionCreateScoreORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TransactionCreateScoreORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TransactionCreateScoreORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionCreateScore(ctx context.Context, in *TransactionCreateScore, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.CreationTransactionHash == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TransactionCreateScoreORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TransactionCreateScoreORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionCreateScoreSet(ctx context.Context, in []*TransactionCreateScore, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.CreationTransactionHash == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.CreationTransactionHash)
	}
	if hook, ok := (interface{}(&TransactionCreateScoreORM{})).(TransactionCreateScoreORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("creation_transaction_hash in (?)", keys).Delete(&TransactionCreateScoreORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TransactionCreateScoreORM{})).(TransactionCreateScoreORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TransactionCreateScoreORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*TransactionCreateScore, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*TransactionCreateScore, *gorm.DB) error
}

// DefaultStrictUpdateTransactionCreateScore clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTransactionCreateScore(ctx context.Context, in *TransactionCreateScore, db *gorm.DB) (*TransactionCreateScore, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransactionCreateScore")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TransactionCreateScoreORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("creation_transaction_hash=?", ormObj.CreationTransactionHash).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type TransactionCreateScoreORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTransactionCreateScore executes a basic gorm update call with patch behavior
func DefaultPatchTransactionCreateScore(ctx context.Context, in *TransactionCreateScore, updateMask *field_mask.FieldMask, db *gorm.DB) (*TransactionCreateScore, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj TransactionCreateScore
	var err error
	if hook, ok := interface{}(&pbObj).(TransactionCreateScoreWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(TransactionCreateScoreWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTransactionCreateScore(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TransactionCreateScoreWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTransactionCreateScore(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TransactionCreateScoreWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TransactionCreateScoreWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *TransactionCreateScore, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *TransactionCreateScore, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *TransactionCreateScore, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *TransactionCreateScore, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTransactionCreateScore executes a bulk gorm update call with patch behavior
func DefaultPatchSetTransactionCreateScore(ctx context.Context, objects []*TransactionCreateScore, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*TransactionCreateScore, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*TransactionCreateScore, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTransactionCreateScore(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTransactionCreateScore patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionCreateScore(ctx context.Context, patchee *TransactionCreateScore, patcher *TransactionCreateScore, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*TransactionCreateScore, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"CreationTransactionHash" {
			patchee.CreationTransactionHash = patcher.CreationTransactionHash
			continue
		}
		if f == prefix+"AcceptTransactionHash" {
			patchee.AcceptTransactionHash = patcher.AcceptTransactionHash
			continue
		}
		if f == prefix+"RejectTransactionHash" {
			patchee.RejectTransactionHash = patcher.RejectTransactionHash
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransactionCreateScore executes a gorm list call
func DefaultListTransactionCreateScore(ctx context.Context, db *gorm.DB) ([]*TransactionCreateScore, error) {
	in := TransactionCreateScore{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TransactionCreateScoreORM{}, &TransactionCreateScore{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("creation_transaction_hash")
	ormResponse := []TransactionCreateScoreORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCreateScoreORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TransactionCreateScore{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionCreateScoreORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionCreateScoreORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TransactionCreateScoreORM) error
}
