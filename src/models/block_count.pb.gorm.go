package models

import (
	context "context"
	fmt "fmt"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	gorm "gorm.io/gorm"
)

type BlockCountORM struct {
	Count int64
	Type  string `gorm:"primaryKey"`
}

// TableName overrides the default tablename generated by GORM
func (BlockCountORM) TableName() string {
	return "block_counts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *BlockCount) ToORM(ctx context.Context) (BlockCountORM, error) {
	to := BlockCountORM{}
	var err error
	if prehook, ok := interface{}(m).(BlockCountWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Count = m.Count
	if posthook, ok := interface{}(m).(BlockCountWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BlockCountORM) ToPB(ctx context.Context) (BlockCount, error) {
	to := BlockCount{}
	var err error
	if prehook, ok := interface{}(m).(BlockCountWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Type = m.Type
	to.Count = m.Count
	if posthook, ok := interface{}(m).(BlockCountWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type BlockCount the arg will be the target, the caller the one being converted from

// BlockCountBeforeToORM called before default ToORM code
type BlockCountWithBeforeToORM interface {
	BeforeToORM(context.Context, *BlockCountORM) error
}

// BlockCountAfterToORM called after default ToORM code
type BlockCountWithAfterToORM interface {
	AfterToORM(context.Context, *BlockCountORM) error
}

// BlockCountBeforeToPB called before default ToPB code
type BlockCountWithBeforeToPB interface {
	BeforeToPB(context.Context, *BlockCount) error
}

// BlockCountAfterToPB called after default ToPB code
type BlockCountWithAfterToPB interface {
	AfterToPB(context.Context, *BlockCount) error
}

// DefaultCreateBlockCount executes a basic gorm create call
func DefaultCreateBlockCount(ctx context.Context, in *BlockCount, db *gorm.DB) (*BlockCount, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BlockCountORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadBlockCount(ctx context.Context, in *BlockCount, db *gorm.DB) (*BlockCount, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Type == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := BlockCountORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(BlockCountORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type BlockCountORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteBlockCount(ctx context.Context, in *BlockCount, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Type == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&BlockCountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type BlockCountORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteBlockCountSet(ctx context.Context, in []*BlockCount, db *gorm.DB) error {
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
		if ormObj.Type == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Type)
	}
	if hook, ok := (interface{}(&BlockCountORM{})).(BlockCountORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("type in (?)", keys).Delete(&BlockCountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&BlockCountORM{})).(BlockCountORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type BlockCountORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*BlockCount, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*BlockCount, *gorm.DB) error
}

// DefaultStrictUpdateBlockCount clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateBlockCount(ctx context.Context, in *BlockCount, db *gorm.DB) (*BlockCount, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateBlockCount")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &BlockCountORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("type=?", ormObj.Type).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithAfterStrictUpdateSave); ok {
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

type BlockCountORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchBlockCount executes a basic gorm update call with patch behavior
func DefaultPatchBlockCount(ctx context.Context, in *BlockCount, updateMask *field_mask.FieldMask, db *gorm.DB) (*BlockCount, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj BlockCount
	var err error
	if hook, ok := interface{}(&pbObj).(BlockCountWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(BlockCountWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskBlockCount(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(BlockCountWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateBlockCount(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(BlockCountWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type BlockCountWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *BlockCount, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type BlockCountWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *BlockCount, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type BlockCountWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *BlockCount, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type BlockCountWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *BlockCount, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetBlockCount executes a bulk gorm update call with patch behavior
func DefaultPatchSetBlockCount(ctx context.Context, objects []*BlockCount, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*BlockCount, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*BlockCount, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchBlockCount(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskBlockCount patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBlockCount(ctx context.Context, patchee *BlockCount, patcher *BlockCount, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*BlockCount, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBlockCount executes a gorm list call
func DefaultListBlockCount(ctx context.Context, db *gorm.DB) ([]*BlockCount, error) {
	in := BlockCount{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("type")
	ormResponse := []BlockCountORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockCountORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*BlockCount{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BlockCountORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type BlockCountORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]BlockCountORM) error
}
