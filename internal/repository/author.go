package repository

import (
	"context"
	"database/sql"
	"kumparan-tech-test/internal/domain/model"
)

type AuthorRepo interface {
	GetAuthorByIDs(ctx context.Context, req []string) (resp []*model.Author, err error)
	GetAuthorByID(ctx context.Context, req string) (resp *model.Author, err error)
}

type authorRepo struct {
	masterDB *sql.DB
}

func NewAuthorRepo(masterDB *sql.DB) AuthorRepo {
	return &authorRepo{
		masterDB: masterDB,
	}
}

func (r *authorRepo) GetAuthorByIDs(ctx context.Context, req []string) (resp []*model.Author, err error) {
	return nil, nil
}

func (r *authorRepo) GetAuthorByID(ctx context.Context, req string) (resp *model.Author, err error) {
	return nil, nil
}
