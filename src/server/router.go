package server

import (
	"fmt"
	. "webServer/src/common/myStruct"
)

var (
	serverRouter = make(map[string]*ApiInfo)
)

func init() {
	// setServer(testIndex, "testIndex", "/test")
}

// swagger 文档输出
func SwaggerOut(filePath string) {

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

func GetServer(apiName string) *ApiInfo {
	return serverRouter[apiName]
}
