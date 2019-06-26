package token

import (
	"testing"
	"webServer/common/redis"
)

func Test(t *testing.T) {
	Config(false, 3600, "")
	token := SetToken(1, `{"id":"1"}`)
	t.Log(token)

	t.Log(CheckToken(token) == `{"id":"1"}`)
	t.Log(CheckToken("5da9a5284394ce77aad864dfdcec6fb5") == "")

	newToken := SetToken(1, `{"id":"1", "info":"new ID"}`)
	t.Log(newToken)
	t.Log(CheckToken(token) == "")
	// t.Log(saveKey[token], saveKey[newToken], saveKey["1"])
	t.Log(CheckToken(newToken) == `{"id":"1", "info":"new ID"}`)
	// t.Log(saveKey)
	DelToken(1)
	t.Log(CheckToken(token) == "")
	t.Log(CheckToken(newToken) == "")

	t.Log("---- redis test ---")

	err := redis.Conn("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		return
	}
	Config(true, 3600, "test:")
	token = SetToken(1, `{"id":"1"}`)
	t.Log(token)

	t.Log(CheckToken(token) == `{"id":"1"}`)
	t.Log(CheckToken("5da9a5284394ce77aad864dfdcec6fb5") == "")

	newToken = SetToken(1, `{"id":"1", "info":"new ID"}`)
	t.Log(newToken)
	t.Log(CheckToken(token) == "")
	t.Log(CheckToken(newToken) == `{"id":"1", "info":"new ID"}`)
	DelToken(1)
	t.Log(CheckToken(token) == "")
	t.Log(CheckToken(newToken) == "")
	SetToken(1, `{"id":"1", "info":"new ID"}`)
}
