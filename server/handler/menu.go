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

func registerMenuRouter(r *gin.Engine) {
	menuRouter := r.Group("/menu")

	// 以下接口登陆后才能访问, 加中间件
	menuRouter.GET("list", middleware.AccessTokenAuthenticator, menuList)
	menuRouter.POST("add", middleware.AccessTokenAuthenticator, menuAdd)
	menuRouter.PUT("update", middleware.AccessTokenAuthenticator, menuUpdate)
	menuRouter.DELETE(":id", middleware.AccessTokenAuthenticator, menuDel)
}

func menuList(c *gin.Context) {
	menus := new([]db.Menu)
	err := db.Find1(menus, -1)
	if err != nil {
		log.Warn("菜单信息获取失败, %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("菜单信息获取失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "菜单信息获取成功",
		"data": gin.H{
			"total": len(*menus),
			"menus": menus,
		},
	})
}

func menuAdd(c *gin.Context) {
	r := new(db.Menu)
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
	log.Debug("menuAdd request data [%v]", r)

	menu := &db.Menu{
		Name:     r.Name,
		Url:      r.Url,
		ParentId: r.ParentId,
	}
	if err := db.Insert(menu); err != nil {
		log.Warn("菜单增加失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("菜单增加失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "菜单增加成功",
		"data":    menu,
	})
}

func menuUpdate(c *gin.Context) {
	r := new(db.Menu)
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
	log.Debug("menuUpdate request data [%v]", r)

	menu := &db.Menu{Id: r.Id, Name: r.Name, Url: r.Url, ParentId: r.ParentId}
	if err := db.UpdateRow(menu); err != nil {
		log.Warn("菜单更新失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("菜单更新失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "菜单更新成功",
		"data":    "",
	})
}

func menuDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
	}
	log.Debug("menuDel id [%v]", id)

	menu := &db.Menu{Id: id}
	err = db.Delete(menu, "id = ?", menu.Id)
	if err != nil {
		log.Warn("菜单删除失败")
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "菜单删除失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "菜单删除成功",
		"data":    "",
	})
}
