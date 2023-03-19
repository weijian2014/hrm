package handler

import (
	"hrm/log"
	"hrm/token"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type LoginRequest struct {
	UserName string `xml:"username" json:"username" description:"登录的用户名"`
	Password string `xml:"password" json:"password" description:"登录的密码"`
}

func (LoginRequest) SwaggerDoc() map[string]string {
	return map[string]string{
		"username": "登录的用户名",
		"password": "登录的密码",
	}
}

type LoginResponse struct {
	Token string `xml:"token" json:"token" description:"登录成功生成的token"`
}

func registerLoginWebService() {
	ws := new(restful.WebService)

	ws.Path("/")
	ws.Route(
		ws.POST("login").Consumes(restful.MIME_JSON).Produces((restful.MIME_JSON)).To(login).Doc("登录接口").
			Param(ws.BodyParameter("username", "用户名").DataType("string")).
			Param(ws.BodyParameter("password", "密码").DataType("string")))
	restful.Add(ws)
}

func login(req *restful.Request, resp *restful.Response) {
	lr := new(LoginRequest)
	if err := req.ReadEntity(lr); err != nil {
		log.Warn("LoginRequest error")
		resp.WriteHeaderAndEntity(http.StatusBadRequest, LoginResponse{Token: ""})
		return
	}
	log.Debug("LoginRequest[%v]", lr)

	token, err := token.GenerateToken(lr.UserName)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, LoginResponse{Token: ""})
		return
	}

	resp.WriteHeaderAndEntity(http.StatusOK, LoginResponse{Token: token})
}
