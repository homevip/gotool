package hverify

import (
	"fmt"
	"testing"
)

// 判断是否是 json 格式
func TestIsJson(t *testing.T) {
	u := `{"name":"sojson"}`
	b := IsJson(u)
	fmt.Printf("b: %v\n", b)
}

// InSlice
func TestStrInSlice(t *testing.T) {
	arr1 := []string{"abc", "aaaa", "bbbb"}
	i, b := InSlice(arr1, "aaaa")
	fmt.Printf("i: %v\n", i)
	fmt.Printf("b: %v\n", b)
}
