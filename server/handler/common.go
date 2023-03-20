package handler

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	registerUserRouter(r)
}

// type responseData[T interface{}] struct {
// 	Message string `xml:"message" json:"message" description:"消息"`
// 	Data    *T     `xml:"data" json:"data" description:"数据"`
// }
