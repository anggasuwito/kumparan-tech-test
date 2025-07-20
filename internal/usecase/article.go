package usecase

import (
	"context"
	"database/sql"
	"errors"
	"kumparan-tech-test/internal/constant"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
	"kumparan-tech-test/internal/repository"
	"kumparan-tech-test/internal/utils"
	"math"
)

type ArticleUC interface {
	CreateArticle(ctx context.Context, req *entity.CreateArticleRequest) (resp *entity.CreateArticleResponse, err error)
	GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp *entity.GetListArticleResponse, err error)
}

type articleUC struct {
	articleRepo repository.ArticleRepo
	authorRepo  repository.AuthorRepo
}

func NewArticleUC(
	articleRepo repository.ArticleRepo,
	authorRepo repository.AuthorRepo,
) ArticleUC {
	return &articleUC{
		articleRepo: articleRepo,
		authorRepo:  authorRepo,
	}
}

func (u *articleUC) CreateArticle(ctx context.Context, req *entity.CreateArticleRequest) (resp *entity.CreateArticleResponse, err error) {
	var (
		article = &model.Article{}
	)

	if err = req.ValidateRequest(); err != nil {
		return nil, err
	}

	//validate author
	author, err := u.authorRepo.GetAuthorByID(ctx, req.AuthorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrBadRequest("author not found", "articleUC.CreateArticle.authorRepo.GetAuthorByID")
		}
		return nil, utils.ErrInternal("failed get author by id : "+err.Error(), "articleUC.CreateArticle.authorRepo.GetAuthorByID")
	}

	article.NewArticle(req)
	err = u.articleRepo.CreateArticle(ctx, article)
	if err != nil {
		return nil, utils.ErrInternal("failed create new article : "+err.Error(), "articleUC.CreateArticle.articleRepo.CreateArticle")
	}

	return &entity.CreateArticleResponse{
		article.ToEntity(author),
	}, nil
}

func (u *articleUC) GetListArticle(ctx context.Context, req *entity.GetListArticleRequest) (resp *entity.GetListArticleResponse, err error) {
	var (
		authorIDs []string
		authorMap = make(map[string]*model.Author)
	)

	if req.Limit > constant.MaxLimit || req.Limit < constant.MinLimit {
		req.Limit = constant.MaxLimit
	}

	if req.Page < constant.MinPage {
		req.Page = constant.MinPage
	}

	data, totalData, err := u.articleRepo.GetListArticle(ctx, req)
	if err != nil {
		return nil, utils.ErrInternal("Failed get article list : "+err.Error(), "articleUC.GetListArticle.articleRepo.GetListArticle")
	}

	//get unique author id
	for _, d := range data {
		if authorMap[d.AuthorID] == nil {
			authorMap[d.AuthorID] = &model.Author{}
			authorIDs = append(authorIDs, d.AuthorID)
		}
	}

	authors, err := u.authorRepo.GetAuthorByIDs(ctx, authorIDs)
	if err != nil {
		return nil, utils.ErrInternal("Failed get author by ids : "+err.Error(), "articleUC.GetListArticle.authorRepo.GetAuthorByIDs")
	}

	//map author data
	for _, author := range authors {
		authorMap[author.ID] = author
	}

	var respData []*entity.Article
	for _, d := range data {
		author := authorMap[d.AuthorID]
		respData = append(respData, d.ToEntity(author))
	}

	return &entity.GetListArticleResponse{
		ListPaginationResponse: &entity.ListPaginationResponse{
			CurrentPage: req.Page,
			TotalPage:   int64(math.Ceil(float64(totalData) / float64(req.Limit))),
			TotalData:   totalData,
			PerPage:     req.Limit,
		},
		Data: respData,
	}, nil
}
