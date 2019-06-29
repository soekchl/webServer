package redis

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(Get("test") == "")

	err := Conn("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(Get("test") == "")
	t.Log(SetEx("test", "asdfasdf", 22))
	t.Log(SetNx("test1", "asd"))
	t.Log(Set("test1", "asdfasdf"))
	h, err := SetNx("test1", "asd")
	t.Log(h == false, err == nil)
	n, err := Del("test", "test1", "testset")
	if n != 2 || err != nil {
		t.Error(n, err)
	}
}
