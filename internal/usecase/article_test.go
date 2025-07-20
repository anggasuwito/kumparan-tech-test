package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/domain/model"
	mockrepository "kumparan-tech-test/mock/repository"
	"reflect"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		authorID     = uuid.NewString()
		articleTitle = "article-title"
		articleBody  = "article-body"
	)

	articleRepo := mockrepository.NewMockArticleRepo(ctrl)
	authorRepo := mockrepository.NewMockAuthorRepo(ctrl)

	articleUsecase := NewArticleUC(articleRepo, authorRepo)

	type args struct {
		ctx context.Context
		req *entity.CreateArticleRequest
	}
	tests := []struct {
		name     string
		mockFunc func(ctx context.Context)
		args     args
		wantResp *entity.Article
		wantErr  error
	}{
		{
			name: "Success",
			mockFunc: func(ctx context.Context) {
				authorRepo.EXPECT().GetAuthorByID(ctx, authorID).Return(&model.Author{
					BaseModel: model.BaseModel{ID: authorID},
				}, nil)

				articleRepo.EXPECT().CreateArticle(ctx, gomock.Any()).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: &entity.CreateArticleRequest{
					AuthorID: authorID,
					Title:    articleTitle,
					Body:     articleBody,
				},
			},
			wantResp: &entity.Article{
				Author: &entity.Author{
					ID: authorID,
				},
				Title: articleTitle,
				Body:  articleBody,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args.ctx)

			gotResp, err := articleUsecase.CreateArticle(tt.args.ctx, tt.args.req)
			if err != nil || tt.wantErr != nil {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			expectResp := &entity.Article{
				Author: &entity.Author{
					ID: gotResp.Author.ID,
				},
				Title: gotResp.Title,
				Body:  gotResp.Body,
			}
			if !reflect.DeepEqual(expectResp, tt.wantResp) {
				t.Errorf("CreateArticle() gotResp = %v, want %v", expectResp, tt.wantResp)
			}
		})
	}
}
