package email

import (
	"testing"
)

func Test(t *testing.T) {
	Config("soekchl@163.com", "123456")
	t.Log(SendCodeInMail("soekchl@163.com", "soekchl", "123456"))
}
