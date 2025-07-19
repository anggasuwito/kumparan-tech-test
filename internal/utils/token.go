package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const (
	ClaimsDataKey = "claims_data"
)

type JWTClaim struct {
	ID        string      `json:"id"`
	ExpiredAt int64       `json:"expired_at"`
	Data      interface{} `json:"data"`
}

func GenerateJWT(claimData interface{}, secretKey string, expiredDuration time.Duration) (tokenStr string, data JWTClaim, err error) {
	var (
		tokenB = jwt.New(jwt.SigningMethodHS256)
		claims = tokenB.Claims.(jwt.MapClaims)
	)

	// Set payload
	expiredAt := TimeNow().Add(expiredDuration).Unix()
	data = JWTClaim{
		ID:        uuid.New().String(),
		ExpiredAt: expiredAt,
		Data:      claimData,
	}

	claims["expired_at"] = expiredAt
	claims[ClaimsDataKey] = data

	tokenStr, err = tokenB.SignedString([]byte(secretKey))
	if err != nil {
		return "", data, err
	}
	return tokenStr, data, err
}

func VerifyJWT(token, secretKey string) (jwtClaim JWTClaim, err error) {
	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return jwtClaim, err
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return jwtClaim, err
	}

	jsonClaims, err := json.Marshal(claims[ClaimsDataKey])
	if err != nil {
		return jwtClaim, err
	}

	err = json.Unmarshal(jsonClaims, &jwtClaim)
	if err != nil {
		return jwtClaim, err
	}
	return jwtClaim, nil
}
