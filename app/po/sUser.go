package po

// SUser 用户表
type SUser struct {
	BaseField
	Username string `json:"username"`
	Password string `json:"password"`
	Realname string `json:"realname"`
	Status   bool   `json:"status"`
}
