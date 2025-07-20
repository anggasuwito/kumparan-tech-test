package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
	"kumparan-tech-test/internal/utils"
)

type ArticleRepo interface {
	CreateArticle(ctx context.Context, req *model.Article) error
	GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp []*model.Article, total int64, err error)
}

type articleRepo struct {
	masterDB         *sql.DB
	allowedOrder     []string
	allowedDirection []string
}

func NewArticleRepo(masterDB *sql.DB) ArticleRepo {
	return &articleRepo{
		masterDB: masterDB,
		allowedOrder: []string{
			"a.title",
			"au.name",
		},
		allowedDirection: []string{
			"ASC",
			"DESC",
		},
	}
}

func (r *articleRepo) CreateArticle(ctx context.Context, req *model.Article) error {
	var (
		query = `INSERT INTO article 
    			(id, created_at, author_id, title, body)
			  VALUES 
			    (?, ?, ?, ?, ?)`
		params = []interface{}{
			req.ID,
			req.CreatedAt.Time,
			req.AuthorID,
			req.Title,
			req.Body,
		}
	)

	_, err := r.masterDB.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *articleRepo) GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp []*model.Article, total int64, err error) {
	var (
		offset = (req.Page - 1) * req.Limit
		query  = `SELECT 
    			a.id,
    			a.created_at,
    			a.updated_at,
    			a.author_id,
    			a.title,
    			a.body,
    			au.id,
    			au.created_at,
    			au.updated_at,
    			au.name    			
    		FROM article a
    		LEFT JOIN author au ON a.author_id = au.id
    		WHERE a.deleted_at IS NULL `
		queryCount = `SELECT
    			COUNT(*)
    		FROM article a
    		LEFT JOIN author au ON a.author_id = au.id
    		WHERE a.deleted_at IS NULL `
		params = []interface{}{}
	)

	//build WHERE condition
	for _, s := range req.Search {
		if s.Value == "" {
			continue
		}

		switch s.Field {
		case "keyword":
			condition := `AND (a.title ILIKE ? OR a.body ILIKE ?) `
			query += condition
			queryCount += condition
			params = append(params, []string{s.Value, s.Value})
		case "author_name":
			condition := `AND au.name = ? `
			query += condition
			queryCount += condition
			params = append(params, s.Value)
		}
	}

	//build ORDER BY
	query += `ORDER BY `
	if len(req.Sort) < 1 {
		query += `a.created_at DESC `
	} else {
		for _, s := range req.Sort {
			if utils.InArray(s.Field, r.allowedOrder) && utils.InArray(s.Value, r.allowedDirection) {
				query += fmt.Sprintf("%v %v ", s.Field, s.Value)
			}
		}
	}

	//get count data
	err = r.masterDB.QueryRowContext(ctx, queryCount, params...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	//build LIMIT OFFSET
	query += `LIMIT ? OFFSET ? `
	params = append(params, []int64{req.Limit, offset})

	//get data
	rows, err := r.masterDB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var article model.Article
		var author model.Author
		err = rows.Scan(
			&article.ID, &article.CreatedAt, &article.UpdatedAt, &article.AuthorID, &article.Title, &article.Body,
			&author.ID, &author.CreatedAt, &author.UpdatedAt, &author.Name,
		)
		if err != nil {
			return nil, 0, err
		}

		article.Author = &author
		resp = append(resp, &article)
	}

	if err = rows.Close(); err != nil {
		return nil, 0, err
	}

	if rows.Err() != nil {
		return nil, 0, err
	}

	return resp, total, nil
}
