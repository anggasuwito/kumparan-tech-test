package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, err error) {
	if err != nil {
		var e *Error
		if errors.As(err, &e) {
			log.Println(fmt.Sprintf("[%s] - %s", e.Trace, e.Message))
			c.AbortWithStatusJSON(e.Code,
				Response{
					Code:    e.Code,
					Message: e.Message,
					Data:    e.Data,
				},
			)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotImplemented,
		Response{
			Code:    http.StatusNotImplemented,
			Message: "Invalid Error Type",
			Data:    nil,
		},
	)
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK,
		Response{
			Code:    http.StatusOK,
			Message: message,
			Data:    data,
		},
	)
}
