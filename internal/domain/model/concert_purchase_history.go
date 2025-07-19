package model

import (
	"github.com/google/uuid"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/utils"
)

func (m *ConcertPurchaseHistory) TableName() string {
	return "concert_purchase_history"
}

type ConcertPurchaseHistory struct {
	BaseModel

	Concert *Concert `gorm:"foreignKey:ConcertID;references:ID"`

	ConcertID   string `gorm:"column:concert_id;type:uuid;size:36;"`
	UserPhone   string `gorm:"column:user_phone;size:100;"`
	ConcertName string `gorm:"column:concert_name;size:255;"`
	Price       int64  `gorm:"column:price"`
	Qty         int64  `gorm:"column:qty"`
	TotalPrice  int64  `gorm:"column:total_price"`
}

func (m *ConcertPurchaseHistory) ToEntity() *entity.ConcertPurchaseHistory {
	return &entity.ConcertPurchaseHistory{
		ID:          m.ID,
		ConcertID:   m.ConcertID,
		UserPhone:   m.UserPhone,
		ConcertName: m.ConcertName,
		Price:       m.Price,
		Qty:         m.Qty,
		TotalPrice:  m.TotalPrice,
		CreatedAt:   utils.ParseTime(m.CreatedAt.Time),
	}
}

func (m *ConcertPurchaseHistory) CreateBookingConcert(req *entity.BookingConcertRequest, concert *Concert) {
	m.ID = uuid.New().String()
	m.ConcertID = req.ConcertID
	m.UserPhone = req.UserPhone
	m.ConcertName = concert.Name
	m.Price = concert.Price
	m.Qty = req.Qty
	m.TotalPrice = concert.Price * req.Qty
}
