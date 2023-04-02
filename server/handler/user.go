package handler

import (
	"fmt"
	"hrm/common"
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

type userLoginRequest struct {
	Name     string `json:"username" description:"登录的用户名"`
	Password string `json:"password" description:"登录的密码"`
}

type userUpdateRequest struct {
	Id       uint64 `json:"id" description:"用户ID"`
	Name     string `json:"username" description:"登录的用户名"`
	Password string `json:"password" description:"登录的密码"`
}

func userLogin(c *gin.Context) {
	r := new(userLoginRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("Login request data [%v]", r)

	user := new(db.User)
	err := db.Take(user, "name = ?", r.Name)
	if err != nil || user.Password != r.Password {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("Find user in database, [%v]", user)
	info, err := middleware.GenerateToken(user.Id, user.Name, common.JsonConfigs.TokenExpiredMinutes)
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
		log.Debug("Generate user info [%+v]", info)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data":    info,
	})
}

func userInfo(c *gin.Context) {
	info, isExists := c.Get("UserInfo")
	if !isExists {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	userInfo := info.(middleware.UserInfo)

	user := new(db.User)
	if err := db.Take(user, "name = ?", userInfo.UserName); err != nil {
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
		"message": "获取用户信息成功",
		"data":    user.Data,
	})
}

func userAdd(c *gin.Context) {
	ulr := new(userLoginRequest)
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

	// 检查用户已存在
	user := new(db.User)
	if err := db.Take(user, "name = ?", ulr.Name); err == nil {
		log.Warn("用户增加出错, 用户名已存在")
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "用户增加出错, 用户名已存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	user.Name = ulr.Name
	user.Password = ulr.Password
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
			"username": user.Name,
		},
	})
}

func userUpdate(c *gin.Context) {
	info, isExists := c.Get("UserInfo")
	if !isExists {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	userInfo := info.(middleware.UserInfo)

	r := new(userUpdateRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("login user info[%+v], userUpdate request[%+v]", userInfo.UserName, r)
	if userInfo.UserName != "admin" && r.Name != userInfo.UserName {
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
		Id:       r.Id,
		Name:     r.Name,
		Password: r.Password,
	}

	if err := db.UpdateSingleColumn(user, "password", user.Password); err != nil {
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
