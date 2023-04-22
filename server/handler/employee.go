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
	employeeRouter.GET("list", middleware.AccessTokenAuthenticator, employeeList)
	employeeRouter.POST("search", middleware.AccessTokenAuthenticator, employeeSearch)
	employeeRouter.POST("add", middleware.AccessTokenAuthenticator, employeeAdd)
	employeeRouter.PUT("update", middleware.AccessTokenAuthenticator, employeeUpdate)
	employeeRouter.DELETE(":id", middleware.AccessTokenAuthenticator, employeeDel)
}

func employeeList(c *gin.Context) {
	employees := new([]db.Employee)
	err := db.Find(employees, -1)
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

type searchRequest struct {
	Key string `json:"key" description:"模糊搜索的关键字"`
}

func employeeSearch(c *gin.Context) {
	r := new(searchRequest)
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
	log.Debug("employeeSearch request data [%v]", r)

	employees := new([]db.Employee)
	// 模糊查询
	err := db.Find(employees,
		-1,
		"name || gender || birthday || height || weight || degree || identifier || phone || political_status || social_security || current_address || first_work_time || former_employer || post || salary || security_card || comments like ?",
		"%"+r.Key+"%")
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
	r := new(db.Employee)
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
	log.Debug("employeeAdd request data [%v]", r)

	if err := db.Insert(r); err != nil {
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
	r := new(db.Employee)
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
	log.Debug("employeeUpdate request data [%v]", r)

	if err := db.UpdateRow(r); err != nil {
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
	err = db.Delete(m, "id = ?", m.Id)
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
