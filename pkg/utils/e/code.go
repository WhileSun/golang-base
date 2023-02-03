package e

const LANG = "CN"

const (
	SUCCESS = 0
	FAILED  = 1

	ERROR_CAPTCHA_GENERATE = 10001
	ERROR_CAPTCHA_VERIFY   = 10002
	ERROR_ACCOUNT_LOGIN    = 10003
	ERROR_ACCOUNT_CLOSE    = 10004
	ERROR_LOGIN_AUTH       = 10008
	ERROR_API_PERMS        = 10009

	ERROR_API_PARAMS = 10011
	ERROR_DB_FIND    = 10012
	ERROR_DB_ADD     = 10013
	ERROR_DB_UPDATE  = 10014
	ERROR_DB_DELETE  = 10015
)

var CnMessage = map[uint]string{
	SUCCESS: "操作成功",
	FAILED:  "操作失败",

	ERROR_CAPTCHA_GENERATE: "登录验证码生成失败",
	ERROR_CAPTCHA_VERIFY:   "验证码验证不通过",
	ERROR_ACCOUNT_LOGIN:    "账号密码不正确",
	ERROR_ACCOUNT_CLOSE:    "账户已经被关闭访问，请联系管理员",
	ERROR_LOGIN_AUTH:       "请先登录系统",
	ERROR_API_PERMS:        "抱歉,无访问权限",

	ERROR_API_PARAMS: "提交参数不正确",
	ERROR_DB_FIND:    "数据查询异常",
	ERROR_DB_ADD:     "数据添加异常",
	ERROR_DB_UPDATE:  "数据更新异常",
	ERROR_DB_DELETE:  "数据删除异常",
}

var EnMessage = map[uint]string{
	SUCCESS: "succeed",
	FAILED:  "failed",
}

func GetMessage(code uint) string {
	if LANG == "CN" {
		message, ok := CnMessage[code]
		if !ok {
			return ""
		}
		return message
	} else {
		message, ok := EnMessage[code]
		if !ok {
			return ""
		}
		return message
	}
}
