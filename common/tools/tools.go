package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	. "github.com/soekchl/myUtils"
)

func Md5(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

func JsonMarshal(obj interface{}) string {
	buff, err := json.Marshal(obj)
	if err != nil {
		Debug("obj=", fmt.Sprintf("%#v", obj), "  err=", err)
		return ""
	}
	return string(buff)
}
