package swagger

import (
	"encoding/json"
	"os"
	"strings"
	"webServer/src/common/config"
	. "webServer/src/common/myStruct"

	. "github.com/soekchl/myUtils"
)

var (
	definitions = make(map[string]Schema)
)

func init() {
	initDefinitions()
}

// 指定目录生成 swagger.json文档
func OutPutSwagger(serverMaps map[string]*ApiInfo, filePath string) {
	getCfgStrList := func(key string, s string) []string {
		v := config.GetString(key)
		if len(v) < 1 {
			return nil
		}
		return strings.Split(v, s)
	}

	s := Swagger{
		SwaggerVersion: "2.0",
		Host:           config.GetString("swagger.Host"),
		BasePath:       config.GetString("swagger.BasePath"),
		Schemes:        getCfgStrList("swagger.Schemes", ","),
		Consumes:       getCfgStrList("swagger.Consumes", ","),
		Produces:       getCfgStrList("swagger.Produces", ","),
		Infos:          getInfos(),
		Definitions:    definitions,
	}

	// path
	s.Paths = getPaths(serverMaps)
	// tag ?

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		Error(err)
		return
	}

	os.Remove(filePath)
	fi, err := os.Create(filePath)
	if err != nil {
		Notice("\n", string(b), "\n\n")
		Error(err)
		return
	}
	defer fi.Close()
	fi.Write(b)
}

func getPaths(serverMaps map[string]*ApiInfo) (m map[string]*Item) {
	m = make(map[string]*Item)
	for k, v := range serverMaps {
		m[k] = getApiItems(v)
	}
	return
}

func getApiItems(apiInfo *ApiInfo) (item *Item) {
	item = &Item{}
	normal := definitions["normalReturn"]
	parmErr := definitions["parmErrReturn"]
	authErr := definitions["authErrReturn"]
	for k, v := range apiInfo.Method {
		if !v {
			continue
		}
		parm := getParm(apiInfo)
		oper := &Operation{
			Tags:        []string{strings.Split(apiInfo.ApiName, "/")[1]},
			Summary:     apiInfo.Summary,
			Description: apiInfo.Description,
			Consumes:    []string{"application/json"},
			Produces:    []string{"application/json"},
			Parameters:  parm,
			Responses: map[string]Response{
				"200": Response{
					Ref:    "#/definitions/normalReturn",
					Schema: &normal,
				},
				"400": Response{
					Ref:    "#/definitions/parmErrReturn",
					Schema: &parmErr,
				},
				"401": Response{
					Ref:    "#/definitions/authErrReturn",
					Schema: &authErr,
				},
			},
		}

		switch strings.ToUpper(k) {
		case "GET":
			item.Get = oper
		case "POST":
			item.Post = oper
		case "DELETE":
			item.Delete = oper
		case "PUT":
			item.Put = oper
		}
	}
	return
}

func getParm(apiInfo *ApiInfo) (parms []Parameter) {
	for _, v := range apiInfo.Parms {
		in := "body" // query/body/path/formData/header
		if apiInfo.Method["GET"] {
			in = "query"
		}
		parms = append(parms, Parameter{
			Required:    v.Req,
			In:          in,
			Name:        v.Name,
			Description: v.Summary,
			Type:        strings.ToLower(v.Type),
		})
	}
	return
}

// 从配置文件获取 项目信息 和 作者信息
func getInfos() Information {
	return Information{
		Title:          config.GetString("swagger.Title"),
		Description:    config.GetString("swagger.Description"),
		Version:        config.GetString("swagger.Version"),
		TermsOfService: config.GetString("swagger.TermsOfService"),
		Contact: Contact{
			Name:  config.GetString("swagger.Contact.Name"),
			URL:   config.GetString("swagger.Contact.URL"),
			EMail: config.GetString("swagger.Contact.EMail"),
		},
	}
}

func initDefinitions() {
	properties := make(map[string]Propertie)
	properties["data"] = Propertie{
		Type:        "object",
		Example:     "{}",
		Description: "相应数据",
	}
	properties["timestamp"] = Propertie{
		Type:        "int32",
		Example:     "1552709427",
		Description: "时间戳",
	}
	properties["msg"] = Propertie{
		Type:        "string",
		Example:     "ok",
		Description: "消息",
	}
	properties["code"] = Propertie{
		Type:        "integer",
		Format:      "int32",
		Example:     "200",
		Description: "数据返回值",
	}
	definitions["normalReturn"] = Schema{
		Description: "通用返回值",
		Type:        "object",
		Title:       "normalReturn",
		Properties:  properties,
	}

	properties = make(map[string]Propertie)
	properties["data"] = Propertie{
		Type:        "object",
		Example:     "{}",
		Description: "空值",
	}
	properties["timestamp"] = Propertie{
		Type:        "int32",
		Example:     "1552709427",
		Description: "时间戳",
	}
	properties["msg"] = Propertie{
		Type:        "string",
		Example:     "参数必填",
		Description: "错误提示消息",
	}
	properties["code"] = Propertie{
		Type:        "integer",
		Format:      "int32",
		Example:     "400",
		Description: "数据返回值",
	}
	definitions["parmErrReturn"] = Schema{
		Description: "参数错误返回值",
		Type:        "object",
		Title:       "parmErrReturn",
		Properties:  properties,
	}

	properties = make(map[string]Propertie)
	properties["data"] = Propertie{
		Type:        "object",
		Example:     "{}",
		Description: "空值",
	}
	properties["timestamp"] = Propertie{
		Type:        "int32",
		Example:     "1552709427",
		Description: "时间戳",
	}
	properties["msg"] = Propertie{
		Type:        "string",
		Example:     "接口需要认证",
		Description: "错误提示消息",
	}
	properties["code"] = Propertie{
		Type:        "integer",
		Format:      "int32",
		Example:     "401",
		Description: "数据返回值",
	}
	definitions["authErrReturn"] = Schema{
		Description: "认证错误返回值",
		Type:        "object",
		Title:       "authErrReturn",
		Properties:  properties,
	}
}
