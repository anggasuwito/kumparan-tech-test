package entity

import (
	"github.com/google/uuid"
	"kumparan-tech-test/internal/utils"
)

type (
	CreateArticleRequest struct {
		AuthorID string `json:"author_id"`
		Title    string `json:"title"`
		Body     string `json:"body"`
	}

	CreateArticleResponse struct {
		*Article
	}

	GetListArticleRequest struct {
		ListPaginationRequest
	}

	GetListArticleResponse struct {
		*ListPaginationResponse
		Data []*Article `json:"data"`
	}

	Article struct {
		ID        string  `json:"id"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		Author    *Author `json:"author"`
		Title     string  `json:"title"`
		Body      string  `json:"body"`
	}
)

func (e *CreateArticleRequest) ValidateRequest() error {
	if _, err := uuid.Parse(e.AuthorID); err != nil {
		return utils.ErrBadRequest("Invalid author id", "CreateArticleRequest.Validate.AuthorID")
	}

	if e.Title == "" {
		return utils.ErrBadRequest("Please input title", "CreateArticleRequest.Validate.Title")
	}

	if e.Body == "" {
		return utils.ErrBadRequest("Please input body", "CreateArticleRequest.Validate.Body")
	}

	return nil
}
