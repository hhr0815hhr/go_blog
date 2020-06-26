package util

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const Md5Salt = "mcsaflFF*321"

func BCrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash), err
}

func BCryptVerify(str, encodedStr []byte) bool {
	return bcrypt.CompareHashAndPassword(encodedStr, str) == nil
}

func Md5(str string) string {
	data := []byte(str + Md5Salt)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func Md5Verify(str, encodedStr string) bool {
	return Md5(str) == encodedStr
}
