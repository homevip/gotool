package hfile

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
)

// 创建目录(深层)
func CreateDir(dirPath string) bool {

	if ok := fileutil.IsExist(dirPath); ok {
		return true
	}

	if err := fileutil.CreateDir(dirPath); err != nil {
		return false
	}

	return true
}

// 移除目录
func RemoveDir(dirPath string) bool {

	if err := os.Remove(dirPath); err != nil {
		return false
	}

	return true
}

// 创建文件 (/xxx/xxx/xxx.txt)
func CreateFile(filePath string) bool {

	if ok := CreateDir(filePath); !ok {
		return false
	}

	if ok := fileutil.CreateFile(filePath); !ok {
		return false
	}

	return true
}

// 写入文件
// filePath: /xxx/xxx/xxx.txt
// content: 写入的内容
// append: 是否追加写入(true:追加 否则不追加)
func WriteFile(filePath string, content string, append bool) bool {

	if err := fileutil.WriteStringToFile(filePath, content, append); err != nil {
		return false
	}

	return true
}

// 返回目录下所有文件(包括子目录)
func ListFileName(dirPath string) (list []string, err error) {

	// 递归遍历目录
	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		// 如果有错误，直接返回
		if err != nil {
			return err
		}

		// 只收集文件，跳过目录
		if !d.IsDir() {
			list = append(list, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}
