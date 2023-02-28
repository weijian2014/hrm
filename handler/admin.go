package handler

import (
	"hrm/api"
	"hrm/common"
	"hrm/log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AdminHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if nil != apiCheckCookie(w, r) {
		return
	}

	cookie, err := r.Cookie(common.JsonConfigs.CookieName)
	if nil != err || nil == cookie {
		log.Error("AdminHandler get cookie fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sessionId := cookie.Value
	isAdmin, ok := sessionMgr.GetSessionValuse(sessionId, "isAdmin")
	if !ok {
		log.Error("AdminHandler get isAdmin fail\n")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// 一个非管理用户调用admin API时
	if !isAdmin.(bool) {
		log.Error("AdminHandler a illegal request!")
		http.NotFound(w, r)
		return
	}

	action := p.ByName("action")
	switch action {
	case "allparts":
		infos, err := api.AllPartsRow(true)
		if nil != err {
			log.Error("AdminHandler allparts fail, %s\n", err.Error())
			respString(w, 0, "AdminHandler allparts fail, "+err.Error())
			return
		}

		log.Debug("AdminHandler allparts ok!")
		respArrayOrStruct(w, 1, infos)
		return
	case "paging":
		info := &api.PagingInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail for search,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("PagingInfo: %+v", info)
		infos, err := api.Paging(true, info)
		if nil != err {
			log.Error("AdminHandler paging fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("AdminHandler paging ok!")
		respArrayOrStruct(w, 1, infos)
	case "search":
		info := &api.SearchInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail for search,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("SearchInfo: %+v", info)
		infos, err := api.Search(info, true)
		if nil != err {
			log.Error("AdminHandler search fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("AdminHandler search ok!")
		respArrayOrStruct(w, 1, infos)
	case "add":
		info := &api.PartsRowInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail for add,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("PartsRowInfo: %+v", info)
		insertInfo, err := api.InsertPartsRow(info)
		if nil != err {
			log.Error("AdminHandler add fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("AdminHandler add ok!")
		respArrayOrStruct(w, 1, insertInfo)
	case "delete":
		info := &api.PartsRowInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail for delete,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("PartsRowInfo: %+v", info)
		err := api.DeletePartsRow(info)
		if nil != err {
			log.Error("AdminHandler delete fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("AdminHandler delete ok!")
		respArrayOrStruct(w, 1, "Delete ok!")
	case "update":
		info := &api.PartsRowInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail for update,  %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("PartsRowInfo: %+v", info)
		err := api.UpdatePartsRow(info)
		if nil != err {
			log.Error("AdminHandler update fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("AdminHandler update ok!")
		respArrayOrStruct(w, 1, "Update ok!")
	default:
		log.Error("AdminHandler invalid api\n")
		respString(w, 0, "AdminHandler invalid api")
	}
}
