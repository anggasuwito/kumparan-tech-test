package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/usecase"
	"kumparan-tech-test/internal/utils"
	"time"
)

type ConcertHandler struct {
	concertUC usecase.ConcertUC
}

func NewConcertHandler(
	concertUC usecase.ConcertUC,
) *ConcertHandler {
	return &ConcertHandler{
		concertUC: concertUC,
	}
}

func (h *ConcertHandler) SetupHandlers(r *gin.Engine) {
	concertPathV1 := r.Group("/v1")
	concertPathV1.POST("/concert/list", h.getConcertList)
	concertPathV1.POST("/concert/booking", h.bookingConcert)
	concertPathV1.GET("/concert/purchase/history/:concert_id", h.getConcertPurchaseHistory)
}

func (h *ConcertHandler) getConcertList(c *gin.Context) {
	var req entity.GetConcertListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "ConcertHandler.getConcertList.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.concertUC.GetConcertList(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *ConcertHandler) bookingConcert(c *gin.Context) {
	var req entity.BookingConcertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "ConcertHandler.bookingConcert.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.concertUC.BookingConcert(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}

func (h *ConcertHandler) getConcertPurchaseHistory(c *gin.Context) {
	var req entity.GetConcertPurchaseHistoryListRequest

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	req.Search = append(req.Search, &entity.Filter{Field: "concert_id", Value: c.Param("concert_id")})
	resp, err := h.concertUC.GetConcertPurchaseHistoryList(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
