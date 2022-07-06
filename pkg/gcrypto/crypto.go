package gcrypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/agclqq/goencryption"
	"github.com/whilesun/go-admin/pkg/gconf"
	"io"
)

var salt string

func init() {
	salt = gconf.Config.GetString("app.saltPwd")
	if salt == ""{
		salt = "ws_oms"
	}
}

//PwdEncode 密码加密
func PwdEncode(password string) string{
	return Md5Encode(Sha256Encode(password,salt))
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
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5Encode16(value string) string{
	return Md5Encode(value)[8:24]
}

func AesEncode(value string, key string, iv string)  string{

	cryptText, _ := goencryption.EasyEncrypt("aes/cbc/pkcs7/base64", value, key, iv)
	return cryptText
}

func AesDecode(cryptText string, key string, iv string)  string{
	value, _ := goencryption.EasyDecrypt("aes/cbc/pkcs7/base64", cryptText, key, iv)
	return value
}