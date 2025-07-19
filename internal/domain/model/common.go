package model

import (
	"database/sql"
	"gorm.io/gorm"
	"kumparan-tech-test/internal/utils"
)

type BaseModel struct {
	ID        string       `gorm:"column:id;type:uuid;primaryKey;size:36"`
	CreatedAt sql.NullTime `gorm:"column:created_at;autoCreateTime:milli"`
	CreatedBy string       `gorm:"column:created_by;size:36;"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;autoUpdateTime:milli"`
	UpdatedBy string       `gorm:"column:updated_by;size:36;"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;index;"`
	DeletedBy string       `gorm:"column:deleted_by;size:36;"`
	Note      string       `gorm:"column:note;size:255;"`
}

func (u *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = sql.NullTime{Time: utils.TimeNow(), Valid: true}
	u.UpdatedAt = sql.NullTime{Time: utils.TimeNow(), Valid: true}
	return
}

func (u *BaseModel) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (u *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = sql.NullTime{Time: utils.TimeNow(), Valid: true}
	return
}

func (u *BaseModel) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
