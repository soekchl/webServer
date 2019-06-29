package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var (
	client       *redis.Client = nil
	NotConnError               = errors.New("请先配置redis连接在调用")
)

// address=localhost:6379
// password=
// dbSelect=0
func Conn(address, password string, dbSelect int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       dbSelect, // use default DB
	})
	err := client.BgSave().Err()
	if err != nil {
		client = nil
		return err
	}
	return nil
}

func Get(key string) string {
	if client == nil {
		return ""
	}
	val, err := client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func Set(key, value string) (bool, error) {
	if client == nil {
		return false, NotConnError
	}
	return SetEx(key, value, 0)
}

func SetNx(key, value string) (bool, error) {
	if client == nil {
		return false, NotConnError
	}
	return client.SetNX(key, value, 0).Result()
}

func SetEx(key, value string, saveSecond int) (bool, error) {
	if client == nil {
		return false, NotConnError
	}
	if saveSecond < 0 {
		saveSecond = 0
	}
	v, err := client.Set(key, value, time.Second*time.Duration(saveSecond)).Result()
	if err != nil {
		return false, err
	}
	if v == "OK" {
		return true, nil
	}
	return false, errors.New("None Error Set key")
}

func Del(key ...string) (int64, error) {
	if client == nil {
		return 0, NotConnError
	}
	return client.Del(key...).Result()
}
