package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/gcaptcha"
	"github.com/whilesun/go-admin/pkg/gsys"
)

type Sys struct {
	Base
}

func NewSys() *Sys{
	return &Sys{}
}

//GetLoginCaptcha 登录验证码
func (c *Sys) GetLoginCaptcha(req *gin.Context){
	id,b64s,err := gcaptcha.Generate()
	if err != nil{
		gsys.Logger.Error("验证码生成失败 ->", err.Error())
		e.New(req).Msg(e.ERROR_CAPTCHA_GENERATE)
		return
	}
	e.New(req).Data(e.SUCCESS,map[string]string{"id":id,"src":b64s})
}


