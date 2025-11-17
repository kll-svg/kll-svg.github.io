package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {

	sessionManager := NewSessionManager()
	sessionManager.AddSession("token1", "user1", "127.0.0.1")
	if session, ok := sessionManager.GetSession("token1"); ok {
		fmt.Println("找到户户:", session.UserId, session.Address, session.LoginTime.Format("2006-01-02 15:04:05"))
	}
	sessionManager.ClearExpiredSessions(time.Hour)

	sessionManager.RemoveSession("token1")

	//	嵌套map
	userPreferences := map[string]map[string]string{
		"user1": {
			"theme": "dark",
			"lang":  "en",
		},
	}
	fmt.Println(userPreferences)
	//	为为集合使用,可以去重。map中的key一一对应，可以把无序数组
	uniqueIds := make(map[string]bool)
	ids := []int{1, 2, 2, 4, 5}
	for _, id := range ids {
		uniqueIds[fmt.Sprintf("%d", id)] = true
	}
	fmt.Println(len(uniqueIds))
	//	键值翻转
	//	将map中的key和value交换

	//	map无序，不能依赖于key的顺序
	//	如果需要有序的key值，需要使用slice来存储key
	ageMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	keys := make([]string, 0)
	for k := range uniqueIds {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//	打印有序的key值
	for _, k := range keys {
		fmt.Println(k, ageMap[k])
	}
	//尽量避免使用，读写多读写场景下，使用sync.Map
	var syncMap sync.Map
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	syncMap.Store("key3", "value3")
	//取值
	fmt.Println(syncMap.Load("key3"))

	//	删除key
	syncMap.Delete("key2")
	fmt.Println(syncMap.Load("key2"))
	//	遍历map
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}

//用户会话管理，session

type UserSession struct {
	UserId           string
	LoginTime        time.Time
	LastActivityTime time.Time
	Address          string
}

type SessionManager struct {
	sessions map[string]*UserSession
	//解决map中的并发安全问题
	mutex sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*UserSession),
	}
}

func (sm *SessionManager) AddSession(token string, userId string, ipAddr string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	session := &UserSession{
		UserId:           userId,
		LoginTime:        time.Now(),
		LastActivityTime: time.Now(),
		Address:          ipAddr,
	}
	sm.sessions[token] = session
}

func (sm *SessionManager) GetSession(token string) (*UserSession, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	session, ok := sm.sessions[token]
	if ok {
		session.LastActivityTime = time.Now()
	}
	return session, ok
}

func (sm *SessionManager) RemoveSession(token string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	delete(sm.sessions, token)
}

func (sm *SessionManager) ClearExpiredSessions(expireTime time.Duration) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	//ticker := time.NewTicker(expireTime)
	for token, session := range sm.sessions {
		if time.Since(session.LastActivityTime) > expireTime {
			delete(sm.sessions, token)
		}
	}
}
