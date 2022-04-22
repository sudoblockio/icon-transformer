// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token_address.proto

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

type TokenAddressORM struct {
	Address              string  `gorm:"primary_key"`
	Balance              float64 `gorm:"index:token_address_idx_balance"`
	TokenContractAddress string  `gorm:"primary_key;index:token_address_idx_token_contract_address"`
}

// TableName overrides the default tablename generated by GORM
func (TokenAddressORM) TableName() string {
	return "token_addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenAddress) ToORM(ctx context.Context) (TokenAddressORM, error) {
	to := TokenAddressORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenAddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.TokenContractAddress = m.TokenContractAddress
	to.Balance = m.Balance
	if posthook, ok := interface{}(m).(TokenAddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenAddressORM) ToPB(ctx context.Context) (TokenAddress, error) {
	to := TokenAddress{}
	var err error
	if prehook, ok := interface{}(m).(TokenAddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.TokenContractAddress = m.TokenContractAddress
	to.Balance = m.Balance
	if posthook, ok := interface{}(m).(TokenAddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenAddress the arg will be the target, the caller the one being converted from

// TokenAddressBeforeToORM called before default ToORM code
type TokenAddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenAddressORM) error
}

// TokenAddressAfterToORM called after default ToORM code
type TokenAddressWithAfterToORM interface {
	AfterToORM(context.Context, *TokenAddressORM) error
}

// TokenAddressBeforeToPB called before default ToPB code
type TokenAddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenAddress) error
}

// TokenAddressAfterToPB called after default ToPB code
type TokenAddressWithAfterToPB interface {
	AfterToPB(context.Context, *TokenAddress) error
}

// DefaultCreateTokenAddress executes a basic gorm create call
func DefaultCreateTokenAddress(ctx context.Context, in *TokenAddress, db *gorm1.DB) (*TokenAddress, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenAddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenAddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenAddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTokenAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenAddress(ctx context.Context, patchee *TokenAddress, patcher *TokenAddress, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TokenAddress, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Address" {
			patchee.Address = patcher.Address
			continue
		}
		if f == prefix+"TokenContractAddress" {
			patchee.TokenContractAddress = patcher.TokenContractAddress
			continue
		}
		if f == prefix+"Balance" {
			patchee.Balance = patcher.Balance
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTokenAddress executes a gorm list call
func DefaultListTokenAddress(ctx context.Context, db *gorm1.DB) ([]*TokenAddress, error) {
	in := TokenAddress{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenAddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TokenAddressORM{}, &TokenAddress{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenAddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("token_contract_address")
	ormResponse := []TokenAddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenAddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenAddress{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenAddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TokenAddressORM) error
}
