package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// TODO config init
	go main() // start server
	time.Sleep(time.Second * 1)

	url := "http://localhost:8080"
	c, b, err := httpGet(url)
	if c != 401 {
		t.Error(c, b, err)
		return
	} else {
		t.Log("---> OK\t", url)
	}

	url = "http://localhost:8080/test"
	c, b, err = httpGet(url)
	if c != 200 || b != "Hello World" || err != nil {
		t.Error(c, b, err)
		return
	} else {
		t.Log("---> OK\t", url)
	}

	url = fmt.Sprintf("http://localhost:8080/test/login?account=%v&password=%v", "root", "admin")
	c, b, err = httpGet(url)
	if c != 200 || err != nil {
		t.Error(c, b, err)
		return
	} else {
		t.Log("---> OK\t", url)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(b), &m)
	if err != nil {
		t.Error(err, b)
	}
	md := m["data"].(map[string]interface{})
	token := md["token"].(string)
	t.Log(token)

	url = fmt.Sprintf("http://localhost:8080/test/home")
	c, b, err = HttpPost(url, map[string]string{"token": token})
	if c != 200 || err != nil {
		t.Error(c, b, err)
		return
	} else {
		t.Log("---> OK\t", url)
	}
	t.Log(b)

	time.Sleep(time.Second * 3)
}

func httpGet(url string) (code int, body string, err error) {
	response, err := http.Get(url)
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
	if response.StatusCode != 200 {
		code = response.StatusCode
		return
	}
	body = string(b)
	code = 200
	return
}

func HttpPost(url string, headers map[string]string) (code int, body string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	if resp.StatusCode != 200 {
		code = resp.StatusCode
		return
	}
	body = string(b)
	code = 200
	return
}
