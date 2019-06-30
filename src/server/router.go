package server

import (
	"fmt"
	"net/http"
	"strings"
	. "webServer/src/common/myStruct"
	"webServer/src/common/swagger"
	"webServer/src/common/token"
	"webServer/src/common/tools"

	. "github.com/soekchl/myUtils"
)

var (
	serverRouter = make(map[string]*ApiInfo)
)

// swagger 文档输出
func SwaggerOut(filePath string) {
	swagger.OutPutSwagger(serverRouter, filePath)
}

// markdown 文档输出
func MarkdownOut(filePath string) {

}

// 设置服务路由 f-函数 funcName-函数名 apiName-映射名臣
func setServer(apiInfo *ApiInfo) {
	if serverRouter[apiInfo.ApiName] != nil {
		panic(fmt.Sprintf("API 已存在！apiName=%v", apiInfo.ApiName))
	}
	serverRouter[apiInfo.ApiName] = apiInfo
}

// 中间件
func Middleware(w http.ResponseWriter, r *http.Request) {
	// Warn(r.Header)
	// Warn(r.RemoteAddr)
	// Warn(r.Method)
	// Warn(r.Header.Get("Content-Type"))

	Debug(fmt.Sprintf("url=%v ip=%v", r.RequestURI, tools.GetRealIp(r)))

	urlList := strings.Split(r.RequestURI, "?")

	// 判断路由是否存在
	ff := serverRouter[urlList[0]]
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
	parm, err := tools.CheckParm(r, ff.Parms)
	if err != nil {
		tools.ReturnJson(w, 400, err.Error())
		return
	}
	Notice(ti, parm)
	ff.Func(w, r, &Data{TokenInfo: ti, ParmMap: parm})
}
