package main

import (
	"os"
	"testing"
	"time"
	"webServer/common/config"
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

	config.Config(configName)

	go StartServer(config.GetString("server.port"))

	time.Sleep(time.Second * 1)

	os.Remove(configName) // del config.ini
}

func saveConfig(f *os.File) {
	f.WriteString(`server.port	= 	:8080	#	测试端口`)
}
