package entity

import (
	"github.com/google/uuid"
	"kumparan-tech-test/internal/utils"
)

type (
	GetConcertListRequest struct {
		ListPaginationRequest
	}

	GetConcertListResponse struct {
		*ListPaginationResponse
		Data []*Concert `json:"data"`
	}

	BookingConcertRequest struct {
		UserPhone string `json:"user_phone"`
		ConcertID string `json:"concert_id"`
		Qty       int64  `json:"qty"`
	}

	BookingConcertResponse struct {
		*ConcertPurchaseHistory
	}

	GetConcertPurchaseHistoryListRequest struct {
		ListPaginationRequest
	}

	GetConcertPurchaseHistoryListResponse struct {
		*ListPaginationResponse
		Data []*ConcertPurchaseHistory `json:"data"`
	}

	Concert struct {
		ID            string `json:"id"`
		AvailableFrom string `json:"available_from"`
		AvailableTo   string `json:"available_to"`
		TicketQuota   int64  `json:"ticket_quota"`
		PlayAt        string `json:"play_at"`
		Name          string `json:"name"`
		Price         int64  `json:"price"`
		CreatedAt     string `json:"created_at"`
	}

	ConcertPurchaseHistory struct {
		ID          string `json:"id"`
		ConcertID   string `json:"concert_id"`
		UserPhone   string `json:"user_phone"`
		ConcertName string `json:"concert_name"`
		Price       int64  `json:"price"`
		Qty         int64  `json:"qty"`
		TotalPrice  int64  `json:"total_price"`
		CreatedAt   string `json:"created_at"`
	}
)

func (e *BookingConcertRequest) ValidateRequest() error {
	if _, err := uuid.Parse(e.ConcertID); err != nil {
		return utils.ErrBadRequest("Invalid concert id", "BookingConcertRequest.Validate.ConcertID")
	}

	if e.UserPhone == "" {
		return utils.ErrBadRequest("Please fill user phone", "BookingConcertRequest.Validate.UserPhone")
	}

	if e.Qty < 1 {
		return utils.ErrBadRequest("Please input quantity", "BookingConcertRequest.Validate.Qty")
	}

	return nil
}
