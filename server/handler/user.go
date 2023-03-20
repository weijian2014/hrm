package handler

import (
	"hrm/log"
	"hrm/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserName string `xml:"username" json:"username" description:"登录的用户名"`
	Password string `xml:"password" json:"password" description:"登录的密码"`
}

type loginResponse struct {
	TokenHeader string `xml:"token_header" json:"token_header" description:"登录成功生成的令牌头"`
	Token       string `xml:"token" json:"token" description:"登录成功生成的令牌"`
}

type userInfoResponse struct {
	TokenHeader string `xml:"token_header" json:"token_header" description:"登录成功生成的令牌头"`
	Token       string `xml:"token" json:"token" description:"登录成功生成的令牌"`
}

func registerUser(r *gin.Engine) {
	r.POST("/user/login", login)
	r.POST("/user/info", info)
}

func login(c *gin.Context) {
	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("Login request data error")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Login request data error",
			"data":    "",
		})
		return
	}
	log.Debug("Login request data [%v]", lr)

	// // todo find in sqlite

	token, err := token.GenerateToken(lr.UserName)
	if err != nil {
		log.Warn("Can not generate token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not generate token",
			"data":    "",
		})
		return
	} else {
		log.Debug("Generate token [%v] for username [%v]", token, lr.UserName)
	}

	log.Warn("Can not generate token")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gin.H{
			"token_header": "hrm",
			"token":        token,
		},
	})
}

func info(c *gin.Context) {
}
