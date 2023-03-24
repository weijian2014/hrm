package handler

import (
	"fmt"
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

func registerUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("login", userLogin)

	// 以下接口登陆后才能访问, 加中间件
	userRouter.GET("info", middleware.JwtAuthenticator, userInfo)
	userRouter.POST("add", middleware.JwtAuthenticator, userAdd)
	userRouter.PUT("update", middleware.JwtAuthenticator, userUpdate)
	userRouter.DELETE(":id", middleware.JwtAuthenticator, userDel)
}

func userLogin(c *gin.Context) {
	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("Login request data [%v]", lr)

	u := &db.User{Name: lr.UserName}
	if err := u.Find(); err != nil || u.Password != lr.Password {
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

func userInfo(c *gin.Context) {
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
	if err := u.Find(); err != nil {
		log.Warn("用户不存在")
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

func userAdd(c *gin.Context) {
	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("userAdd request data [%v]", lr)

	// todo verify the username have add user permission

	u := &db.User{Name: lr.UserName, Password: lr.Password}
	if err := u.Insert(); err != nil {
		log.Warn("用户增加出错, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("用户增加出错, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "用户增加成功",
		"data":    "",
	})
}

func userUpdate(c *gin.Context) {
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

	lr := new(loginRequest)
	if err := c.ShouldBindJSON(lr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}

	if username != "admin" || lr.UserName != username.(string) {
		log.Warn("非超级管理员只能修改自己的密码")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "非超级管理员只能修改自己的密码",
			"data":    "",
		})
		c.Abort()
		return
	}

	u := &db.User{Name: lr.UserName, Password: lr.Password}
	if err := u.Update(); err != nil {
		log.Warn("密码修改失败")
		c.JSON(http.StatusNotModified, gin.H{
			"code":    http.StatusNotModified,
			"message": "密码修改失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "密码修改成功",
		"data":    u.Data,
	})
}

func userDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("Delete id [%v]", id)

	u := &db.User{Id: id}
	if err = u.Delete(); err != nil {
		log.Warn("用户不存在")
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
		"message": "用户删除成功",
		"data":    u.Data,
	})
}
