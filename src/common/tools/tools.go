package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	. "github.com/soekchl/myUtils"
)

func GetInputArgs() map[string]string {
	m := make(map[string]string)
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			l := strings.Split(os.Args[i], "=")
			m[l[0]] = l[1]
		}
	}
	Warn(fmt.Sprintf("Input Arg %#v", m))
	return m
}

// code=0 dataJson=msg  code!=0 dataJson=json(obj)
func ReturnJson(w http.ResponseWriter, code int, dataJson string) {
	w.Header().Set("Content-Type", "application/json") // 设置返回值为JSON
	msg := "ok"
	data := "{}"
	if code == 0 {
		code = 200
		data = dataJson
		if len(data) < 1 {
			data = "{}"
		}
	} else {
		msg = dataJson
	}
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"code":%v,"data": %v,"msg":"%v","timestamp" : %v}`,
		code,
		data,
		msg,
		time.Now().Unix(),
	)))
}

// 从请求中获取相应参数，先获取post参数
func GetValue(r *http.Request, valueName string) string {
	value := r.PostFormValue(valueName)
	if len(value) < 1 {
		value = r.FormValue(valueName)
	}
	return value
}

// 获取当前请求ip
func GetRealIp(r *http.Request) (ip string) {
	if r == nil {
		return ""
	}
	if r.Header != nil {
		ip = r.Header.Get("X-Real-IP")
	}
	if len(ip) > 0 {
		return ip
	}

	host := r.Host
	if len(host) > 0 {
		list := strings.Split(host, ":")
		return list[0]
	}
	Warn("没有查找到 IP地址 ", fmt.Sprintf("%#v", r))
	return ""
}

func Md5(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

func JsonMarshal(obj interface{}) string {
	buff, err := json.Marshal(obj)
	if err != nil {
		Debug("obj=", fmt.Sprintf("%#v", obj), "  err=", err)
		return ""
	}
	return string(buff)
}
