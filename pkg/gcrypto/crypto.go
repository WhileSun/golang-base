package gcrypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

//PwdEncode 密码加密
func PwdEncode(password string) string{
	return Md5Encode(Sha256Encode(password,"ws_oms"))
}

func Sha256Encode(value string, salt string) string{
	h := sha256.New()
	_, _ = h.Write([]byte(value))
	_, _ = io.WriteString(h, salt)
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return s
}

func Md5Encode(value string) string{
	w := md5.New()
	io.WriteString(w, value)
	bydate := w.Sum(nil)
	result := fmt.Sprintf("%x", bydate)
	return result
}