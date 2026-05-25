package hcurl

import (
	"time"

	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ctx    = gctx.New()
	client = gclient.New()
)

type Options struct {
	Url             string         // 请求地址
	Params          any            // 参数
	Header          map[string]any // Header 信息
	Auth_Basic_User string         // Auth 校验信息
	Auth_Basic_Pass string         // Auth 校验信息
	Timeout         int            // 请求超时时间
}

func NewOption() *Options {
	return &Options{
		Timeout: 3, // 请求超时时间, 默认3秒
	}
}

// CURL Get 请求
func (o *Options) CurlGet() (res string, err error) {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		Get(ctx, o.Url, o.Params)

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL Get 请求
// 获取状态码
func (o *Options) GetStatusCode() int {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		Get(ctx, o.Url, o.Params)

	if err != nil {
		return 0
	}
	defer r.Close()

	return r.StatusCode
}

// CURL Post 请求
func (o *Options) CurlPost() (res string, err error) {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		PostForm(ctx, o.Url, gconv.MapStrStr(o.Params))

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL PostJson 请求
func (o *Options) CurlPostJson() (res string, err error) {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	JsonStr := gconv.String(o.Params)

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		Post(ctx, o.Url, JsonStr)

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL PutJson 请求
func (o *Options) CurlPutJson() (res string, err error) {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	JsonStr := gconv.String(o.Params)

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		Put(ctx, o.Url, JsonStr)

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL Delete 请求
func (o *Options) CurlDelete() (res string, err error) {

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			client.SetHeader(k, s[0])
		}
	}

	r, err := client.
		Timeout(time.Second*time.Duration(o.Timeout)).
		BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass).
		Delete(ctx, o.Url)

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}
