package hmap

import (
	"fmt"
	"sort"
	"strings"
)

// 实现的php中http_build_query
func HttpBuildQuery(param map[string]any) (res string) {

	// 检查输入的 map 是否为空
	if len(param) == 0 {
		return ""
	}

	// 提取所有的键
	keys := make([]string, 0, len(param))
	for k := range param {
		keys = append(keys, k)
	}

	// 对键进行排序
	sort.Strings(keys)

	m := make([]string, 0, len(param))

	// 按照排序后的键顺序遍历 map
	for _, k := range keys {
		v := param[k]
		m = append(m, fmt.Sprintf("%s=%s", k, fmt.Sprint(v)))
	}
	res = strings.Join(m, "&")
	return
}
