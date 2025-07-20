package model

import (
	"database/sql"
	"github.com/google/uuid"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/utils"
)

type Article struct {
	BaseModel

	Author *Author

	AuthorID string
	Title    string
	Body     string
}

func (m *Article) NewArticle(req *entity.CreateArticleRequest) {
	m.ID = uuid.NewString()
	m.CreatedAt = sql.NullTime{Time: utils.TimeNow(), Valid: true}
	m.AuthorID = req.AuthorID
	m.Title = req.Title
	m.Body = req.Body
}

func (m *Article) ToEntity(author *Author) *entity.Article {
	return &entity.Article{
		ID:        m.ID,
		CreatedAt: utils.ParseTime(m.CreatedAt.Time),
		UpdatedAt: utils.ParseTime(m.UpdatedAt.Time),
		Title:     m.Title,
		Body:      m.Body,
		Author: &entity.Author{
			ID:        author.ID,
			CreatedAt: utils.ParseTime(author.CreatedAt.Time),
			UpdatedAt: utils.ParseTime(author.UpdatedAt.Time),
			Name:      author.Name,
		},
	}
}
