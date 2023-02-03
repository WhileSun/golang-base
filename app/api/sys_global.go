package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/e"
)

type SysGlobalApi struct {
}

// GetCaptcha 登录验证码
func (c *SysGlobalApi) GetCaptcha(req *gin.Context) {
	id, b64s, err := gctx.GCaptcha.Generate()
	if err != nil {
		gctx.Logger.Error("验证码生成失败 ->", err.Error())
		e.New(req).Msg(e.ERROR_CAPTCHA_GENERATE)
		return
	}
	e.New(req).Data(e.SUCCESS, map[string]string{"id": id, "src": b64s})
}
