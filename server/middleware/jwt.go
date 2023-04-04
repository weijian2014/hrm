package middleware

import (
	"errors"
	"fmt"
	"hrm/common"
	"hrm/log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	signKey    = []byte("Copyright © 2023 HRM")
	bearerName = "Bearer"
)

type TokenInfo struct {
	UserName     string    `json:"user_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type ContextInfo struct {
	UserId   uint64
	UserName string
}

type JwtClaims struct {
	UserId    uint64
	UserName  string
	ExpiredAt time.Time
	jwt.RegisteredClaims
}

func GenerateToken(userId uint64, userName string, expiredAtSeconds uint64) (*TokenInfo, error) {
	accessTokenExpiredAt := time.Now().Add(time.Duration(expiredAtSeconds) * time.Second)
	accessClaims := JwtClaims{
		userId,
		userName,
		accessTokenExpiredAt,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt), // 有效期，结束时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),           // 签发时间（比如：时间戳）
			NotBefore: jwt.NewNumericDate(time.Now()),           // 有效期，开始时间（比如：时间戳）
			Issuer:    "Copyright © 2023 HRM Admin",             // Token 颁发者的唯一标识
			Subject:   "Copyright © 2023 HRM Token Generate",    // 主题（比如：用户id 或 用户名）
			ID:        userName,                                 // jwt唯一标识（比如：UUID
			Audience:  []string{"HRM users"},                    // JWT的接收者（比如：APP的包名）
		},
	}

	refreshTokenExpiredAt := accessTokenExpiredAt.Add(time.Duration(expiredAtSeconds) * time.Second)
	refreshClaims := JwtClaims{
		userId,
		userName,
		refreshTokenExpiredAt,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiredAt), // 有效期，结束时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),            // 签发时间（比如：时间戳）
			NotBefore: jwt.NewNumericDate(time.Now()),            // 有效期，开始时间（比如：时间戳）
			Issuer:    "Copyright © 2023 HRM Admin",              // Token 颁发者的唯一标识
			Subject:   "Copyright © 2023 HRM Token Generate",     // 主题（比如：用户id 或 用户名）
			ID:        userName,                                  // jwt唯一标识（比如：UUID
			Audience:  []string{"HRM users"},                     // JWT的接收者（比如：APP的包名）
		},
	}

	unsignAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := unsignAccessToken.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	unsignRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := unsignRefreshToken.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	tokenInfo := &TokenInfo{
		UserName:     userName,
		AccessToken:  bearerName + " " + accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    accessTokenExpiredAt,
	}
	tokenInfo.AccessToken = bearerName + " " + accessToken
	tokenInfo.RefreshToken = refreshToken
	return tokenInfo, nil
}

func isTokenValid(token string) (*JwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	}, jwt.WithLeeway(5*time.Second))

	if claims, ok := t.Claims.(*JwtClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func AccessTokenAuthenticator(c *gin.Context) {
	// 从请求头中取出Access Token
	signToken := c.Request.Header.Get("Authorization")
	if signToken == "" {
		log.Warn("请求未认证, 请在HTTP请求头中携带Key为Authorization的token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "请求未认证, 请在HTTP请求头中携带Key为Authorization的token",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("signToken=[%v]", signToken)
	// token="Bearer xxxxxxxxxxxxxxxxxxxxx"
	ht := strings.Split(signToken, " ")
	if len(ht) != 2 || ht[0] != bearerName {
		log.Warn("非法token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "非法token",
			"data":    "",
		})
		c.Abort()
		return
	}

	// 校验token, 包括token是否过期
	claims, err := isTokenValid(ht[1])
	if err != nil {
		log.Warn("非法token, err[%+v]", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": fmt.Sprintf("非法token, err: %+v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("claims[%+v]", claims)

	contextInfo := ContextInfo{
		UserId:   claims.UserId,
		UserName: claims.UserName,
	}
	// 后续的处理函数可以用过c.Get("UserInfo")来获取当前请求的UserInfo
	c.Set("ContextInfo", contextInfo)
	c.Next()
}

func RefreshToken(refreshToken string) (*TokenInfo, error) {
	log.Debug("refreshToken [%v]", refreshToken)

	// 校验token
	claims, err := isTokenValid(refreshToken)
	if err != nil {
		log.Warn("IsRefreshTokenValid 非法token, err[%+v]", err)
		return nil, err
	}

	log.Debug("claims[%+v]", claims)

	// token是否过期
	if claims.ExpiredAt.Before(time.Now()) {
		log.Warn("token过期, ExpiredAt[%+v], Now[%+v]", claims.ExpiredAt.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
		return nil, errors.New("refresh token过期")
	}

	newTokenInfo, err := GenerateToken(claims.UserId, claims.UserName, common.JsonConfigs.TokenExpiredSeconds)
	if err != nil {
		log.Warn("重新生成token失败, err[%+v]", err)
		return nil, err
	}

	return newTokenInfo, nil
}
