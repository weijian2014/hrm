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

type roleAddRequest struct {
	Name string `xml:"role_name" json:"role_name" description:"角色名"`
}

type roleUpdateRequest struct {
	Id   uint64 `xml:"role_id" json:"role_id" description:"角色ID"`
	Name string `xml:"role_name" json:"role_name" description:"角色名"`
}

func registerRoleRouter(r *gin.Engine) {
	roleRouter := r.Group("/role")

	// 以下接口登陆后才能访问, 加中间件
	roleRouter.POST("add", roleAdd)
	roleRouter.PUT("update", middleware.AccessTokenAuthenticator, roleUpdate)
	roleRouter.DELETE(":id", middleware.AccessTokenAuthenticator, roleDel)
}

func roleAdd(c *gin.Context) {
	r := new(roleAddRequest)
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
	log.Debug("roleAdd request data [%v]", r)

	role := &db.Role{Name: r.Name}
	if err := db.Insert(role); err != nil {
		log.Warn("角色增加失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("角色增加失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "角色增加成功",
		"data":    "",
	})
}

func roleUpdate(c *gin.Context) {
	r := new(roleUpdateRequest)
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
	log.Debug("roleUpdate request data [%v]", r)

	role := &db.Role{Id: r.Id, Name: r.Name}
	if err := db.UpdateRow(role); err != nil {
		log.Warn("角色更新失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("角色更新失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "角色更新成功",
		"data":    "",
	})
}

func roleDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
	}
	log.Debug("roleDel id [%v]", id)

	role := &db.Role{Id: id}
	err = db.Delete(role, "id = ?", role.Id)
	if err != nil {
		log.Warn("角色删除失败")
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "角色删除失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "角色删除成功",
		"data":    role.Id,
	})
}
