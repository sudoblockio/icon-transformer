package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type TransactionInternalByAddressORM struct {
	Address         string `gorm:"primary_key"`
	BlockNumber     int64
	LogIndex        int64  `gorm:"primary_key"`
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (TransactionInternalByAddressORM) TableName() string {
	return "transaction_internal_by_addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TransactionInternalByAddress) ToORM(ctx context.Context) (TransactionInternalByAddressORM, error) {
	to := TransactionInternalByAddressORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionInternalByAddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TransactionInternalByAddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionInternalByAddressORM) ToPB(ctx context.Context) (TransactionInternalByAddress, error) {
	to := TransactionInternalByAddress{}
	var err error
	if prehook, ok := interface{}(m).(TransactionInternalByAddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TransactionInternalByAddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TransactionInternalByAddress the arg will be the target, the caller the one being converted from

// TransactionInternalByAddressBeforeToORM called before default ToORM code
type TransactionInternalByAddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionInternalByAddressORM) error
}

// TransactionInternalByAddressAfterToORM called after default ToORM code
type TransactionInternalByAddressWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionInternalByAddressORM) error
}

// TransactionInternalByAddressBeforeToPB called before default ToPB code
type TransactionInternalByAddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *TransactionInternalByAddress) error
}

// TransactionInternalByAddressAfterToPB called after default ToPB code
type TransactionInternalByAddressWithAfterToPB interface {
	AfterToPB(context.Context, *TransactionInternalByAddress) error
}

// DefaultCreateTransactionInternalByAddress executes a basic gorm create call
func DefaultCreateTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, db *gorm.DB) (*TransactionInternalByAddress, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionInternalByAddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, db *gorm.DB) (*TransactionInternalByAddress, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.TransactionHash == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TransactionInternalByAddressORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TransactionInternalByAddressORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TransactionInternalByAddressORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TransactionInternalByAddressORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.TransactionHash == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TransactionInternalByAddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TransactionInternalByAddressORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTransactionInternalByAddressSet(ctx context.Context, in []*TransactionInternalByAddress, db *gorm.DB) error {
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
		if ormObj.TransactionHash == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.TransactionHash)
	}
	if hook, ok := (interface{}(&TransactionInternalByAddressORM{})).(TransactionInternalByAddressORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("transaction_hash in (?)", keys).Delete(&TransactionInternalByAddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TransactionInternalByAddressORM{})).(TransactionInternalByAddressORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TransactionInternalByAddressORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*TransactionInternalByAddress, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*TransactionInternalByAddress, *gorm.DB) error
}

// DefaultStrictUpdateTransactionInternalByAddress clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, db *gorm.DB) (*TransactionInternalByAddress, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTransactionInternalByAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TransactionInternalByAddressORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("transaction_hash=?", ormObj.TransactionHash).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithAfterStrictUpdateSave); ok {
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

type TransactionInternalByAddressORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTransactionInternalByAddress executes a basic gorm update call with patch behavior
func DefaultPatchTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, updateMask *field_mask.FieldMask, db *gorm.DB) (*TransactionInternalByAddress, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj TransactionInternalByAddress
	var err error
	if hook, ok := interface{}(&pbObj).(TransactionInternalByAddressWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(TransactionInternalByAddressWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTransactionInternalByAddress(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TransactionInternalByAddressWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTransactionInternalByAddress(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TransactionInternalByAddressWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TransactionInternalByAddressWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *TransactionInternalByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *TransactionInternalByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *TransactionInternalByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *TransactionInternalByAddress, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTransactionInternalByAddress executes a bulk gorm update call with patch behavior
func DefaultPatchSetTransactionInternalByAddress(ctx context.Context, objects []*TransactionInternalByAddress, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*TransactionInternalByAddress, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*TransactionInternalByAddress, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTransactionInternalByAddress(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTransactionInternalByAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionInternalByAddress(ctx context.Context, patchee *TransactionInternalByAddress, patcher *TransactionInternalByAddress, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*TransactionInternalByAddress, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
		if f == prefix+"Address" {
			patchee.Address = patcher.Address
			continue
		}
		if f == prefix+"BlockNumber" {
			patchee.BlockNumber = patcher.BlockNumber
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransactionInternalByAddress executes a gorm list call
func DefaultListTransactionInternalByAddress(ctx context.Context, db *gorm.DB) ([]*TransactionInternalByAddress, error) {
	in := TransactionInternalByAddress{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TransactionInternalByAddressORM{}, &TransactionInternalByAddress{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TransactionInternalByAddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalByAddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TransactionInternalByAddress{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionInternalByAddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TransactionInternalByAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TransactionInternalByAddressORM) error
}
