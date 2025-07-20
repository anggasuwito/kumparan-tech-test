package repository

import (
	"context"
	"database/sql"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
)

type ArticleRepo interface {
	CreateArticle(ctx context.Context, req *model.Article) error
	GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp []*model.Article, total int64, err error)
}

type articleRepo struct {
	masterDB *sql.DB
}

func NewArticleRepo(masterDB *sql.DB) ArticleRepo {
	return &articleRepo{
		masterDB: masterDB,
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, req *model.Article) error {
	return nil
}

func (r *articleRepo) GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp []*model.Article, total int64, err error) {
	return nil, 0, nil
}
