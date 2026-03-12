package hpasswordhash

import "golang.org/x/crypto/bcrypt"

// 字符串Hash加密
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码与加密的值是否相等
// 成功返回 true 否则false
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
