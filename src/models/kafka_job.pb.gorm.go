package models

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type KafkaJobORM struct {
	JobId       string `gorm:"primary_key"`
	Partition   uint64 `gorm:"primary_key"`
	StopOffset  uint64
	Topic       string `gorm:"primary_key"`
	WorkerGroup string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (KafkaJobORM) TableName() string {
	return "kafka_jobs"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *KafkaJob) ToORM(ctx context.Context) (KafkaJobORM, error) {
	to := KafkaJobORM{}
	var err error
	if prehook, ok := interface{}(m).(KafkaJobWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.WorkerGroup = m.WorkerGroup
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.StopOffset = m.StopOffset
	if posthook, ok := interface{}(m).(KafkaJobWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *KafkaJobORM) ToPB(ctx context.Context) (KafkaJob, error) {
	to := KafkaJob{}
	var err error
	if prehook, ok := interface{}(m).(KafkaJobWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.WorkerGroup = m.WorkerGroup
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.StopOffset = m.StopOffset
	if posthook, ok := interface{}(m).(KafkaJobWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type KafkaJob the arg will be the target, the caller the one being converted from

// KafkaJobBeforeToORM called before default ToORM code
type KafkaJobWithBeforeToORM interface {
	BeforeToORM(context.Context, *KafkaJobORM) error
}

// KafkaJobAfterToORM called after default ToORM code
type KafkaJobWithAfterToORM interface {
	AfterToORM(context.Context, *KafkaJobORM) error
}

// KafkaJobBeforeToPB called before default ToPB code
type KafkaJobWithBeforeToPB interface {
	BeforeToPB(context.Context, *KafkaJob) error
}

// KafkaJobAfterToPB called after default ToPB code
type KafkaJobWithAfterToPB interface {
	AfterToPB(context.Context, *KafkaJob) error
}

// DefaultCreateKafkaJob executes a basic gorm create call
func DefaultCreateKafkaJob(ctx context.Context, in *KafkaJob, db *gorm.DB) (*KafkaJob, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type KafkaJobORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadKafkaJob(ctx context.Context, in *KafkaJob, db *gorm.DB) (*KafkaJob, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Partition == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &KafkaJobORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := KafkaJobORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(KafkaJobORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type KafkaJobORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteKafkaJob(ctx context.Context, in *KafkaJob, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.JobId == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&KafkaJobORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type KafkaJobORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteKafkaJobSet(ctx context.Context, in []*KafkaJob, db *gorm.DB) error {
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
		if ormObj.JobId == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.JobId)
	}
	if hook, ok := (interface{}(&KafkaJobORM{})).(KafkaJobORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("job_id in (?)", keys).Delete(&KafkaJobORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&KafkaJobORM{})).(KafkaJobORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type KafkaJobORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*KafkaJob, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*KafkaJob, *gorm.DB) error
}

// DefaultStrictUpdateKafkaJob clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateKafkaJob(ctx context.Context, in *KafkaJob, db *gorm.DB) (*KafkaJob, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateKafkaJob")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &KafkaJobORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("job_id=?", ormObj.JobId).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterStrictUpdateSave); ok {
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

type KafkaJobORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchKafkaJob executes a basic gorm update call with patch behavior
func DefaultPatchKafkaJob(ctx context.Context, in *KafkaJob, updateMask *field_mask.FieldMask, db *gorm.DB) (*KafkaJob, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj KafkaJob
	var err error
	if hook, ok := interface{}(&pbObj).(KafkaJobWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(KafkaJobWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskKafkaJob(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(KafkaJobWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateKafkaJob(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(KafkaJobWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type KafkaJobWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *KafkaJob, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *KafkaJob, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *KafkaJob, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *KafkaJob, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetKafkaJob executes a bulk gorm update call with patch behavior
func DefaultPatchSetKafkaJob(ctx context.Context, objects []*KafkaJob, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*KafkaJob, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*KafkaJob, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchKafkaJob(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskKafkaJob patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskKafkaJob(ctx context.Context, patchee *KafkaJob, patcher *KafkaJob, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*KafkaJob, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"JobId" {
			patchee.JobId = patcher.JobId
			continue
		}
		if f == prefix+"WorkerGroup" {
			patchee.WorkerGroup = patcher.WorkerGroup
			continue
		}
		if f == prefix+"Topic" {
			patchee.Topic = patcher.Topic
			continue
		}
		if f == prefix+"Partition" {
			patchee.Partition = patcher.Partition
			continue
		}
		if f == prefix+"StopOffset" {
			patchee.StopOffset = patcher.StopOffset
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListKafkaJob executes a gorm list call
func DefaultListKafkaJob(ctx context.Context, db *gorm.DB) ([]*KafkaJob, error) {
	in := KafkaJob{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &KafkaJobORM{}, &KafkaJob{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("partition")
	ormResponse := []KafkaJobORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*KafkaJob{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type KafkaJobORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type KafkaJobORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]KafkaJobORM) error
}
