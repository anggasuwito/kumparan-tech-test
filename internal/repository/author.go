package repository

import (
	"context"
	"database/sql"
	"fmt"
	"kumparan-tech-test/internal/domain/model"
	"strings"
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
	var (
		query = `SELECT 
    			au.id,
    			au.created_at,
    			au.updated_at,
    			au.name    			
    		FROM author au `
		params = []interface{}{}
	)
	//build param base on req length
	queryParam := []string{}
	for i, v := range req {
		queryParam = append(queryParam, fmt.Sprintf("$%v", i+1))
		params = append(params, v)
	}
	if len(queryParam) > 0 {
		query += fmt.Sprintf(`WHERE au.id IN (%v) `, strings.Join(queryParam, ", "))
	}

	//get data
	rows, err := r.masterDB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author model.Author
		err = rows.Scan(
			&author.ID, &author.CreatedAt, &author.UpdatedAt, &author.Name,
		)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &author)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	return resp, nil
}

func (r *authorRepo) GetAuthorByID(ctx context.Context, req string) (resp *model.Author, err error) {
	var (
		author = model.Author{}
		query  = `SELECT 
    			au.id,
    			au.created_at,
    			au.updated_at,
    			au.name    			
    		FROM author au
    		WHERE au.id = $1`
		params = []interface{}{
			req,
		}
	)

	err = r.masterDB.QueryRowContext(ctx, query, params...).
		Scan(&author.ID, &author.CreatedAt, &author.UpdatedAt, &author.Name)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
