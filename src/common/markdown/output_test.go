package markdown

import (
	"testing"
	"webServer/src/common/config"
	. "webServer/src/common/myStruct"
)

func Test(t *testing.T) {
	config.Config("../../../config/config.ini") // config.ini

	m := addTestData()

	OutPutMarkdown(m, "./markdown.md")
}

func addTestData() map[string]*ApiInfo {
	serverRouter := make(map[string]*ApiInfo)
	serverRouter["/test"] = &ApiInfo{
		FuncName:    "testIndex",
		ApiName:     "/test",
		Summary:     "测试接口",
		Description: "测试服务是否正常启动接口",
		Method:      map[string]bool{"GET": true, "POST": true},
	}

	serverRouter["/test/login"] = &ApiInfo{
		FuncName:    "testLogin",
		ApiName:     "/test/login",
		Summary:     "测试登陆",
		Description: "测试登陆",
		Method:      map[string]bool{"GET": true, "POST": true},
		Parms: []ParmInfo{
			ParmInfo{Name: "account", Req: true, Summary: "帐号"},
			ParmInfo{Name: "password", Req: true, Summary: "密码"},
		},
	}

	serverRouter["/test/home"] = &ApiInfo{
		FuncName:    "testHome",
		ApiName:     "/test/home",
		Summary:     "测试主页",
		Description: "测试主页返回相应 用户登陆数据",
		Method:      map[string]bool{"POST": true},
		NeedToken:   true,
	}
	return serverRouter
}
