package usecase

import (
	"context"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
	"kumparan-tech-test/internal/repository"
	"kumparan-tech-test/internal/utils"
	"math"
)

type ConcertUC interface {
	GetConcertList(ctx context.Context, req *entity.GetConcertListRequest) (*entity.GetConcertListResponse, error)
	BookingConcert(ctx context.Context, req *entity.BookingConcertRequest) (*entity.BookingConcertResponse, error)
	GetConcertPurchaseHistoryList(ctx context.Context, req *entity.GetConcertPurchaseHistoryListRequest) (*entity.GetConcertPurchaseHistoryListResponse, error)
}

type concertUC struct {
	txWrapper   repository.TransactionWrapper
	concertRepo repository.ConcertRepo
}

func NewConcertUC(
	txWrapper repository.TransactionWrapper,
	concertRepo repository.ConcertRepo,
) ConcertUC {
	return &concertUC{
		txWrapper:   txWrapper,
		concertRepo: concertRepo,
	}
}

func (u *concertUC) GetConcertList(ctx context.Context, req *entity.GetConcertListRequest) (*entity.GetConcertListResponse, error) {
	var (
		limit = req.Limit
		page  = req.Page
	)

	if limit > 100 || limit == 0 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	data, totalData, err := u.concertRepo.GetConcertList(ctx, req.Search, req.Sort, page, limit)
	if err != nil {
		return nil, utils.ErrInternal("Failed get transaction report list : "+err.Error(), "concertUC.GetConcertList.concertRepo.GetConcertList")
	}

	var resp []*entity.Concert
	for _, d := range data {
		resp = append(resp, d.ToEntity())
	}

	return &entity.GetConcertListResponse{
		ListPaginationResponse: &entity.ListPaginationResponse{
			CurrentPage: page,
			TotalPage:   int64(math.Ceil(float64(totalData) / float64(limit))),
			TotalData:   totalData,
			PerPage:     limit,
		},
		Data: resp,
	}, nil
}

func (u *concertUC) BookingConcert(ctx context.Context, req *entity.BookingConcertRequest) (*entity.BookingConcertResponse, error) {
	var (
		now                    = utils.TimeNow()
		concertPurchaseHistory = &model.ConcertPurchaseHistory{}
	)

	if err := req.ValidateRequest(); err != nil {
		return nil, err
	}

	if err := u.txWrapper.ExecuteTransaction(ctx,
		func(ctxTX context.Context) error {
			//get and lock concert
			concert, err := u.concertRepo.GetAndLockConcert(ctxTX, req.ConcertID)
			if err != nil {
				return err
			}

			//validate booking request
			if req.Qty > concert.TicketQuota {
				return utils.ErrBadRequest("Concert has reached max ticket quota", "concertUC.BookingConcert.TicketQuota")
			}

			if now.Before(concert.AvailableFrom.Time) {
				return utils.ErrBadRequest("Concert ticket are not available for puchase at this time", "concertUC.BookingConcert.AvailableFrom")
			}

			if now.After(concert.AvailableTo.Time) {
				return utils.ErrBadRequest("Concert ticket sales have ended", "concertUC.BookingConcert.AvailableTo")
			}

			//update ticket quota
			err = u.concertRepo.UpdateConcertQuota(ctxTX, concert.ID, concert.TicketQuota-req.Qty)
			if err != nil {
				return err
			}

			//create concert purchase history
			concertPurchaseHistory.CreateBookingConcert(req, concert)
			err = u.concertRepo.CreateConcertPurchaseHistory(ctxTX, concertPurchaseHistory)
			if err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return &entity.BookingConcertResponse{
		concertPurchaseHistory.ToEntity(),
	}, nil
}

func (u *concertUC) GetConcertPurchaseHistoryList(ctx context.Context, req *entity.GetConcertPurchaseHistoryListRequest) (*entity.GetConcertPurchaseHistoryListResponse, error) {
	var (
		limit = req.Limit
		page  = req.Page
	)

	if limit > 100 || limit == 0 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	data, totalData, err := u.concertRepo.GetConcertPurchaseHistoryList(ctx, req.Search, req.Sort, page, limit)
	if err != nil {
		return nil, utils.ErrInternal("Failed get transaction report list : "+err.Error(), "concertUC.GetConcertPurchaseHistory.concertRepo.GetConcertPurchaseHistoryList")
	}

	var resp []*entity.ConcertPurchaseHistory
	for _, d := range data {
		resp = append(resp, d.ToEntity())
	}

	return &entity.GetConcertPurchaseHistoryListResponse{
		ListPaginationResponse: &entity.ListPaginationResponse{
			CurrentPage: page,
			TotalPage:   int64(math.Ceil(float64(totalData) / float64(limit))),
			TotalData:   totalData,
			PerPage:     limit,
		},
		Data: resp,
	}, nil
}
