package hcurl

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	url   = "http://192.168.xx.164:10102"
	token = "xxxx"
)

// 获取状态码
func TestStatusCode(t *testing.T) {
	co := &Options{
		Url: url,
	}
	code := co.GetStatusCode()
	fmt.Printf("code: %v\n", code)
}

// Get 请求
func TestGet(t *testing.T) {
	o := Options{
		Url:    url + "/api/department?",
		Params: map[string]any{
			// "keyword": "总",
		},
		Header: map[string]any{
			"token": token,
		},
	}
	s, _ := o.CurlGet()
	fmt.Printf("s: %v\n", s)
}

// Post 请求
func TestPost(t *testing.T) {
	o := Options{
		Url: url + "/api/login",
		Params: map[string]any{
			"account":  "xxx",
			"password": "xxx",
		},
	}
	s, _ := o.CurlPost()
	fmt.Printf("s: %v\n", s)
}

// Post Json参数请求
func TestPostJson(t *testing.T) {

	JsonParam := fmt.Sprintf(`{"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"%v"}]}`, "先有鸡还是先有蛋")

	// 将map 参数转 json
	maps := map[string]any{}
	err := json.Unmarshal([]byte(JsonParam), &maps)
	if err != nil {
		panic("报错")
	}

	o := Options{
		Url:    url,
		Params: maps,
	}
	s, _ := o.CurlPostJson()

	fmt.Printf("s: %v\n", s)

}
