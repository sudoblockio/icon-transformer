// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction_internal_by_address.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type TransactionInternalByAddressORM struct {
	Address         string `gorm:"primary_key"`
	BlockNumber     int64  `gorm:"index:transaction_internal_by_address_idx_block_number"`
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
func DefaultCreateTransactionInternalByAddress(ctx context.Context, in *TransactionInternalByAddress, db *gorm1.DB) (*TransactionInternalByAddress, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
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
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalByAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTransactionInternalByAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionInternalByAddress(ctx context.Context, patchee *TransactionInternalByAddress, patcher *TransactionInternalByAddress, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TransactionInternalByAddress, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
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
func DefaultListTransactionInternalByAddress(ctx context.Context, db *gorm1.DB) ([]*TransactionInternalByAddress, error) {
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
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TransactionInternalByAddressORM{}, &TransactionInternalByAddress{}, nil, nil, nil, nil)
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
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalByAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalByAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TransactionInternalByAddressORM) error
}