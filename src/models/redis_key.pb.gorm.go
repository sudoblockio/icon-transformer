package models

import (
	context "context"
	fmt "fmt"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	gorm "gorm.io/gorm"
)

type RedisKeyORM struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

// TableName overrides the default tablename generated by GORM
func (RedisKeyORM) TableName() string {
	return "redis_keys"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *RedisKey) ToORM(ctx context.Context) (RedisKeyORM, error) {
	to := RedisKeyORM{}
	var err error
	if prehook, ok := interface{}(m).(RedisKeyWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Key = m.Key
	to.Value = m.Value
	if posthook, ok := interface{}(m).(RedisKeyWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *RedisKeyORM) ToPB(ctx context.Context) (RedisKey, error) {
	to := RedisKey{}
	var err error
	if prehook, ok := interface{}(m).(RedisKeyWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Key = m.Key
	to.Value = m.Value
	if posthook, ok := interface{}(m).(RedisKeyWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type RedisKey the arg will be the target, the caller the one being converted from

// RedisKeyBeforeToORM called before default ToORM code
type RedisKeyWithBeforeToORM interface {
	BeforeToORM(context.Context, *RedisKeyORM) error
}

// RedisKeyAfterToORM called after default ToORM code
type RedisKeyWithAfterToORM interface {
	AfterToORM(context.Context, *RedisKeyORM) error
}

// RedisKeyBeforeToPB called before default ToPB code
type RedisKeyWithBeforeToPB interface {
	BeforeToPB(context.Context, *RedisKey) error
}

// RedisKeyAfterToPB called after default ToPB code
type RedisKeyWithAfterToPB interface {
	AfterToPB(context.Context, *RedisKey) error
}

// DefaultCreateRedisKey executes a basic gorm create call
func DefaultCreateRedisKey(ctx context.Context, in *RedisKey, db *gorm.DB) (*RedisKey, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type RedisKeyORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadRedisKey(ctx context.Context, in *RedisKey, db *gorm.DB) (*RedisKey, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Key == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := RedisKeyORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(RedisKeyORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type RedisKeyORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteRedisKey(ctx context.Context, in *RedisKey, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Key == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&RedisKeyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type RedisKeyORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteRedisKeySet(ctx context.Context, in []*RedisKey, db *gorm.DB) error {
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
		if ormObj.Key == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Key)
	}
	if hook, ok := (interface{}(&RedisKeyORM{})).(RedisKeyORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("key in (?)", keys).Delete(&RedisKeyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&RedisKeyORM{})).(RedisKeyORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type RedisKeyORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*RedisKey, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*RedisKey, *gorm.DB) error
}

// DefaultStrictUpdateRedisKey clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateRedisKey(ctx context.Context, in *RedisKey, db *gorm.DB) (*RedisKey, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateRedisKey")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &RedisKeyORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("key=?", ormObj.Key).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithAfterStrictUpdateSave); ok {
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

type RedisKeyORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchRedisKey executes a basic gorm update call with patch behavior
func DefaultPatchRedisKey(ctx context.Context, in *RedisKey, updateMask *field_mask.FieldMask, db *gorm.DB) (*RedisKey, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj RedisKey
	var err error
	if hook, ok := interface{}(&pbObj).(RedisKeyWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(RedisKeyWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskRedisKey(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(RedisKeyWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateRedisKey(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(RedisKeyWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type RedisKeyWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *RedisKey, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *RedisKey, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *RedisKey, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *RedisKey, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetRedisKey executes a bulk gorm update call with patch behavior
func DefaultPatchSetRedisKey(ctx context.Context, objects []*RedisKey, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*RedisKey, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*RedisKey, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchRedisKey(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskRedisKey patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskRedisKey(ctx context.Context, patchee *RedisKey, patcher *RedisKey, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*RedisKey, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
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

// DefaultListRedisKey executes a gorm list call
func DefaultListRedisKey(ctx context.Context, db *gorm.DB) ([]*RedisKey, error) {
	in := RedisKey{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("key")
	ormResponse := []RedisKeyORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RedisKeyORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*RedisKey{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type RedisKeyORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RedisKeyORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]RedisKeyORM) error
}
