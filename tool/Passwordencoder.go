package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func EncoderSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	//由于是16进制需要转换
	s := hex.EncodeToString(sum)
	return string(s)
}

func MD5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	bydate := h.Sum(nil)
	result := fmt.Sprintf("%x", bydate)
	return result
}
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
