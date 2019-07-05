package markdown

import (
	"fmt"
	"os"
	"strings"
	. "webServer/src/common/myStruct"

	. "github.com/soekchl/myUtils"
)

var markdownFormat = "\n\n### %v\n\n%v\n\n**地址**：  `%v`\n\n**方法**： `%v`\n\n%v\n\n---\n\n"

// 指定目录生成 swagger.json文档
func OutPutMarkdown(serverMaps map[string]*ApiInfo, filePath string) {
	buff := ""

	for _, v := range serverMaps {
		parms := `
**formData/body参数**：

| 参数名 | 必须 | 类型 | 说明 |
| ---: | :---: | :---: | :--- |`
		for _, vv := range v.Parms {
			parms = fmt.Sprintf("%s\n| %v | %v | %v | **%v** |",
				parms,
				vv.Name,
				getBoolToStr(vv.Req),
				vv.Type,
				vv.GetDesc(),
			)
		}
		if len(v.Parms) < 1 {
			parms = "**参数：无**"
		}
		buff += fmt.Sprintf(markdownFormat,
			v.Summary,
			v.Description,
			v.ApiName,
			getMethodToString(v.Method),
			parms,
		)
	}

	os.Remove(filePath)
	fi, err := os.Create(filePath)
	if err != nil {
		Notice("\n", buff, "\n\n")
		Error(err)
		return
	}
	defer fi.Close()
	fi.WriteString(buff)
}

func getBoolToStr(f bool) string {
	if f {
		return "是"
	}
	return "否"
}

func getMethodToString(m map[string]bool) string {
	var list []string
	for k, _ := range m {
		list = append(list, k)
	}
	return strings.Join(list, "|")
}
