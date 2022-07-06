package gcaptcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

func Generate() (id string, b64s string, err error) {
	param := configJsonBody{
		Id:          "",
		CaptchaType: "string",
		VerifyValue: "",
		DriverString: &base64Captcha.DriverString{
			Height:64,
			Width:200,
			Length:4,
			NoiseCount:0,
			Source:"123456789qwertyuipkjhgfdsazxcvbnm",
			ShowLineOptions:base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
			Fonts: []string{"wqy-microhei.ttc"},
			BgColor:&color.RGBA{254, 254, 254,254},
		},
	}
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return id, b64s, err
}

func Verify(Id string, VerifyValue string) (res bool) {
	if store.Verify(Id, VerifyValue, true){
		return true
	}else{
		return false
	}
}