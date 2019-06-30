package swagger

import (
	"testing"
	"webServer/src/common/config"
	. "webServer/src/common/myStruct"
)

func Test(t *testing.T) {
	config.Config("../../../config/config.ini") // config.ini

	OutPutSwagger(map[string]*ApiInfo{}, "./swagger.json")
}
