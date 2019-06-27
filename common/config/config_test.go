package config

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	configName := "./config.ini"
	os.Remove(configName)
	f, err := os.Create(configName)
	if err != nil {
		t.Error(err)
		return
	}
	saveConfig(f)
	f.Close()

	Config(configName)

	t.Log(GetString("mysql.port") == ":3060")
	b, err := GetBool("server.bool")
	t.Log(b == false, err)
	i, err := GetInt("limit")
	t.Log(i == 8888, err)
}

func saveConfig(f *os.File) {
	f.WriteString(`server.bool	=	false	#	测试
mysql.port	= 	:3060	#	测试端口
limit		= 	8888	#	测试
`)
}
