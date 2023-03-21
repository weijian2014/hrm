package handler

import (
	"hrm/db"
	"hrm/log"
	"hrm/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserName string `xml:"username" json:"username" description:"登录的用户名"`
	Password string `xml:"password" json:"password" description:"登录的密码"`
}

// type loginResponse struct {
// 	TokenHeader string `xml:"token_header" json:"token_header" description:"登录成功生成的令牌头"`
// 	Token       string `xml:"token" json:"token" description:"登录成功生成的令牌"`
// }

// type userInfoResponse struct {
// 	TokenHeader string `xml:"token_header" json:"token_header" description:"登录成功生成的令牌头"`
// 	Token       string `xml:"token" json:"token" description:"登录成功生成的令牌"`
// }

func registerUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("login", login)
	userRouter.GET("info", info)
}

func login(c *gin.Context) {
	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求数据格式错误",
			"data":    "",
		})
		return
	}
	log.Debug("Login request data [%v]", lr)

	u := new(db.User)
	err := u.Find(lr.UserName)
	if err != nil || u.Password != lr.Password {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		return
	}

	log.Debug("Find user in database, [%v]", u)
	token, err := token.GenerateToken(lr.UserName)
	if err != nil {
		log.Warn("系统无法生成token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统无法生成token",
			"data":    "",
		})
		return
	} else {
		log.Debug("Generate token [%v] for username [%v]", token, lr.UserName)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gin.H{
			"token_header": "hrm",
			"token":        token,
		},
	})
}

func info(c *gin.Context) {
	tk := c.Request.Header.Get("Authorization")
	if len(tk) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "请求未认证, 请在HTTP请求头中携带Key为Authorization的token",
			"data":    "",
		})
		return
	}

	username, err := token.IsTokenValid(tk)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "非法token",
			"data":    "",
		})
		return
	}

	log.Debug("User name [%v]", username)

	u := new(db.User)
	err = u.Find(username)
	if err != nil {
		log.Warn("用户不存在")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
			"data":    "",
		})
		return
	}

	log.Debug("Find user in database, [%v]", u)

	c.JSON(http.StatusOK, gin.H{
		"message": "用户合法",
		"data":    u.Data,
	})
}
