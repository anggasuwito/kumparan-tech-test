package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
	"kumparan-tech-test/internal/utils"
	"strings"
)

type ConcertRepo interface {
	GetConcertList(ctx context.Context, filter []*entity.Filter, sort []*entity.Filter, page, limit int64) ([]*model.Concert, int64, error)
	GetConcertPurchaseHistoryList(ctx context.Context, filter []*entity.Filter, sort []*entity.Filter, page, limit int64) ([]*model.ConcertPurchaseHistory, int64, error)
	GetAndLockConcert(ctx context.Context, concertID string) (*model.Concert, error)
	CreateConcertPurchaseHistory(ctx context.Context, req *model.ConcertPurchaseHistory) error
	UpdateConcertQuota(ctx context.Context, concertID string, quota int64) error
}

type concertRepo struct {
	masterDB *gorm.DB
}

func NewConcertRepo(masterDB *gorm.DB) ConcertRepo {
	return &concertRepo{
		masterDB: masterDB,
	}
}

func (r *concertRepo) useTX(ctx context.Context) *gorm.DB {
	if tx := utils.GetTransactionFromContext(ctx); tx != nil {
		return tx
	}
	return r.masterDB
}

func (r *concertRepo) GetConcertList(ctx context.Context, filter []*entity.Filter, sort []*entity.Filter, page, limit int64) ([]*model.Concert, int64, error) {
	var (
		res   = []*model.Concert{}
		count int64
		err   error
	)

	q := r.masterDB.
		Debug().
		Model(&model.Concert{}).
		Where("deleted_at is NULL")
	for _, v := range filter {
		if v.Value == "" {
			continue
		}

		valSlice := strings.Split(v.Value, "|")

		switch v.Field {
		case "id":
			if len(valSlice) > 1 {
				q.Where(fmt.Sprintf("%s IN(?)", v.Field), valSlice)
			} else {
				q.Where(fmt.Sprintf("%s = ?", v.Field), v.Value)
			}
		}
	}

	for _, s := range sort {
		if s.Field != "" {
			q.Order(fmt.Sprintf("%s %s", s.Field, s.Value))
		}
	}

	if len(sort) < 1 {
		q.Order("created_at DESC")
	}

	q.Count(&count)

	err = q.Limit(int(limit)).Offset(int((page - 1) * limit)).Find(&res).Error
	if err != nil {
		return res, count, utils.ErrInternal("Failed get concert list : "+err.Error(), "concertRepo.GetConcertList")
	}
	return res, count, err
}

func (r *concertRepo) GetConcertPurchaseHistoryList(ctx context.Context, filter []*entity.Filter, sort []*entity.Filter, page, limit int64) ([]*model.ConcertPurchaseHistory, int64, error) {
	var (
		res   = []*model.ConcertPurchaseHistory{}
		count int64
		err   error
	)

	q := r.masterDB.
		Debug().
		Model(&model.ConcertPurchaseHistory{}).
		Where("deleted_at is NULL")
	for _, v := range filter {
		if v.Value == "" {
			continue
		}

		valSlice := strings.Split(v.Value, "|")

		switch v.Field {
		case "id", "user_phone", "concert_id":
			if len(valSlice) > 1 {
				q.Where(fmt.Sprintf("%s IN(?)", v.Field), valSlice)
			} else {
				q.Where(fmt.Sprintf("%s = ?", v.Field), v.Value)
			}
		}
	}

	for _, s := range sort {
		if s.Field != "" {
			q.Order(fmt.Sprintf("%s %s", s.Field, s.Value))
		}
	}

	if len(sort) < 1 {
		q.Order("created_at DESC")
	}

	q.Count(&count)

	err = q.Limit(int(limit)).Offset(int((page - 1) * limit)).Find(&res).Error
	if err != nil {
		return res, count, utils.ErrInternal("Failed get concert purchase history list : "+err.Error(), "concertRepo.GetConcertPurchaseHistoryList")
	}
	return res, count, err
}

func (r *concertRepo) GetAndLockConcert(ctx context.Context, concertID string) (*model.Concert, error) {
	var data model.Concert
	err := r.useTX(ctx).
		Debug().
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.Concert{}).
		Where("deleted_at IS NULL").
		Where("id = ?", concertID).
		First(&data).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound("Concert Not Found", "concertRepo.GetAndLockConcert.ErrRecordNotFound")
		}
		return nil, utils.ErrInternal("Failed get concert : "+err.Error(), "concertRepo.GetAndLockConcert")
	}

	return &data, nil
}

func (r *concertRepo) CreateConcertPurchaseHistory(ctx context.Context, req *model.ConcertPurchaseHistory) error {
	err := r.useTX(ctx).
		Debug().
		Create(req).
		Error
	if err != nil {
		return utils.ErrInternal("Failed create concert purchase history : "+err.Error(), "concertRepo.CreateConcertPurchaseHistory")
	}
	return nil
}

func (r *concertRepo) UpdateConcertQuota(ctx context.Context, concertID string, quota int64) error {
	err := r.useTX(ctx).
		Debug().
		Model(&model.Concert{}).
		Where("id = ?", concertID).
		Update("ticket_quota", quota).
		Error
	if err != nil {
		return utils.ErrInternal("Failed update concert quota : "+err.Error(), "concertRepo.UpdateConcertQuota")
	}
	return nil
}
