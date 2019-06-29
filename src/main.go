package main

import (
	"fmt"
	"webServer/src/common/config"
	"webServer/src/common/redis"
	"webServer/src/common/token"
	"webServer/src/common/tools"

	. "github.com/soekchl/myUtils"
)

var serverPort = ":8080"

func init() {
	configFile := "../config/config.ini"
	// 从启动参数 更改配置文件目录
	im := tools.GetInputArgs()
	if len(im["configFile"]) > 0 {
		configFile = im["configFile"]
		Warn("Config File Edit configFile=", configFile)
	}
	config.Config(configFile)
	ConnRedis()
	serverPort = config.GetString("server.port")

	b, _ := config.GetBool("token.saveRedis")
	st, _ := config.GetInt("token.saveSecond")
	token.Config(b, st, config.GetString("token.saveKey"))
}

func main() {
	StartServer(serverPort)
}

func ConnRedis() {
	addr := config.GetString("redis.server")
	pwd := config.GetString("redis.auth")
	db, err := config.GetInt("redis.db")
	if err != nil {
		db = 0
	}

	Warn(fmt.Sprintf("Redis server=%v pwd=%v db=%v", addr, pwd, db))
	err = redis.Conn(addr, pwd, db)
	if err != nil {
		panic(err)
	}
}
