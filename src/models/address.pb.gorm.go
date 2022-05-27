package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type AddressORM struct {
	Address                  string  `gorm:"primary_key"`
	Balance                  float64 `gorm:"index:address_idx_balance"`
	ContractUpdatedBlock     int64
	CreatedTimestamp         int64 `gorm:"index:address_idx_created_timestamp"`
	IsContract               bool  `gorm:"index:address_idx_is_contract"`
	IsPrep                   bool  `gorm:"index:address_idx_is_governance_prep"`
	IsToken                  bool  `gorm:"index:address_idx_is_token"`
	LogCount                 int64 `gorm:"index:address_idx_log_count"`
	Name                     string
	Status                   string
	TokenTransferCount       int64 `gorm:"index:address_idx_token_transfer_count"`
	TransactionCount         int64 `gorm:"index:address_idx_transaction_count"`
	TransactionInternalCount int64 `gorm:"index:address_idx_transaction_internal_count"`
	Type                     string
}

// TableName overrides the default tablename generated by GORM
func (AddressORM) TableName() string {
	return "addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Address) ToORM(ctx context.Context) (AddressORM, error) {
	to := AddressORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.IsContract = m.IsContract
	to.TransactionCount = m.TransactionCount
	to.TransactionInternalCount = m.TransactionInternalCount
	to.LogCount = m.LogCount
	to.TokenTransferCount = m.TokenTransferCount
	to.Balance = m.Balance
	to.Type = m.Type
	to.Name = m.Name
	to.Status = m.Status
	to.CreatedTimestamp = m.CreatedTimestamp
	to.IsToken = m.IsToken
	to.ContractUpdatedBlock = m.ContractUpdatedBlock
	to.IsPrep = m.IsPrep
	if posthook, ok := interface{}(m).(AddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressORM) ToPB(ctx context.Context) (Address, error) {
	to := Address{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.IsContract = m.IsContract
	to.TransactionCount = m.TransactionCount
	to.TransactionInternalCount = m.TransactionInternalCount
	to.LogCount = m.LogCount
	to.TokenTransferCount = m.TokenTransferCount
	to.Balance = m.Balance
	to.Type = m.Type
	to.Name = m.Name
	to.Status = m.Status
	to.CreatedTimestamp = m.CreatedTimestamp
	to.IsToken = m.IsToken
	to.ContractUpdatedBlock = m.ContractUpdatedBlock
	to.IsPrep = m.IsPrep
	if posthook, ok := interface{}(m).(AddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Address the arg will be the target, the caller the one being converted from

// AddressBeforeToORM called before default ToORM code
type AddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressORM) error
}

// AddressAfterToORM called after default ToORM code
type AddressWithAfterToORM interface {
	AfterToORM(context.Context, *AddressORM) error
}

// AddressBeforeToPB called before default ToPB code
type AddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *Address) error
}

// AddressAfterToPB called after default ToPB code
type AddressWithAfterToPB interface {
	AfterToPB(context.Context, *Address) error
}

// DefaultCreateAddress executes a basic gorm create call
func DefaultCreateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Address == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &AddressORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := AddressORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(AddressORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type AddressORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteAddress(ctx context.Context, in *Address, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&AddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type AddressORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteAddressSet(ctx context.Context, in []*Address, db *gorm.DB) error {
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
		if ormObj.Address == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Address)
	}
	if hook, ok := (interface{}(&AddressORM{})).(AddressORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("address in (?)", keys).Delete(&AddressORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&AddressORM{})).(AddressORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type AddressORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Address, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Address, *gorm.DB) error
}

// DefaultStrictUpdateAddress clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &AddressORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("address=?", ormObj.Address).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterStrictUpdateSave); ok {
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

type AddressORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchAddress executes a basic gorm update call with patch behavior
func DefaultPatchAddress(ctx context.Context, in *Address, updateMask *field_mask.FieldMask, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Address
	var err error
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskAddress(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(AddressWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateAddress(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(AddressWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type AddressWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type AddressWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Address, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetAddress executes a bulk gorm update call with patch behavior
func DefaultPatchSetAddress(ctx context.Context, objects []*Address, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Address, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Address, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchAddress(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAddress(ctx context.Context, patchee *Address, patcher *Address, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Address, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Address" {
			patchee.Address = patcher.Address
			continue
		}
		if f == prefix+"IsContract" {
			patchee.IsContract = patcher.IsContract
			continue
		}
		if f == prefix+"TransactionCount" {
			patchee.TransactionCount = patcher.TransactionCount
			continue
		}
		if f == prefix+"TransactionInternalCount" {
			patchee.TransactionInternalCount = patcher.TransactionInternalCount
			continue
		}
		if f == prefix+"LogCount" {
			patchee.LogCount = patcher.LogCount
			continue
		}
		if f == prefix+"TokenTransferCount" {
			patchee.TokenTransferCount = patcher.TokenTransferCount
			continue
		}
		if f == prefix+"Balance" {
			patchee.Balance = patcher.Balance
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Status" {
			patchee.Status = patcher.Status
			continue
		}
		if f == prefix+"CreatedTimestamp" {
			patchee.CreatedTimestamp = patcher.CreatedTimestamp
			continue
		}
		if f == prefix+"IsToken" {
			patchee.IsToken = patcher.IsToken
			continue
		}
		if f == prefix+"ContractUpdatedBlock" {
			patchee.ContractUpdatedBlock = patcher.ContractUpdatedBlock
			continue
		}
		if f == prefix+"IsPrep" {
			patchee.IsPrep = patcher.IsPrep
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAddress executes a gorm list call
func DefaultListAddress(ctx context.Context, db *gorm.DB) ([]*Address, error) {
	in := Address{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &AddressORM{}, &Address{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("address")
	ormResponse := []AddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Address{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]AddressORM) error
}
