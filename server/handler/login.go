package handler

import (
	"hrm/log"
	"hrm/token"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type loginRequest struct {
	UserName string `xml:"username" json:"username" description:"登录的用户名"`
	Password string `xml:"password" json:"password" description:"登录的密码"`
}

type loginResponse struct {
	TokenHeader string `xml:"token_header" json:"token_header" description:"登录成功生成的令牌头"`
	Token       string `xml:"token" json:"token" description:"登录成功生成的令牌"`
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
	lr := new(loginRequest)
	if err := req.ReadEntity(lr); err != nil {
		log.Warn("Login request data error")
		resp.WriteHeaderAndEntity(http.StatusBadRequest,
			responseData[loginResponse]{
				Code:    http.StatusBadRequest,
				Message: "Login request data error",
				Data: &loginResponse{
					TokenHeader: "", Token: "",
				}})
		return
	}
	log.Debug("Login request data [%v]", lr)

	token, err := token.GenerateToken(lr.UserName)
	if err != nil {
		log.Warn("Can not generate token")
		resp.WriteHeaderAndEntity(http.StatusInternalServerError,
			responseData[loginResponse]{
				Code:    http.StatusInternalServerError,
				Message: "Can not generate token",
				Data: &loginResponse{
					TokenHeader: "", Token: "",
				}})
		return
	} else {
		log.Debug("Generate token [%v] for username [%v]", token, lr.UserName)
	}

	resp.WriteHeaderAndEntity(http.StatusOK,
		responseData[loginResponse]{
			Code:    http.StatusOK,
			Message: "ok",
			Data: &loginResponse{
				TokenHeader: "hrm", Token: token,
			}})
}
