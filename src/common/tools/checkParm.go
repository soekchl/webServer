package tools

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	. "webServer/src/common/myStruct"
)

func CheckParm(r *http.Request, parms []ParmInfo) (parm map[string]string, err error) {
	parm = make(map[string]string)
	if len(parms) < 1 {
		return
	}

	for _, v := range parms {
		value := GetValue(r, v.Name)
		if v.Req && len(value) < 1 {
			err = fmt.Errorf("%v 必填", v.GetDesc())
			return
		}
		if strings.Index(strings.ToUpper(v.Type), "INT") >= 0 {
			_, err = strconv.Atoi(value)
			if err != nil {
				err = fmt.Errorf("%v 类型错误", v.GetDesc())
				return
			}
		}
		parm[v.Name] = value
	}
	return
}
