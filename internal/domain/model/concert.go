package model

import (
	"database/sql"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/utils"
)

func (m *Concert) TableName() string {
	return "concert"
}

type Concert struct {
	BaseModel

	ConcertPurchaseHistory []*ConcertPurchaseHistory `gorm:"foreignKey:ConcertID"`

	AvailableFrom sql.NullTime `gorm:"column:available_from"`
	AvailableTo   sql.NullTime `gorm:"column:available_to"`
	TicketQuota   int64        `gorm:"column:ticket_quota"`
	PlayAt        sql.NullTime `gorm:"column:play_at"`
	Price         int64        `gorm:"column:price"`
	Name          string       `gorm:"column:name;size:255;"`
}

func (m *Concert) ToEntity() *entity.Concert {
	return &entity.Concert{
		ID:            m.ID,
		AvailableFrom: utils.ParseTime(m.AvailableFrom.Time),
		AvailableTo:   utils.ParseTime(m.AvailableTo.Time),
		TicketQuota:   m.TicketQuota,
		PlayAt:        utils.ParseTime(m.PlayAt.Time),
		Name:          m.Name,
		Price:         m.Price,
		CreatedAt:     utils.ParseTime(m.CreatedAt.Time),
	}
}
