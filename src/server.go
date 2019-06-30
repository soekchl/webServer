package main

import (
	"fmt"
	"net/http"
	"strings"
	. "webServer/src/common/myStruct"
	"webServer/src/common/token"
	"webServer/src/common/tools"
	"webServer/src/server"

	. "github.com/soekchl/myUtils"
)

// 启动服务
func StartServer(port string) {
	http.HandleFunc("/", middleware)
	Warn("Server listen port = ", port)
	if strings.Index(port, ":") != 0 {
		port = ":" + port
	}
	if len(port) < 1 {
		panic("Config Need [server.port] ")
	}
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

// 中间件
func middleware(w http.ResponseWriter, r *http.Request) {
	// Warn(r.Header)
	// Warn(r.RemoteAddr)
	// Warn(r.Method)
	// Warn(r.Header.Get("Content-Type"))

	Debug(fmt.Sprintf("url=%v ip=%v", r.RequestURI, tools.GetRealIp(r)))

	urlList := strings.Split(r.RequestURI, "?")

	// 判断路由是否存在
	ff := server.GetServer(urlList[0])
	if ff == nil {
		tools.ReturnJson(w, 401, "无此接口")
		return
	}

	if !ff.Method[r.Method] {
		tools.ReturnJson(w, 400, "请求方式错误")
		return
	}

	ti := ""
	// 判断是否需要token
	if ff.NeedToken {
		t := r.Header.Get("token")
		if len(t) < 1 {
			tools.ReturnJson(w, 400, "接口需要认证")
			return
		}
		ti = token.CheckToken(t)
		if len(ti) < 1 {
			tools.ReturnJson(w, 401, "凭证失效")
			return
		}
	}

	// 判断参数是否正确？
	parm, err := checkParm(r, ff.Parms)
	if err != nil {
		tools.ReturnJson(w, 400, err.Error())
		return
	}
	Notice(ti, parm)
	ff.Func(w, r, &Data{TokenInfo: ti, ParmMap: parm})
}

func checkParm(r *http.Request, parms []ParmInfo) (parm map[string]string, err error) {
	parm = make(map[string]string)
	if len(parms) < 1 {
		return
	}

	for _, v := range parms {
		value := tools.GetValue(r, v.Name)
		if v.Req && len(value) < 1 {
			err = fmt.Errorf("%v 必填", v.GetDesc())
			return
		}
		parm[v.Name] = value
	}
	return
}
