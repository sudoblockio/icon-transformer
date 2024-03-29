package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type TokenTransferORM struct {
	BlockNumber          int64 `gorm:"index:token_transfer_idx_block_number"`
	BlockTimestamp       int64
	FromAddress          string `gorm:"index:token_transfer_idx_from_address"`
	LogIndex             int64  `gorm:"primary_key"`
	NftId                int64
	ToAddress            string `gorm:"index:token_transfer_idx_to_address"`
	TokenContractAddress string `gorm:"index:token_transfer_idx_token_contract_address"`
	TokenContractName    string
	TokenContractSymbol  string
	TransactionFee       string
	TransactionHash      string `gorm:"primary_key"`
	TransactionIndex     int64
	Value                string
	ValueDecimal         float64
}

// TableName overrides the default tablename generated by GORM
func (TokenTransferORM) TableName() string {
	return "token_transfers"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenTransfer) ToORM(ctx context.Context) (TokenTransferORM, error) {
	to := TokenTransferORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.TokenContractAddress = m.TokenContractAddress
	to.FromAddress = m.FromAddress
	to.ToAddress = m.ToAddress
	to.BlockNumber = m.BlockNumber
	to.Value = m.Value
	to.ValueDecimal = m.ValueDecimal
	to.BlockTimestamp = m.BlockTimestamp
	to.TokenContractName = m.TokenContractName
	to.TokenContractSymbol = m.TokenContractSymbol
	to.TransactionFee = m.TransactionFee
	to.NftId = m.NftId
	to.TransactionIndex = m.TransactionIndex
	if posthook, ok := interface{}(m).(TokenTransferWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenTransferORM) ToPB(ctx context.Context) (TokenTransfer, error) {
	to := TokenTransfer{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.TokenContractAddress = m.TokenContractAddress
	to.FromAddress = m.FromAddress
	to.ToAddress = m.ToAddress
	to.BlockNumber = m.BlockNumber
	to.Value = m.Value
	to.ValueDecimal = m.ValueDecimal
	to.BlockTimestamp = m.BlockTimestamp
	to.TokenContractName = m.TokenContractName
	to.TokenContractSymbol = m.TokenContractSymbol
	to.TransactionFee = m.TransactionFee
	to.NftId = m.NftId
	to.TransactionIndex = m.TransactionIndex
	if posthook, ok := interface{}(m).(TokenTransferWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenTransfer the arg will be the target, the caller the one being converted from

// TokenTransferBeforeToORM called before default ToORM code
type TokenTransferWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenTransferORM) error
}

// TokenTransferAfterToORM called after default ToORM code
type TokenTransferWithAfterToORM interface {
	AfterToORM(context.Context, *TokenTransferORM) error
}

// TokenTransferBeforeToPB called before default ToPB code
type TokenTransferWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenTransfer) error
}

// TokenTransferAfterToPB called after default ToPB code
type TokenTransferWithAfterToPB interface {
	AfterToPB(context.Context, *TokenTransfer) error
}

// DefaultCreateTokenTransfer executes a basic gorm create call
func DefaultCreateTokenTransfer(ctx context.Context, in *TokenTransfer, db *gorm.DB) (*TokenTransfer, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTokenTransfer(ctx context.Context, in *TokenTransfer, db *gorm.DB) (*TokenTransfer, error) {
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
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &TokenTransferORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TokenTransferORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TokenTransferORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTokenTransfer(ctx context.Context, in *TokenTransfer, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TokenTransferORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TokenTransferORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTokenTransferSet(ctx context.Context, in []*TokenTransfer, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&TokenTransferORM{})).(TokenTransferORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("transaction_hash in (?)", keys).Delete(&TokenTransferORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TokenTransferORM{})).(TokenTransferORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TokenTransferORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*TokenTransfer, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*TokenTransfer, *gorm.DB) error
}

// DefaultStrictUpdateTokenTransfer clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTokenTransfer(ctx context.Context, in *TokenTransfer, db *gorm.DB) (*TokenTransfer, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTokenTransfer")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TokenTransferORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("transaction_hash=?", ormObj.TransactionHash).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithAfterStrictUpdateSave); ok {
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

type TokenTransferORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTokenTransfer executes a basic gorm update call with patch behavior
func DefaultPatchTokenTransfer(ctx context.Context, in *TokenTransfer, updateMask *field_mask.FieldMask, db *gorm.DB) (*TokenTransfer, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj TokenTransfer
	var err error
	if hook, ok := interface{}(&pbObj).(TokenTransferWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(TokenTransferWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTokenTransfer(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TokenTransferWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTokenTransfer(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TokenTransferWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TokenTransferWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *TokenTransfer, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *TokenTransfer, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *TokenTransfer, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *TokenTransfer, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTokenTransfer executes a bulk gorm update call with patch behavior
func DefaultPatchSetTokenTransfer(ctx context.Context, objects []*TokenTransfer, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*TokenTransfer, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*TokenTransfer, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTokenTransfer(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTokenTransfer patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenTransfer(ctx context.Context, patchee *TokenTransfer, patcher *TokenTransfer, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*TokenTransfer, error) {
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
		if f == prefix+"TokenContractAddress" {
			patchee.TokenContractAddress = patcher.TokenContractAddress
			continue
		}
		if f == prefix+"FromAddress" {
			patchee.FromAddress = patcher.FromAddress
			continue
		}
		if f == prefix+"ToAddress" {
			patchee.ToAddress = patcher.ToAddress
			continue
		}
		if f == prefix+"BlockNumber" {
			patchee.BlockNumber = patcher.BlockNumber
			continue
		}
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
		if f == prefix+"ValueDecimal" {
			patchee.ValueDecimal = patcher.ValueDecimal
			continue
		}
		if f == prefix+"BlockTimestamp" {
			patchee.BlockTimestamp = patcher.BlockTimestamp
			continue
		}
		if f == prefix+"TokenContractName" {
			patchee.TokenContractName = patcher.TokenContractName
			continue
		}
		if f == prefix+"TokenContractSymbol" {
			patchee.TokenContractSymbol = patcher.TokenContractSymbol
			continue
		}
		if f == prefix+"TransactionFee" {
			patchee.TransactionFee = patcher.TransactionFee
			continue
		}
		if f == prefix+"NftId" {
			patchee.NftId = patcher.NftId
			continue
		}
		if f == prefix+"TransactionIndex" {
			patchee.TransactionIndex = patcher.TransactionIndex
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTokenTransfer executes a gorm list call
func DefaultListTokenTransfer(ctx context.Context, db *gorm.DB) ([]*TokenTransfer, error) {
	in := TokenTransfer{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &TokenTransferORM{}, &TokenTransfer{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TokenTransferORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenTransfer{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenTransferORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TokenTransferORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TokenTransferORM) error
}
