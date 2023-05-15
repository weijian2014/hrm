package handler

import (
	"fmt"
	"hrm/db"
	"hrm/log"
	"hrm/middleware"
	"net/http"
	"strconv"
	"time"

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

type EmployeeWithClient struct {
	Id              uint64    `json:"id" description:"用户ID"`
	Name            string    `json:"name" description:"姓名"`
	Gender          string    `json:"gender" description:"性别"`
	Birthday        time.Time `json:"birthday" description:"出生日期"`
	Height          uint64    `json:"height" description:"身高"`
	Weight          uint64    `json:"weight" description:"体重"`
	Degree          string    `json:"degree" description:"学历"`
	Identifier      string    `json:"identifier" description:"身份证号"`
	Phone           string    `json:"phone" description:"手机号码"`
	PoliticalStatus string    `json:"political_status" description:"政治面貌"`
	SocialSecurity  string    `json:"social_security" description:"社保"`
	CurrentAddress  string    `json:"current_address" description:"现住址"`
	FirstWorkTime   time.Time `json:"first_work_time" description:"首次工作"`
	FormerEmployer  string    `json:"former_employer" description:"原单位"`
	PostId          uint64    `json:"post_id" description:"岗位ID"`
	Post            string    `json:"post" description:"岗位"`
	Salary          uint64    `json:"salary" description:"工资"`
	SecurityCard    string    `json:"security_card" description:"保安证"`
	Comments        string    `json:"comments" description:"备注"`
	UpdatedAt       time.Time `json:"updated_at" description:"更新时间"`
}

func employeeList(c *gin.Context) {
	employees := new([]db.Employee)
	err := db.Find2(employees,
		"Post",
		"Post.id == Employee.id",
		-1)
	if err != nil {
		log.Warn("职工信息获取失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNoContent,
			"message": fmt.Sprintf("职工信息获取失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	ers := convertToClient(*employees)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工信息获取成功",
		"data": gin.H{
			"total":     len(ers),
			"employees": ers,
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
		c.JSON(http.StatusOK, gin.H{
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
	err := db.Find2(employees,
		"Post",
		"Post.id == Employee.id",
		-1,
		"name || gender || birthday || height || weight || degree || identifier || phone || political_status || social_security || current_address || first_work_time || former_employer || post || salary || security_card || comments like ?",
		"%"+r.Key+"%")
	if err != nil {
		log.Warn("职工信息获取失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNoContent,
			"message": fmt.Sprintf("职工信息获取失败, %v", err),
			"data":    "",
		})
		c.Abort()
		return
	}

	ers := convertToClient(*employees)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "职工信息获取成功",
		"data": gin.H{
			"total":     len(ers),
			"employees": ers,
		},
	})
}

func employeeAdd(c *gin.Context) {
	r := new(EmployeeWithClient)
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
	log.Debug("employeeAdd request data [%v]", r)

	e := convertToServer([]EmployeeWithClient{*r})
	if err := db.Insert(e[0]); err != nil {
		log.Warn("职工增加失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
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
	r := new(EmployeeWithClient)
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
	log.Debug("employeeUpdate request data [%v]", r)

	e := convertToServer([]EmployeeWithClient{*r})
	if err := db.UpdateRow(e[0]); err != nil {
		log.Warn("职工更新失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusNotAcceptable,
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

func convertToClient(srcs []db.Employee) []EmployeeWithClient {
	var ewc []EmployeeWithClient
	var e EmployeeWithClient
	for _, v := range srcs {
		e.Id = v.Id
		e.Name = v.Name
		e.Gender = v.Gender
		e.Birthday = v.Birthday
		e.Height = v.Height
		e.Weight = v.Weight
		e.Degree = v.Degree
		e.Identifier = v.Identifier
		e.Phone = v.Phone
		e.PoliticalStatus = v.PoliticalStatus
		e.SocialSecurity = v.SocialSecurity
		e.CurrentAddress = v.CurrentAddress
		e.FirstWorkTime = v.FirstWorkTime
		e.FormerEmployer = v.FormerEmployer
		e.PostId = v.Post.Id
		e.Post = v.Post.Name
		e.Salary = v.Salary
		e.SecurityCard = v.SecurityCard
		e.Comments = v.Comments
		e.UpdatedAt = v.UpdatedAt

		ewc = append(ewc, e)
	}

	return ewc
}

func convertToServer(srcs []EmployeeWithClient) []db.Employee {
	var employees []db.Employee
	var e db.Employee
	for _, v := range srcs {
		e.Id = v.Id
		e.Name = v.Name
		e.Gender = v.Gender
		e.Birthday = v.Birthday
		e.Height = v.Height
		e.Weight = v.Weight
		e.Degree = v.Degree
		e.Identifier = v.Identifier
		e.Phone = v.Phone
		e.PoliticalStatus = v.PoliticalStatus
		e.SocialSecurity = v.SocialSecurity
		e.CurrentAddress = v.CurrentAddress
		e.FirstWorkTime = v.FirstWorkTime
		e.FormerEmployer = v.FormerEmployer
		e.PostId = v.PostId
		e.Salary = v.Salary
		e.SecurityCard = v.SecurityCard
		e.Comments = v.Comments
		e.UpdatedAt = v.UpdatedAt
		employees = append(employees, e)
	}

	return employees
}
