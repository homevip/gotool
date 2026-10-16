package hcurl

import (
	"time"

	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ctx = gctx.New()
)

type Options struct {
	Url             string         // 请求地址
	Params          any            // 参数
	Header          map[string]any // Header 信息
	AuthBasicUser   string         // Basic Auth 账号
	AuthBasicPass   string         // Basic Auth 密码
	AuthBearerToken string         // Bearer Token
	Timeout         int            // 请求超时时间
}

func NewOption() *Options {
	return &Options{
		Timeout: 3, // 请求超时时间, 默认3秒
	}
}

// 创建客户端
func (o *Options) newClient() *gclient.Client {

	client := gclient.New()

	// Header
	if o.Header != nil {
		for k, v := range o.Header {
			s := gconv.Strings(v)
			if len(s) > 0 {
				client.SetHeader(k, s[0])
			}
		}
	}

	// Basic Auth
	if o.Auth_Basic_User != "" {
		client.BasicAuth(o.Auth_Basic_User, o.Auth_Basic_Pass)
	}

	// Bearer Token
	if o.Auth_Bearer_Token != "" {
		client.SetHeader("Authorization", "Bearer "+o.Auth_Bearer_Token)
	}

	// 请求超时时间
	if o.Timeout > 0 {
		client.Timeout(time.Second * time.Duration(o.Timeout))
	}

	return client
}

// CURL Get 请求
func (o *Options) CurlGet() (res string, err error) {

	client := o.newClient()

	r, err := client.Get(ctx, o.Url, o.Params)

	if err != nil {
		return
	}

	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL Get 请求
// 获取状态码
func (o *Options) GetStatusCode() (code int, err error) {

	client := o.newClient()

	r, err := client.Get(ctx, o.Url, o.Params)

	if err != nil {
		return 0, err
	}
	defer r.Close()

	return r.StatusCode, nil
}

// CURL Post 请求
func (o *Options) CurlPost() (res string, err error) {

	client := o.newClient()

	r, err := client.PostForm(ctx, o.Url, gconv.MapStrStr(o.Params))

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL PostJson 请求
func (o *Options) CurlPostJson() (res string, err error) {

	client := o.newClient()

	r, err := client.Post(ctx, o.Url, gconv.String(o.Params))

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL PutJson 请求
func (o *Options) CurlPutJson() (res string, err error) {

	client := o.newClient()

	r, err := client.Put(ctx, o.Url, gconv.String(o.Params))

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}

// CURL Delete 请求
func (o *Options) CurlDelete() (res string, err error) {

	client := o.newClient()

	r, err := client.Delete(ctx, o.Url)

	if err != nil {
		return
	}
	defer r.Close()

	res = r.ReadAllString()
	return
}
