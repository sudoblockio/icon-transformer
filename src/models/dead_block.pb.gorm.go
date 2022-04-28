// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dead_block.proto

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

type DeadBlockORM struct {
	Key       string
	Offset    int64  `gorm:"primary_key"`
	Partition int64  `gorm:"primary_key"`
	Topic     string `gorm:"primary_key"`
	Value     string
}

// TableName overrides the default tablename generated by GORM
func (DeadBlockORM) TableName() string {
	return "dead_blocks"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *DeadBlock) ToORM(ctx context.Context) (DeadBlockORM, error) {
	to := DeadBlockORM{}
	var err error
	if prehook, ok := interface{}(m).(DeadBlockWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.Offset = m.Offset
	to.Key = m.Key
	to.Value = m.Value
	if posthook, ok := interface{}(m).(DeadBlockWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *DeadBlockORM) ToPB(ctx context.Context) (DeadBlock, error) {
	to := DeadBlock{}
	var err error
	if prehook, ok := interface{}(m).(DeadBlockWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.Offset = m.Offset
	to.Key = m.Key
	to.Value = m.Value
	if posthook, ok := interface{}(m).(DeadBlockWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type DeadBlock the arg will be the target, the caller the one being converted from

// DeadBlockBeforeToORM called before default ToORM code
type DeadBlockWithBeforeToORM interface {
	BeforeToORM(context.Context, *DeadBlockORM) error
}

// DeadBlockAfterToORM called after default ToORM code
type DeadBlockWithAfterToORM interface {
	AfterToORM(context.Context, *DeadBlockORM) error
}

// DeadBlockBeforeToPB called before default ToPB code
type DeadBlockWithBeforeToPB interface {
	BeforeToPB(context.Context, *DeadBlock) error
}

// DeadBlockAfterToPB called after default ToPB code
type DeadBlockWithAfterToPB interface {
	AfterToPB(context.Context, *DeadBlock) error
}

// DefaultCreateDeadBlock executes a basic gorm create call
func DefaultCreateDeadBlock(ctx context.Context, in *DeadBlock, db *gorm1.DB) (*DeadBlock, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeadBlockORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeadBlockORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type DeadBlockORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type DeadBlockORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskDeadBlock patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskDeadBlock(ctx context.Context, patchee *DeadBlock, patcher *DeadBlock, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*DeadBlock, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Topic" {
			patchee.Topic = patcher.Topic
			continue
		}
		if f == prefix+"Partition" {
			patchee.Partition = patcher.Partition
			continue
		}
		if f == prefix+"Offset" {
			patchee.Offset = patcher.Offset
			continue
		}
		if f == prefix+"Key" {
			patchee.Key = patcher.Key
			continue
		}
		if f == prefix+"Value" {
			patchee.Value = patcher.Value
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListDeadBlock executes a gorm list call
func DefaultListDeadBlock(ctx context.Context, db *gorm1.DB) ([]*DeadBlock, error) {
	in := DeadBlock{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeadBlockORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &DeadBlockORM{}, &DeadBlock{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeadBlockORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("topic")
	ormResponse := []DeadBlockORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeadBlockORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*DeadBlock{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type DeadBlockORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type DeadBlockORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type DeadBlockORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]DeadBlockORM) error
}
