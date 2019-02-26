package util

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func GetSignature(sign string) (string, string) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	_str := fmt.Sprintf("%s%s", sign, now)
	signature := Md5([]byte(_str))
	return now, signature
}
