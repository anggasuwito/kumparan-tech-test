package middleware

import (
	"github.com/gin-gonic/gin"
	"kumparan-tech-test/config"
	"kumparan-tech-test/internal/domain/entity"
	"kumparan-tech-test/internal/utils"
	"strings"
)

const (
	dataKey = "data"
)

func TokenChecker(c *gin.Context) {
	cfg := config.GetConfig()
	authorization := c.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 || token[1] == "" {
		utils.ResponseError(c, utils.ErrUnauthorized("You are not authorized", "middleware.TokenChecker"))
		return
	}
	tokenStr := token[1]
	tokenClaims, err := utils.VerifyJWT(tokenStr, cfg.AccessTokenSecret)
	if err != nil {
		utils.ResponseError(c, utils.ErrUnauthorized("Invalid Token", "middleware.TokenChecker.VerifyJWT"))
		return
	}

	c.Set(dataKey, tokenClaims)
	c.Next()
}

func GetUserContextValue(c *gin.Context) *entity.JWTClaimUser {
	if value, exists := c.Get(dataKey); exists {
		if res, ok := value.(*entity.JWTClaimUser); ok {
			return res
		}
	}
	return &entity.JWTClaimUser{}
}

func GetAdminContextValue(c *gin.Context) *entity.JWTClaimAdmin {
	if value, exists := c.Get(dataKey); exists {
		if res, ok := value.(*entity.JWTClaimAdmin); ok {
			return res
		}
	}
	return &entity.JWTClaimAdmin{}
}
