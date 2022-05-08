package utils

import (
	"crypto/md5"
	"encoding/hex"
)


//将值转换成MD5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}