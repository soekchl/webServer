package memory

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	saveSecond = 5
	Store("tset", "asdf")

	time.Sleep(time.Second * 6)
	Store("tset1", "asdf")
	Store("tset", "asdf1")
	time.Sleep(time.Second * 3)
	Load("test")
	Store("tset1", "asdf")
	time.Sleep(time.Second * 60)
}
