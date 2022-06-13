package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type TokenTransferByAddressORM struct {
	Address         string `gorm:"primary_key"`
	BlockNumber     int64  `gorm:"index:token_transfer_by_address_idx_block_number"`
	LogIndex        int64  `gorm:"primary_key"`
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (TokenTransferByAddressORM) TableName() string {
	return "token_transfer_by_addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenTransferByAddress) ToORM(ctx context.Context) (TokenTransferByAddressORM, error) {
	to := TokenTransferByAddressORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferByAddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TokenTransferByAddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenTransferByAddressORM) ToPB(ctx context.Context) (TokenTransferByAddress, error) {
	to := TokenTransferByAddress{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferByAddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TokenTransferByAddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenTransferByAddress the arg will be the target, the caller the one being converted from

// TokenTransferByAddressBeforeToORM called before default ToORM code
type TokenTransferByAddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenTransferByAddressORM) error
}

// TokenTransferByAddressAfterToORM called after default ToORM code
type TokenTransferByAddressWithAfterToORM interface {
	AfterToORM(context.Context, *TokenTransferByAddressORM) error
}

// TokenTransferByAddressBeforeToPB called before default ToPB code
type TokenTransferByAddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenTransferByAddress) error
}

// TokenTransferByAddressAfterToPB called after default ToPB code
type TokenTransferByAddressWithAfterToPB interface {
	AfterToPB(context.Context, *TokenTransferByAddress) error
}

// DefaultCreateTokenTransferByAddress executes a basic gorm create call
func DefaultCreateTokenTransferByAddress(ctx context.Context, in *TokenTransferByAddress, db *gorm.DB) (*TokenTransferByAddress, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferByAddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTokenTransferByAddress(ctx context.Context, in *TokenTransferByAddress, db *gorm.DB) (*TokenTransferByAddress, error) {
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
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TokenTransferByAddressORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TokenTransferByAddressORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TokenTransferByAddressORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferByAddressORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTokenTransferByAddress(ctx context.Context, in *TokenTransferByAddress, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Address == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TokenTransferByAddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TokenTransferByAddressORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTokenTransferByAddressSet(ctx context.Context, in []*TokenTransferByAddress, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&TokenTransferByAddressORM{})).(TokenTransferByAddressORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("transaction_hash in (?)", keys).Delete(&TokenTransferByAddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TokenTransferByAddressORM{})).(TokenTransferByAddressORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TokenTransferByAddressORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*TokenTransferByAddress, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*TokenTransferByAddress, *gorm.DB) error
}

// DefaultStrictUpdateTokenTransferByAddress clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTokenTransferByAddress(ctx context.Context, in *TokenTransferByAddress, db *gorm.DB) (*TokenTransferByAddress, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTokenTransferByAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TokenTransferByAddressORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("log_index=?", ormObj.LogIndex).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithAfterStrictUpdateSave); ok {
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

type TokenTransferByAddressORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTokenTransferByAddress executes a basic gorm update call with patch behavior
func DefaultPatchTokenTransferByAddress(ctx context.Context, in *TokenTransferByAddress, updateMask *field_mask.FieldMask, db *gorm.DB) (*TokenTransferByAddress, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj TokenTransferByAddress
	var err error
	if hook, ok := interface{}(&pbObj).(TokenTransferByAddressWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(TokenTransferByAddressWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTokenTransferByAddress(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TokenTransferByAddressWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTokenTransferByAddress(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TokenTransferByAddressWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TokenTransferByAddressWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *TokenTransferByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *TokenTransferByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *TokenTransferByAddress, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *TokenTransferByAddress, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTokenTransferByAddress executes a bulk gorm update call with patch behavior
func DefaultPatchSetTokenTransferByAddress(ctx context.Context, objects []*TokenTransferByAddress, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*TokenTransferByAddress, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*TokenTransferByAddress, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTokenTransferByAddress(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTokenTransferByAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenTransferByAddress(ctx context.Context, patchee *TokenTransferByAddress, patcher *TokenTransferByAddress, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*TokenTransferByAddress, error) {
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

// DefaultListTokenTransferByAddress executes a gorm list call
func DefaultListTokenTransferByAddress(ctx context.Context, db *gorm.DB) ([]*TokenTransferByAddress, error) {
	in := TokenTransferByAddress{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TokenTransferByAddressORM{}, &TokenTransferByAddress{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TokenTransferByAddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferByAddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenTransferByAddress{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenTransferByAddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferByAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TokenTransferByAddressORM) error
}
