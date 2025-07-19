package utils

import (
	"context"
	"gorm.io/gorm"
)

const (
	transactionContextKey = "tx-db"
)

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, transactionContextKey, tx)
}

func GetTransactionFromContext(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(transactionContextKey).(*gorm.DB)
	return tx
}
