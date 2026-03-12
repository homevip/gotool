package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/guuid"
)

var (
	OK              = ResponseCode(200, "success")
	NeedRedirect    = ResponseCode(301, "need redirect")
	InvalidArgs     = ResponseCode(400, "invalid params")
	Unauthorized    = ResponseCode(401, "unauthorized")
	Forbidden       = ResponseCode(403, "forbidden")
	NotFound        = ResponseCode(404, "not found")
	Conflict        = ResponseCode(409, "entry exist")
	TooManyRequests = ResponseCode(429, "too many requests")
	ResultError     = ResponseCode(500, "response error")
	DatabaseError   = ResponseCode(598, "database error")
	CSRFDetected    = ResponseCode(599, "csrf attack detected")
	UserError       = ResponseCode(5001, "username or password error")
	CodeExpire      = ResponseCode(5002, "verification expire")
	CodeError       = ResponseCode(5003, "verification error")
	UserExist       = ResponseCode(5004, "user Exist")
)

type ResponseJson struct {
	Code    int    `json:"code"`               // 状态码
	Msg     string `json:"msg"`                // 描述信息
	Data    any    `json:"data"`               // 数据包
	TraceID string `json:"trace_id,omitempty"` // 跟踪ID
}

var (
	rj ResponseJson // 成功响应
)

// 成功响应
func OutSuccess(c *gin.Context, data any) {

	if rj.Msg == "" {
		rj.Msg = "成功"
	}

	template(c, rj.Code, rj.Msg, data)
}

// 失败响应
func OutError(c *gin.Context, err ResponseJson) {
	// template(c, err.Code, err.Msg, nil)
	template(c, err.Code, err.Msg, err.Data)
}

// 使用 code模板 响应
// 获取错误码对应的模板信息
func ResponseCode(code int, msg string) ResponseJson {

	uuid := gconv.String(guuid.New())
	return ResponseJson{
		code,
		msg,
		nil,
		uuid,
	}
}

// 模板
func template(c *gin.Context, code int, msg string, data any) {
	var (
		uuid, _ = c.Get("trace_id")
	)

	c.JSON(http.StatusOK, ResponseJson{
		code,
		msg,
		data,
		gconv.String(uuid),
	})
}
