package hstring

import (
	"fmt"
	"gotool/util/htime"
	"testing"
)

// 获取 UUID
func TestUUID(t *testing.T) {

	s := UUID()
	fmt.Printf("s: %v\n", s)
}

// 更好的随机值
func TestEnUniqid(t *testing.T) {

	var (
		ti = htime.GetDay()
	)

	fmt.Printf("==============: %v\n", ti)

	s := EnUniqid()
	fmt.Printf("s: %v\n", s)
}

// 更好的随机值
func TestBuilderOrderSn(t *testing.T) {

	s := BuilderOrderSn()
	fmt.Printf("s: %v\n", s)
}
