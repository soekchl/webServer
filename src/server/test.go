package server

import (
	"fmt"
	"net/http"
	. "webServer/src/common/myStruct"
	"webServer/src/common/token"
	"webServer/src/common/tools"
)

func init() {
	setServer(&ApiInfo{
		Func:        testIndex,
		FuncName:    "testIndex",
		ApiName:     "/test",
		Summary:     "测试接口",
		Description: "测试服务是否正常启动接口",
		Method:      map[string]bool{"GET": true, "POST": true},
	})

	setServer(&ApiInfo{
		Func:        testLogin,
		FuncName:    "testLogin",
		ApiName:     "/test/login",
		Summary:     "测试登陆",
		Description: "测试登陆",
		Method:      map[string]bool{"GET": true, "POST": true},
		Parms: []ParmInfo{
			ParmInfo{Name: "account", Req: true, Summary: "帐号", Type: "string"},
			ParmInfo{Name: "password", Req: true, Summary: "密码", Type: "string"},
		},
	})

	setServer(&ApiInfo{
		Func:        testHome,
		FuncName:    "testHome",
		ApiName:     "/test/home",
		Summary:     "测试主页",
		Description: "测试主页返回相应 用户登陆数据",
		Method:      map[string]bool{"POST": true},
		NeedToken:   true,
	})
}

func testLogin(w http.ResponseWriter, r *http.Request, data *Data) {
	account := data.ParmMap["account"]
	password := data.ParmMap["password"]
	if account == "root" && password == "admin" {
		t := token.SetToken(1, fmt.Sprintf(`{"account":"%v", "id":1}`, account))
		tools.ReturnJson(w, 0, fmt.Sprintf(`{"token":"%v"}`, t))
		return
	}
	tools.ReturnJson(w, 401, "帐号或密码错误")
}

func testHome(w http.ResponseWriter, r *http.Request, data *Data) {
	tools.ReturnJson(w, 0, data.TokenInfo)
}

func testIndex(w http.ResponseWriter, r *http.Request, data *Data) {
	w.Write([]byte("Hello World"))
}
