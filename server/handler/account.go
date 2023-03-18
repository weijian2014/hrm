package handler

import (
	"hrm/api"
	"hrm/common"
	"hrm/log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AccountHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	action := p.ByName("action")
	switch action {
	case "login":
		info := &api.LoginInfo{}
		if err := getPostData(r, info); err != nil {
			log.Error("Get post data fail for login,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("LoginInfo %+v", info)
		isAdmin, err := api.Login(info)
		if err != nil {
			log.Error("AccountHandler login fail, %v", err)
			respString(w, 0, err.Error())
			return
		}

		// 从session中获取跳转中的url
		cookie, err := r.Cookie(common.JsonConfigs.CookieName)
		if nil != err || nil == cookie {
			log.Error("AccountHandler login fail, the cookie error\n")
			respString(w, 0, "Cookie错误!")
			return
		}

		sessionId := cookie.Value
		// 设置一个用户类型的session值
		sessionMgr.SetSessionValues(sessionId, "isAdmin", isAdmin)
		u, ok := sessionMgr.GetSessionValues(sessionId, "jump")
		if !ok {
			u = "/index"
		}

		jumpUrl := u.(string)
		if isAdmin {
			jumpUrl = "/admin"
		} else {
			jumpUrl = "/index"
		}
		sessionMgr.DeleteSessionValues(sessionId, "jump")
		log.Debug("AccountHandler login ok, jumpUrl [%s]\n", jumpUrl)
		respString(w, 1, jumpUrl)
		return
	case "logout":
		if nil != apiCheckCookie(w, r) {
			return
		}

		// 删除Cookie
		cookie := http.Cookie{Name: common.JsonConfigs.CookieName, Path: "/", MaxAge: -1}
		http.SetCookie(w, &cookie)

		// 删除Session
		sessionMgr.StopSession(w, r)
		log.Debug("AccountHandler logout ok, delete session and return jump login page.\n")
		respString(w, 1, "/")
		return
	case "change":
		if nil != apiCheckCookie(w, r) {
			return
		}
		// todo
		return
	default:
		log.Error("AccountHandler invalid api\n")
		respString(w, 0, "AccountHandler invalid api")
	}
}
