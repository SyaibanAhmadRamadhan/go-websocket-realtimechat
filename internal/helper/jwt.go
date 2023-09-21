package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	UserID string
	Key    string
	Exp    time.Duration
}

func GenerateJwtHS256(jwtModel *Jwt) (string, error) {
	timeNow := time.Now()
	timeExp := timeNow.Add(jwtModel.Exp).Unix()

	tokenParse := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": timeExp,
		"sub": jwtModel.UserID,
	})

	tokenStr, err := tokenParse.SignedString([]byte(jwtModel.Key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ClaimJwtHS256(tokenStr, key string) (map[string]any, error) {
	tokenParse, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	claim, _ := tokenParse.Claims.(jwt.MapClaims)

	return claim, nil
}
