package handler

import (
	"hrm/db"
	"hrm/log"
	"hrm/middleware"
	"net/http"
	"strconv"

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

	// 该接口登陆后才能访问,加中间件
	userRouter.GET("info", middleware.JwtAuthenticator, info)
	userRouter.DELETE(":id", middleware.JwtAuthenticator, del)
}

func login(c *gin.Context) {
	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		return
	}
	log.Debug("Login request data [%v]", lr)

	u := &db.User{Name: lr.UserName}
	err := u.Find()
	if err != nil || u.Password != lr.Password {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("Find user in database, [%v]", u)
	token, err := middleware.GenerateToken(lr.UserName)
	if err != nil {
		log.Warn("系统无法生成token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "系统无法生成token",
			"data":    "",
		})
		c.Abort()
		return
	} else {
		log.Debug("Generate token [%v] for username [%v]", token, lr.UserName)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data": gin.H{
			"token_header": middleware.TokenHeader,
			"token":        token,
		},
	})
}

func info(c *gin.Context) {
	username, isExists := c.Get("username")
	if !isExists {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	u := &db.User{Name: username.(string)}
	err := u.Find()
	if err != nil {
		log.Warn("用户不存")
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "用户合法",
		"data":    u.Data,
	})
}

func del(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
	}
	log.Debug("Delete id [%v]", id)

	u := &db.User{Id: id}
	err = u.Delete()
	if err != nil {
		log.Warn("用户不存")
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
		"data":    u.Data,
	})
}
