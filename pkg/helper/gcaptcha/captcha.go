package gcaptcha

import (
	"github.com/mojocn/base64Captcha"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"image/color"
)

type configJsonBody struct {
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type CaptchaConfig struct {
	Type   string
	Height int
	Width  int
	Length int
	Store  base64Captcha.Store
}

func New(captchaKey string) *CaptchaConfig {
	if captchaKey == "" {
		captchaKey = "captcha"
	}
	captchaConfig := &CaptchaConfig{
		Type:   "digit",
		Height: 80,
		Width:  240,
		Length: 6,
	}
	gconfig.Config.UnmarshalKey(captchaKey, captchaConfig)
	store := base64Captcha.DefaultMemStore
	captchaConfig.Store = store
	return captchaConfig
}

func (captchaConfig *CaptchaConfig) Generate() (id string, b64s string, err error) {
	param := configJsonBody{
		DriverString: &base64Captcha.DriverString{
			Height:          captchaConfig.Height,
			Width:           captchaConfig.Width,
			Length:          captchaConfig.Length,
			NoiseCount:      0,
			Source:          "123456789qwertyuipkjhgfdsazxcvbnm",
			ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
			Fonts:           []string{"wqy-microhei.ttc"},
			BgColor:         &color.RGBA{254, 254, 254, 254},
		},
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   captchaConfig.Height,
			Width:    captchaConfig.Width,
			Length:   captchaConfig.Length,
			MaxSkew:  0.7,
			DotCount: 80,
		},
	}
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch captchaConfig.Type {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	case "digit":
		driver = param.DriverDigit
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, captchaConfig.Store)
	id, b64s, err = c.Generate()
	return id, b64s, err
}

func (captchaConfig *CaptchaConfig) Verify(Id string, VerifyValue string) (res bool) {
	if captchaConfig.Store.Verify(Id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
