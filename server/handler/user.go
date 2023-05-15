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
	userRouter.POST("refresh", userRefresh)

	// 以下接口登陆后才能访问, 加中间件
	userRouter.POST("logout", userLogout)
	userRouter.GET("info", middleware.AccessTokenAuthenticator, userInfo)
	userRouter.POST("add", middleware.AccessTokenAuthenticator, userAdd)
	userRouter.PUT("update", middleware.AccessTokenAuthenticator, userUpdate)
	userRouter.DELETE(":id", middleware.AccessTokenAuthenticator, userDel)
}

type userLoginRequest struct {
	Name     string `json:"username" description:"登录的用户名"`
	Password string `json:"password" description:"登录的密码"`
}

type userUpdateRequest struct {
	Name              string `json:"username" description:"登录的用户名"`
	OldPassword       string `json:"old_password" description:"旧密码"`
	NewPassword       string `json:"new_password" description:"新密码"`
	NewPasswordRepeat string `json:"new_password_repeat" description:"确认密码"`
}

type userRefreshRequest struct {
	RefreshToken string `json:"refresh_token" description:"用于刷新Access Token的Token"`
}

func userLogin(c *gin.Context) {
	r := new(userLoginRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("userLogin request data [%v]", r)

	user := new(db.User)
	err := db.Take(user, "name = ?", r.Name)
	if err != nil || user.Password != r.Password {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("Find user in database, [%v]", user)
	info, err := middleware.GenerateToken(user.Id, user.Name, common.JsonConfigs.TokenExpiredSeconds)
	if err != nil {
		log.Warn("系统无法生成token")
		c.JSON(http.StatusOK, gin.H{
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

func userRefresh(c *gin.Context) {
	r := new(userRefreshRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("userRefresh request data [%v]", r)

	tokenInfo, err := middleware.RefreshToken(r.RefreshToken)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": "重新生成token失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("userRefresh [%+v]", tokenInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "重新生成token成功",
		"data":    tokenInfo,
	})
}

func userLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "退出成功",
		"data":    "",
	})
}

func userInfo(c *gin.Context) {
	info, isExists := c.Get("ContextInfo")
	if !isExists {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	contextInfo := info.(middleware.ContextInfo)

	user := new(db.User)
	if err := db.Take(user, "id = ?", contextInfo.UserId); err != nil {
		log.Warn("用户不存在")
		c.JSON(http.StatusOK, gin.H{
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
	r := new(userLoginRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("userAdd request data [%v]", r)

	if !isUserNameValid(r.Name) {
		log.Warn("用户增加出错, 无效用户名")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户增加失败, 用户名长度在4~20之间",
			"data":    "",
		})
		c.Abort()
		return
	}

	if !isPasswordValid(r.Password) {
		log.Warn("用户增加出错, 无效密码")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户增加失败, 密码长度在8~20之间",
			"data":    "",
		})
		c.Abort()
		return
	}

	// 检查用户已存在
	user := new(db.User)
	if err := db.Take(user, "name = ?", r.Name); err == nil {
		log.Warn("用户增加出错, 用户名已存在")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户增加出错, 用户名已存在",
			"data":    "",
		})
		c.Abort()
		return
	}

	info, isExists := c.Get("ContextInfo")
	if !isExists {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	contextInfo := info.(middleware.ContextInfo)
	if contextInfo.UserName != "admin" {
		log.Warn("非管理员权限不允许增加用户")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "非管理员权限不允许增加用户",
			"data":    "",
		})
		c.Abort()
		return
	}

	user.Name = r.Name
	user.Password = r.Password
	if err := db.Insert(user); err != nil {
		log.Warn("用户增加出错, %v", err)
		c.JSON(http.StatusOK, gin.H{
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
	info, isExists := c.Get("ContextInfo")
	if !isExists {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	contextInfo := info.(middleware.ContextInfo)

	r := new(userUpdateRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}

	log.Debug("login user info[%+v], userUpdate request[%+v]", contextInfo.UserName, r)
	if contextInfo.UserName != "admin" && r.Name != contextInfo.UserName {
		log.Warn("非超级管理员只能修改自己的密码")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "非超级管理员只能修改自己的密码",
			"data":    "",
		})
		c.Abort()
		return
	}

	if !isPasswordValid(r.NewPassword) {
		log.Warn("密码修改失败, 无效密码")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "密码修改失败, 密码长度在8~20之间",
			"data":    "",
		})
		c.Abort()
		return
	}

	user := new(db.User)
	err := db.Take(user, "name = ?", r.Name)
	if err != nil || user.Password != r.OldPassword {
		log.Warn("用户名或者密码不正确")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名或者密码不正确",
			"data":    "",
		})
		c.Abort()
		return
	}

	if err := db.UpdateSingleColumn(user, "password", user.Password); err != nil {
		log.Warn("密码修改失败, err[%v]", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotModified,
			"message": "密码修改失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "密码修改成功, 请重新登录",
		"data":    "",
	})
}

func userDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("Delete id [%v]", id)

	info, isExists := c.Get("ContextInfo")
	if !isExists {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户不存在",
			"data":    "",
		})
		c.Abort()
		return
	}
	contextInfo := info.(middleware.ContextInfo)
	if contextInfo.UserName != "admin" {
		log.Warn("非管理员权限不允许删除用户")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "非管理员权限不允许删除用户",
			"data":    "",
		})
		c.Abort()
		return
	}

	user := &db.User{Id: id}
	if err = db.Delete(user, "id = ?", id); err != nil {
		log.Warn("用户不存在")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
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

func isUserNameValid(n string) bool {
	len := len(n)
	return len >= 4 && len <= 20
}

func isPasswordValid(p string) bool {
	len := len(p)
	return len >= 8 && len <= 20
}
