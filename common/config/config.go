package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

const ()

var (
	configData     sync.Map
	configPathName = "./config/config.ini"
)

func Config(fileAddress string) {
	configPathName = fileAddress
	getConfigDatas()
}

// reload config
func LoadConfig() {
	getConfigDatas()
}

func GetString(key string) string {
	if temp, ok := configData.Load(key); ok {
		if val, ok := temp.(string); ok {
			return val
		}
	}
	return ""
}

func GetBool(key string) (bool, error) {
	n := GetString(key)
	if n == "" {
		return false, fmt.Errorf("Not Found!")
	}

	return strconv.ParseBool(n)
}

func GetInt(key string) (int, error) {
	n := GetString(key)
	if n == "" {
		return 0, fmt.Errorf("Not Found!")
	}

	return strconv.Atoi(n)
}

func readFile() string {
	fi, err := os.Open(configPathName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func getConfigDatas() {
	fileStr := readFile()
	for _, v := range strings.Split(fileStr, "\n") {
		if len(v) < 1 || v[0] == '#' {
			continue
		}

		// 删除注释
		n := strings.IndexByte(v, '#')
		if n != -1 {
			v = v[:n-1]
		}

		n = strings.IndexByte(v, '=')
		if n == -1 {
			continue
		}
		// \t ---> ""
		k := strings.Replace(v[:n], "\t", "", -1)
		val := strings.Replace(v[n+1:], "\t", "", -1)
		// " " ---> ""
		k = strings.Replace(k, " ", "", -1)
		val = strings.Replace(val, " ", "", -1)
		configData.Store(k, val)
	}
}
