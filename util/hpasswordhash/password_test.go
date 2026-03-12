package hpasswordhash

import (
	"fmt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	pass := "123"
	s, _ := PasswordHash(pass)
	fmt.Printf("s: %v\n", s)
}
func TestPasswordVerify(t *testing.T) {

	password := "123"
	hash := "$2a$10$qE3jlCLTFW5PVPgrgC5NWeMP62yqjrsPeq2wG0Rfw6C8u8CjkBCkq"
	b := PasswordVerify(password, hash)
	fmt.Printf("b: %v\n", b)

}
