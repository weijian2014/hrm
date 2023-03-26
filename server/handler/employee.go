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

func registerEmployeeRouter(r *gin.Engine) {
	employeeRouter := r.Group("/employee")

	// 以下接口登陆后才能访问, 加中间件
	employeeRouter.GET("list", middleware.JwtAuthenticator, employeeList)
	employeeRouter.POST("add", middleware.JwtAuthenticator, employeeAdd)
	employeeRouter.PUT("update", middleware.JwtAuthenticator, employeeUpdate)
	employeeRouter.DELETE(":id", middleware.JwtAuthenticator, employeeDel)
}

func employeeList(c *gin.Context) {
	employees, err := db.SelectAllEmployee()
	if err != nil {
		log.Warn("职工信息获取失败, %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("职工信息获取失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工信息获取成功",
		"data": gin.H{
			"total": len(*employees),
			"rows":  employees,
		},
	})
}

func employeeAdd(c *gin.Context) {
	e := new(db.Employee)
	if err := c.ShouldBindJSON(e); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("employeeAdd request data [%v]", e)

	if err := e.Insert(); err != nil {
		log.Warn("职工增加失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("职工增加失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工增加成功",
		"data":    "",
	})
}

func employeeUpdate(c *gin.Context) {
	e := new(db.Employee)
	if err := c.ShouldBindJSON(e); err != nil {
		log.Warn("请求数据格式错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
		return
	}
	log.Debug("employeeUpdate request data [%v]", e)

	if err := e.Update(); err != nil {
		log.Warn("职工更新失败, %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": fmt.Sprintf("职工更新失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工更新成功",
		"data":    "",
	})
}

func employeeDel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
			"data":    "",
		})
		c.Abort()
	}
	log.Debug("employeeDel id [%v]", id)

	m := &db.Employee{Id: id}
	err = m.Delete()
	if err != nil {
		log.Warn("职工删除失败")
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "职工删除失败",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工删除成功",
		"data":    "",
	})
}
