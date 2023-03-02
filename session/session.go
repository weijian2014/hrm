/*
Thread safe
*/
package session

import (
	"crypto/rand"
	"encoding/base64"
	"hrm/log"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

func New(cn string) *SessionMgr {
	sessionMgr := &SessionMgr{cookieName: cn, maxLifeTime: 3600, sessions: make(map[string]*Session)}
	go sessionMgr.gc()
	return sessionMgr
}

type Session struct {
	id             string
	values         map[interface{}]interface{}
	lastAccessTime time.Time
}

type SessionMgr struct {
	cookieName  string
	mutex       sync.RWMutex
	maxLifeTime int64
	sessions    map[string]*Session //<sessionID, Session>
}

func (mgr *SessionMgr) gc() {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	for sessionId, pSession := range mgr.sessions {
		if pSession.lastAccessTime.Unix()+mgr.maxLifeTime > time.Now().Unix() {
			delete(mgr.sessions, sessionId)
		}
	}

	time.AfterFunc(time.Duration(mgr.maxLifeTime)*time.Second, func() { mgr.gc() })
}

func (mgr *SessionMgr) newSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano()
		return strconv.FormatInt(nano, 10)
	}

	return base64.URLEncoding.EncodeToString(b)
}

func (mgr *SessionMgr) StartSession(w http.ResponseWriter, r *http.Request) string {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	sessionId := url.QueryEscape(mgr.newSessionId())
	session := &Session{id: sessionId, lastAccessTime: time.Now(), values: make(map[interface{}]interface{})}
	mgr.sessions[sessionId] = session
	cookie := http.Cookie{Name: mgr.cookieName, Value: sessionId, Path: "/", HttpOnly: true, MaxAge: int(mgr.maxLifeTime)}
	http.SetCookie(w, &cookie)

	return sessionId
}

func (mgr *SessionMgr) StopSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		mgr.mutex.Lock()
		defer mgr.mutex.Unlock()

		delete(mgr.sessions, cookie.Value)

		expiration := time.Now()
		c := http.Cookie{Name: mgr.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &c)
	}
}

func (mgr *SessionMgr) SetSessionValuse(sessionId string, k interface{}, v interface{}) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	if s, ok := mgr.sessions[sessionId]; ok {
		s.values[k] = v
	}
}

func (mgr *SessionMgr) GetSessionValuse(sessionId string, k interface{}) (interface{}, bool) {
	mgr.mutex.RLock()
	defer mgr.mutex.RUnlock()
	if s, ok := mgr.sessions[sessionId]; ok {
		if v, ok := s.values[k]; ok {
			return v, true
		}
	}

	return nil, false
}

func (mgr *SessionMgr) DeleteSessionValuse(sessionId string, k interface{}) {
	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()
	delete(mgr.sessions[sessionId].values, k)
}

func (mgr *SessionMgr) CheckCookieValid(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(mgr.cookieName)
	if nil != err || nil == cookie {
		return ""
	}

	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	sessionId := cookie.Value
	if s, ok := mgr.sessions[sessionId]; ok {
		// session存在,说明合法,更新最后的访问时间
		s.lastAccessTime = time.Now()
		return sessionId
	}

	return ""
}

func (mgr *SessionMgr) Dump() {
	mgr.mutex.RLock()
	defer mgr.mutex.RUnlock()
	log.Info("Cookies name:%s, max life time:%d\n", mgr.cookieName, mgr.maxLifeTime)
	for id, s := range mgr.sessions {
		log.Info("Session ID:%s, last access time:%d, valuse count:%d\n", id, s.lastAccessTime.Unix(), len(s.values))
		for k, v := range s.values {
			log.Info("k:%s, v:%s\n", k, v)
		}
	}
}

func (mgr *SessionMgr) SessionCount() int {
	return len(mgr.sessions)
}
