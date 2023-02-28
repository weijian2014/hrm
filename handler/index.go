package handler

import (
	"hrm/api"
	"hrm/log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if nil != apiCheckCookie(w, r) {
		return
	}

	action := p.ByName("action")
	switch action {
	case "allparts":
		infos, err := api.AllPartsRow(false)
		if nil != err {
			log.Error("IndexHandler allparts fail, %s\n", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("IndexHandler allparts ok!")
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
		infos, err := api.Paging(false, info)
		if nil != err {
			log.Error("IndexHandler paging fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("IndexHandler paging ok!")
		respArrayOrStruct(w, 1, infos)
	case "search":
		info := &api.SearchInfo{}
		if err := getPostData(r, info); nil != err {
			log.Error("Get post data fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("SearchInfo: %+v", info)
		infos, err := api.Search(info, false)
		if nil != err {
			log.Error("IndexHandler search fail, %s", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("IndexHandler search ok!")
		respArrayOrStruct(w, 1, infos)
	case "count":
		info, err := api.Count()
		if nil != err {
			log.Error("IndexHandler count fail, %s\n", err.Error())
			respString(w, 0, err.Error())
			return
		}

		log.Debug("CountInfo %+v", info)
		respArrayOrStruct(w, 1, info)
		return
	default:
		log.Error("IndexHandler invalid api\n")
		respString(w, 0, "IndexHandler invalid api")
	}
}
