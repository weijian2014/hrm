package handler

import (
	"encoding/json"
	"errors"
	"hrm/common"
	"hrm/log"
	"io"
	"net/http"
)

type result struct {
	Response int         `json:"response"`
	Msg      interface{} `json:"msg"`
}

func apiCheckCookie(w http.ResponseWriter, r *http.Request) error {
	if ret := sessionMgr.CheckCookieValid(w, r); ret == "" {
		w.Header().Add("sessionstatus", "timeout")
		return errors.New("cookie check fail, set session timeout to response header")

		// 这里跳转不成功,原因可能是Redirect后加载的HTML无法发送给浏览器
		// http.Redirect(w, r, "/", http.StatusFound)
	}

	if !isLogin(r) {
		w.Header().Add("sessionstatus", "timeout")
		return errors.New("cookie check fail, Not login")
	}

	return nil
}

// string
func respString(w http.ResponseWriter, isOk int, str string) {
	f := &result{isOk, str}
	b, err := json.Marshal(f)
	if err != nil {
		log.Error("respString fail, %s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// json string
func respJson(w http.ResponseWriter, isOk int, jsonStr []byte) {
	var jf interface{}
	if err := json.Unmarshal(jsonStr, &jf); nil != err {
		log.Error("respJson fail1, %s\n", err.Error())
	}

	f := &result{isOk, jf}
	b, err := json.Marshal(f)
	if err != nil {
		log.Error("respJson fail2, %s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// array and struct
func respArrayOrStruct(w http.ResponseWriter, isOk int, arrOrStruct interface{}) {
	out := &result{isOk, arrOrStruct}
	b, err := json.Marshal(out)
	if err != nil {
		log.Error("respArrayOrStruct fail, %s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getPostData(r *http.Request, out interface{}) error {
	data, err := io.ReadAll(r.Body)
	if nil != err {
		return err
	}

	if len(data) == 0 {
		return errors.New("not post data")
	}

	err = json.Unmarshal(data, out)
	if nil != err {
		return err
	}

	return nil
}

func isLogin(r *http.Request) bool {
	cookie, err := r.Cookie(common.JsonConfigs.CookieName)
	if nil != err {
		return false
	}

	sessionId := cookie.Value
	_, ok := sessionMgr.GetSessionValues(sessionId, "isAdmin")
	return ok
}
