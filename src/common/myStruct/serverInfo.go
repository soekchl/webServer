package myStruct

import (
	"net/http"
)

type Data struct {
	ParmMap   map[string]string
	TokenInfo string
}

type ApiFunc func(http.ResponseWriter, *http.Request, *Data)

type ParmInfo struct {
	Name    string // 名称
	Summary string // 摘要
	Req     bool   // 是否必填
}

type ApiInfo struct {
	Func        ApiFunc         // 跳转函数
	FuncName    string          // 函数名
	ApiName     string          // 接口名称
	Summary     string          // 摘要
	Description string          // 描述
	NeedToken   bool            // token验证
	Method      map[string]bool // 调用格式 [get post delete put]
	Parms       []ParmInfo      // 参数列表
}

func (this ParmInfo) GetDesc() string {
	if len(this.Summary) > 0 {
		return this.Summary
	}
	return this.Name
}
