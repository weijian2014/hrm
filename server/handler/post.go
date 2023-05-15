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

func registerPostRouter(r *gin.Engine) {
	roleRouter := r.Group("/post")

	// 以下接口登陆后才能访问, 加中间件
	roleRouter.GET("list", middleware.AccessTokenAuthenticator, postList)
	roleRouter.POST("add", middleware.AccessTokenAuthenticator, postAdd)
	roleRouter.PUT("update", middleware.AccessTokenAuthenticator, postUpdate)
	roleRouter.DELETE(":id", middleware.AccessTokenAuthenticator, postDel)
}

func postList(c *gin.Context) {
	posts := new([]db.Post)
	err := db.Find1(posts, -1)
	if err != nil {
		log.Warn("岗位信息获取失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNoContent,
			"message": fmt.Sprintf("岗位信息获取失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "岗位信息获取成功",
		"data": gin.H{
			"total": len(*posts),
			"posts": posts,
		},
	})
}

func postAdd(c *gin.Context) {
	r := new(db.Post)
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
	log.Debug("postAdd request data [%v]", r)

	if err := db.Insert(r); err != nil {
		log.Warn("岗位增加失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("岗位增加失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "岗位增加成功",
		"data":    "",
	})
}

func postUpdate(c *gin.Context) {
	r := new(db.Post)
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
	log.Debug("postUpdate request data [%v]", r)

	if err := db.UpdateRow(r); err != nil {
		log.Warn("岗位更新失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("岗位更新失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "岗位更新成功",
		"data":    "",
	})
}

func postDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
	}
	log.Debug("postDel id [%v]", id)

	employees := new([]db.Employee)
	err = db.Find1(employees, -1, "post_id = ?", id)
	if err != nil {
		log.Warn("岗位删除失败, err=%+v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": "岗位删除失败",
			"data":    "",
		})
		c.Abort()
		return
	}
	employeesBelongTo := len(*employees)
	if employeesBelongTo != 0 {
		log.Warn("岗位删除失败, 有%v个员工属于该岗位!", employeesBelongTo)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("岗位删除失败, 有%v个员工属于该岗位, 请修改这些员工岗位再删除!", employeesBelongTo),
			"data":    "",
		})
		c.Abort()
		return
	}

	p := &db.Post{Id: id}
	err = db.Delete(p, "id = ?", p.Id)
	if err != nil {
		log.Warn("岗位删除失败")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": "岗位删除失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "岗位删除成功",
		"data":    "",
	})
}
