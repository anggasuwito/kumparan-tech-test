package repository

import (
	"context"
	"gorm.io/gorm"
	"kumparan-tech-test/internal/utils"
)

type TransactionWrapper interface {
	ExecuteTransaction(ctx context.Context, fns func(ctxTX context.Context) error) error
}

type transactionWrapper struct {
	db *gorm.DB
}

func NewTransactionWrapper(db *gorm.DB) TransactionWrapper {
	return &transactionWrapper{db: db}
}

func (tw *transactionWrapper) ExecuteTransaction(ctx context.Context, fns func(ctxTX context.Context) error) error {
	tx := tw.db.Begin()
	ctxTX := utils.CtxWithTransaction(ctx, tx)

	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fns(ctxTX); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
