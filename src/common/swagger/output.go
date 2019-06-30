package swagger

import (
	"encoding/json"
	"strings"
	"webServer/src/common/config"
	. "webServer/src/common/myStruct"

	. "github.com/soekchl/myUtils"
)

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
		Host:     config.GetString("swagger.Host"),
		BasePath: config.GetString("swagger.BasePath"),
		Schemes:  getCfgStrList("swagger.Schemes", ","),
		Consumes: getCfgStrList("swagger.Consumes", ","),
		Produces: getCfgStrList("swagger.Produces", ","),
		Infos:    getInfos(),
	}

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		Error(err)
		return
	}
	Notice("\n", string(b))
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
