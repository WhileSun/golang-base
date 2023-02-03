package dto

type AddUser struct {
	Username string   `form:"username" binding:"required"  label:"用户账号"`
	Realname string   `form:"realname" binding:"required" label:"用户名称"`
	Status   *bool    `form:"status" binding:"required" label:"用户状态"`
	RoleIds  []string `form:"role_ids" binding:"required"  label:"所属角色"`
}

type UpdateUser struct {
	Id int `form:"id" binding:"required,gt=0" label:"ID"`
	AddUser
}

type LoginUser struct {
	Username    string `form:"username" binding:"required"  label:"用户账号"`
	Password    string `form:"password" binding:"required"  label:"用户密码"`
	CaptchaId   string `form:"captcha_id" binding:"required" label:"验证码ID"`
	CaptchaCode string `form:"captcha_code" binding:"required" label:"验证码"`
}

type UpdateUserPasswd struct {
	OldPasswd string `form:"old_passwd" binding:"required"  label:"旧密码"`
	NewPasswd string `form:"new_passwd" binding:"required"  label:"新密码"`
}
