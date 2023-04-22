package handler

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	registerUserRouter(r)
	registerRoleRouter(r)
	registerMenuRouter(r)
	registerEmployeeRouter(r)
	registerPostRouter(r)
}
