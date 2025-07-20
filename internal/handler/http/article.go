package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/usecase"
	"kumparan-tech-test/internal/utils"
	"time"
)

type ArticleHandler struct {
	articleUC usecase.ArticleUC
}

func NewArticleHandler(
	articleUC usecase.ArticleUC,
) *ArticleHandler {
	return &ArticleHandler{
		articleUC: articleUC,
	}
}

func (h *ArticleHandler) SetupHandlers(r *gin.Engine) {
	articlePathV1 := r.Group("/v1")
	articlePathV1.POST("/article/create", h.createArticle)
	articlePathV1.POST("/article/get-list", h.getListArticle)
}

func (h *ArticleHandler) createArticle(c *gin.Context) {
	var req entity.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "ArticleHandler.createArticle.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.articleUC.CreateArticle(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *ArticleHandler) getListArticle(c *gin.Context) {
	var req entity.GetListArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "ArticleHandler.getListArticle.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.articleUC.GetListArticle(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
