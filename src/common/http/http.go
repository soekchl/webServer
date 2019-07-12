// getComic project main.go
package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	. "github.com/soekchl/myUtils"
)

func test() {
	value := make(url.Values)
	value.Add("phone", "17712345678")
	value.Add("count", "3")
		
	url := "http://localhost:8080/test/getId"
	Warn(HttpReq(url, "get", nil, value))

  headers := map[string]string{"content-type": "application/x-www-form-urlencoded"}
	url = "http://localhost:8080/test"
	Warn(HttpReq(url, "post", headers, value))
}

func HttpReq(reqUrl, method string, headers map[string]string, value url.Values) (code int, body string, err error) {
	method = strings.ToUpper(method)
	client := &http.Client{}
	val := value.Encode()
	if method == "GET" {
		reqUrl = fmt.Sprintf("%s?%v", reqUrl, val)
		val = ""
	}
	req, err := http.NewRequest(method, reqUrl, strings.NewReader(val))
	if err != nil {
		return 0, "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	code = response.StatusCode
	body = string(b)
	return
}
