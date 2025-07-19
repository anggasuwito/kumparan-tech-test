package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/usecase"
	"kumparan-tech-test/internal/utils"
	"time"
)

type UserHandler struct {
	concertUC usecase.ConcertUC
}

func NewUserHandler(
	concertUC usecase.ConcertUC,
) *UserHandler {
	return &UserHandler{
		concertUC: concertUC,
	}
}

func (h *UserHandler) SetupHandlers(r *gin.Engine) {
	userPathV1 := r.Group("/v1")
	userPathV1.GET("/user/purchase/history/:user_phone", h.getUserPurchaseHistory)
}

func (h *UserHandler) getUserPurchaseHistory(c *gin.Context) {
	var req entity.GetConcertPurchaseHistoryListRequest

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	req.Search = append(req.Search, &entity.Filter{Field: "user_phone", Value: c.Param("user_phone")})
	resp, err := h.concertUC.GetConcertPurchaseHistoryList(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
