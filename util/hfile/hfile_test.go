package hfile

import (
	"fmt"
	"os"
	"testing"
)

// 创建目录(深层)
func TestCreateDir(t *testing.T) {

	var (
		pwd, _   = os.Getwd()
		dir_path = pwd + "/test_xxx/abc"
	)
	b := CreateDir(dir_path)
	fmt.Printf("b: %v\n", b)
}

// 删除目录
func TestRemoveDir(t *testing.T) {

	var (
		pwd, _   = os.Getwd()
		dir_path = pwd + "/test_xxx/abc"
	)

	b := RemoveDir(dir_path)
	fmt.Printf("b: %v\n", b)
}
