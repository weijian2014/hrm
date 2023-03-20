package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	signKey = []byte("Copyright © 2023 HRM")
)

func GenerateToken(userName string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效期，结束时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间（比如：时间戳）
		NotBefore: jwt.NewNumericDate(time.Now()),                     // 有效期，开始时间（比如：时间戳）
		Issuer:    "Copyright © 2023 HRM Admin",                       // Token 颁发者的唯一标识
		Subject:   "Copyright © 2023 HRM Token Generate",              // 主题（比如：用户id 或 用户名）
		ID:        userName,                                           // jwt唯一标识（比如：UUID
		Audience:  []string{"HRM user"},                               // JWT的接收者（比如：APP的包名）
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func IsTokenValid(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	}, jwt.WithLeeway(5*time.Second))

	if claims, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
		return claims.ID, nil
	} else {
		return "", err
	}
}
