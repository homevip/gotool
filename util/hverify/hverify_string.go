package hverify

import (
	"encoding/json"
	"net"
	"regexp"
)

// 判断是否是 json 格式
// 存在 true 否则false
func IsJson(str string) bool {
	res := json.Valid([]byte(str))
	return res
}

// 是否在切片中
// 响应1:所在的索引
// 响应2:包含true 否则false
func InSlice[T comparable](slice []T, val T) (int, bool) {
	for index, item := range slice {
		if item == val {
			return index, true
		}
	}
	return -1, false
}

// 验证手机号
func Phone(s string) bool {

	matched, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, s)
	return matched
}

// 字母数字
func ZmSz(s string) bool {

	matched, _ := regexp.MatchString(`^[A-Za-z0-9]+$`, s)
	return matched
}

// 判断是否是IPv4
func IsIPv4(address string) bool {

	ip := net.ParseIP(address)
	if ip == nil {
		return false
	}

	return ip.To4() != nil
}

// 判断是否是IPv6
func IsIPv6(address string) bool {

	ip := net.ParseIP(address)
	if ip == nil {
		return false
	}

	return ip.To4() == nil && len(ip) == net.IPv6len
}
