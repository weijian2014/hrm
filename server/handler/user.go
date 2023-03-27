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
	lr := new(db.UserLoginRequest)
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

	user := new(db.User)
	err := db.First(user, "name = ?", lr.Name)
	if err != nil || user.Ulr.Password != lr.Password {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("Find user in database, [%v]", user)
	token, err := middleware.GenerateToken(lr.Name)
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
		log.Debug("Generate token [%v] for username [%v]", token, lr.Name)
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

	user := &db.User{
		Ulr: db.UserLoginRequest{
			Name: username.(string),
		}}
	if err := db.First(user, "name = ?", user.Ulr.Name); err != nil {
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
		"data":    user.Data,
	})
}

func userAdd(c *gin.Context) {
	ulr := new(db.UserLoginRequest)
	if err := c.ShouldBindJSON(ulr); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("userAdd request data [%v]", ulr)

	// todo verify the username have add user permission

	user := &db.User{
		Ulr: db.UserLoginRequest{
			Name:     ulr.Name,
			Password: ulr.Password,
		}}

	// 检查用户已存在
	if err := db.First(user, "name = ?", user.Ulr.Name); err == nil {
		log.Warn("用户增加出错, 用户名已存在")
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "用户增加出错, 用户名已存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	if err := db.Insert(user); err != nil {
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
		"data": gin.H{
			"id":       user.Id,
			"username": user.Ulr.Name,
		},
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

	uur := new(db.UserUpdateRequest)
	if err := c.ShouldBindJSON(uur); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("login user[%v], userUpdate request[%v]", username.(string), uur)
	if username != "admin" && uur.Name != username.(string) {
		log.Warn("非超级管理员只能修改自己的密码")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "非超级管理员只能修改自己的密码",
			"data":    "",
		})
		c.Abort()
		return
	}

	user := &db.User{
		Id: uur.Id,
		Ulr: db.UserLoginRequest{
			Name:     uur.Name,
			Password: uur.Password,
		}}

	if err := db.UpdateSingleField(user, "password", user.Ulr.Password); err != nil {
		log.Warn("密码修改失败, err[%v]", err)
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
		"message": "密码修改成功, 下次登录生效",
		"data":    "",
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

	user := &db.User{Id: id}
	if err = db.Delete(user, "id = ?", id); err != nil {
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
		"data":    "",
	})
}
