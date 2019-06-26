package token

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
	redis "webServer/common/redis"
	tools "webServer/common/tools"

	. "github.com/soekchl/myUtils"
)

type saveToken struct {
	info         string
	saveTimeUnix int64
}

var (
	saveRedis      = false                       // save memory
	saveTimeSecond = 3600 * 2                    // 2小时
	saveKey        = make(map[string]*saveToken) // key -> jsonStr
	saveKeyMutex   sync.RWMutex                  // 读写锁
	redisKey       = "test"
)

func Config(redisSave bool, saveTokenSecond int, redisHeaderKey string) {
	saveRedis = redisSave
	saveTimeSecond = saveTokenSecond
	if len(redisHeaderKey) > 0 && redisHeaderKey[len(redisHeaderKey)-1] == ':' {
		redisKey = redisHeaderKey[:len(redisHeaderKey)-1]
		fmt.Printf("Change Redis Key %20v -> %-20v\n", redisHeaderKey, redisKey)
	} else {
		redisKey = redisHeaderKey
	}

}

func SetToken(id int, info string) (token string) {
	token = tools.Md5(fmt.Sprint(id, time.Now().UnixNano()+rand.Int63n(1000), redisKey))
	result := false

	if saveRedis {
		result = setTokenRedis(id, info, token)
	} else {
		result = setTokenMemory(id, info, token)
	}

	if !result {
		return ""
	}
	return
}

func DelToken(id int) {
	if saveRedis {
		oldToken := redis.Get(getIdKey(id))
		if len(oldToken) > 0 {
			redis.Del(getTokenKey(oldToken))
		}
	} else {
		idStr := fmt.Sprint(id)
		if saveKey[idStr] != nil {
			saveKeyMutex.Lock()
			delete(saveKey, saveKey[idStr].info)
			saveKeyMutex.Unlock()
		}
	}
}

func CheckToken(token string) string {
	if saveRedis {
		return checkTokenRedis(token)
	}
	return checkTokenMemory(token)
}

func checkTokenRedis(token string) string {
	info := redis.Get(getTokenKey(token))
	if len(info) > 1 {
		// check token and id is OK
		m := make(map[string]string)
		err := json.Unmarshal([]byte(info), &m)
		if err == nil {
			t := redis.Get(getIdStrKey(m["id"]))
			if t == token {
				return info
			}
		}
	}
	return ""
}

func checkTokenMemory(token string) (info string) {
	saveKeyMutex.RLock()
	defer saveKeyMutex.RUnlock()

	r := saveKey[token]
	if r != nil {
		if r.saveTimeUnix == 0 {
			info = r.info
		}
		if time.Now().Unix() <= r.saveTimeUnix+int64(saveTimeSecond) {
			info = r.info
		} else {
			// TODO del memory token data
		}
	}
	if len(info) > 1 { // check id and token is OK
		m := make(map[string]string)
		err := json.Unmarshal([]byte(info), &m)
		t := saveKey[m["id"]]
		if err == nil && t != nil && t.info == token {
			// NOTICE is OK
		} else {
			info = ""
		}
	}
	return
}

func setTokenRedis(id int, info string, token string) bool {
	tokenKey := getTokenKey(token)
	r, err := redis.SetEx(tokenKey, info, saveTimeSecond)
	if err != nil || !r {
		Debug("Err=", err, fmt.Sprintf(" id=%v info=%v token=%v", id, info, token))
		return false
	}

	DelToken(id) // del old token

	r, err = redis.Set(getIdKey(id), token)
	if err != nil || !r {
		Debug("Err=", err, fmt.Sprintf(" id=%v info=%v token=%v", id, info, token))
		redis.Del(tokenKey)
		return false
	}
	return true
}

func setTokenMemory(id int, info string, token string) bool {
	saveKeyMutex.Lock()
	defer saveKeyMutex.Unlock()

	idStr := fmt.Sprint(id)
	saveKey[token] = &saveToken{info: info, saveTimeUnix: time.Now().Unix()}
	if saveKey[idStr] != nil { // del old token
		delete(saveKey, saveKey[idStr].info)
	}
	saveKey[idStr] = &saveToken{info: token, saveTimeUnix: 0}
	return true
}

func getTokenKey(token string) string {
	return fmt.Sprint(redisKey, ":userlogin:", token)
}

func getIdKey(id int) string {
	return getIdStrKey(fmt.Sprint(id))
}

func getIdStrKey(id string) string {
	return fmt.Sprint(redisKey, ":id:", id)
}
