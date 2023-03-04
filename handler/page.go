/*
	Date: 2018.03.29
	Code by weijian 303101610@qq.com
	Copyright (C) 2018 weijian
*/

package handler

import (
	"hrm/common"
	"hrm/log"
	"hrm/session"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	sessionMgr *session.SessionMgr = nil
)

func init() {

}

func Init() {
	sessionMgr = session.New(common.JsonConfigs.CookieName)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if ret := sessionMgr.CheckCookieValid(w, r); ret == "" {
		sessionId := sessionMgr.StartSession(w, r)
		jump := "/index"
		log.Debug("Invalid cookie, start session [%s] and set jump to [%s]\n", sessionId, jump)
		sessionMgr.SetSessionValuse(sessionId, "jump", jump)
	}

	t, err := template.ParseFiles(common.CurrDir + "/web/login.html")
	if err != nil {
		log.Error("Parse login.html file failt, error %s\n", err.Error())
		return
	}

	t.Execute(w, nil)
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !isLogin(r) {
		log.Warn("Not login!")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if ret := sessionMgr.CheckCookieValid(w, r); ret == "" {
		sessionId := sessionMgr.StartSession(w, r)
		jump := r.RequestURI
		log.Debug("Invalid cookie, start session [%s] and set jump to [%s].\n", sessionId, jump)
		sessionMgr.SetSessionValuse(sessionId, "jump", jump)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	t, err := template.ParseFiles(common.CurrDir + "/web/index.html")
	if err != nil {
		log.Error("Parse index.html file failt, error %s\n", err.Error())
		return
	}

	t.Execute(w, nil)
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !isLogin(r) {
		log.Warn("Not login!")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if ret := sessionMgr.CheckCookieValid(w, r); ret == "" {
		sessionId := sessionMgr.StartSession(w, r)
		jump := r.RequestURI
		log.Debug("Invalid cookie, start session [%s] and set jump to [%s].\n", sessionId, jump)
		sessionMgr.SetSessionValuse(sessionId, "jump", jump)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, err := r.Cookie(common.JsonConfigs.CookieName)
	if nil != err || nil == cookie {
		log.Error("AdminPageHandler get cookie fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sessionId := cookie.Value
	isAdmin, ok := sessionMgr.GetSessionValuse(sessionId, "isAdmin")
	if !ok {
		log.Error("AdminPageHandler get isAdmin fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if !isAdmin.(bool) {
		log.Error("AdminPageHandler a illegal request!")
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles(common.CurrDir + "/web/admin.html")
	if err != nil {
		log.Error("Parse admin.html file failt, error %s\n", err.Error())
		return
	}

	t.Execute(w, nil)
}

func EmployeePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !isLogin(r) {
		log.Warn("Not login!")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if ret := sessionMgr.CheckCookieValid(w, r); ret == "" {
		sessionId := sessionMgr.StartSession(w, r)
		jump := r.RequestURI
		log.Debug("Invalid cookie, start session [%s] and set jump to [%s].\n", sessionId, jump)
		sessionMgr.SetSessionValuse(sessionId, "jump", jump)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cookie, err := r.Cookie(common.JsonConfigs.CookieName)
	if nil != err || nil == cookie {
		log.Error("EmployeePageHandler get cookie fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sessionId := cookie.Value
	isAdmin, ok := sessionMgr.GetSessionValuse(sessionId, "isAdmin")
	if !ok {
		log.Error("EmployeePageHandler get isAdmin fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if !isAdmin.(bool) {
		log.Error("EmployeePageHandler a illegal request!")
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles(common.CurrDir + "/web/employee.html")
	if err != nil {
		log.Error("Parse admin.html file failt, error %s\n", err.Error())
		return
	}

	t.Execute(w, nil)
}
