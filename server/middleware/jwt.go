package middleware

import (
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

type UserInfo struct {
	UserName  string    `json:"user_name"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type JwtClaims struct {
	Info UserInfo
	jwt.RegisteredClaims
}

func GenerateToken(userName string, expiredAtMinutes uint32) (*UserInfo, error) {
	claims := JwtClaims{
		UserInfo{
			userName,
			"",
			time.Now().Add(time.Duration(expiredAtMinutes) * time.Minute),
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredAtMinutes) * time.Minute)), // 有效期，结束时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                    // 签发时间（比如：时间戳）
			NotBefore: jwt.NewNumericDate(time.Now()),                                                    // 有效期，开始时间（比如：时间戳）
			Issuer:    "Copyright © 2023 HRM Admin",                                                      // Token 颁发者的唯一标识
			Subject:   "Copyright © 2023 HRM Token Generate",                                             // 主题（比如：用户id 或 用户名）
			ID:        userName,                                                                          // jwt唯一标识（比如：UUID
			Audience:  []string{"HRM users"},                                                             // JWT的接收者（比如：APP的包名）
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	claims.Info.Token = bearerName + " " + token
	return &claims.Info, nil
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

func JwtAuthenticator(c *gin.Context) {
	// 从请求头中取出
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

	// 校验token
	claims, err := isTokenValid(ht[1])
	if err != nil {
		log.Warn("非法token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "非法token",
			"data":    "",
		})
		c.Abort()
		return
	}

	if claims.Info.ExpiredAt.Before(time.Now()) {
		log.Warn("token过期, ExpiredAt[%+v], Now[%+v]", claims.Info.ExpiredAt.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "token过期",
			"data":    "",
		})
		c.Abort()
		return
	}

	// 将用户的id放在到请求的上下文中
	c.Set("UserInfo", claims.Info)

	// 后续的处理函数可以用过c.Get("UserInfo")来获取当前请求的UserInfo
	c.Next()

}
