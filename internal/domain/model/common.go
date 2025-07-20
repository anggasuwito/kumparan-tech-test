package model

import (
	"database/sql"
)

type BaseModel struct {
	ID        string
	CreatedAt sql.NullTime
	CreatedBy string
	UpdatedAt sql.NullTime
	UpdatedBy string
	DeletedAt sql.NullTime
	DeletedBy string
	Note      string
}
