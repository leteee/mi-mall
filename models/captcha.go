package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// var store = base64Captcha.DefaultMemStore
var store base64Captcha.Store = RedisStore{}

func MakeCaptcha() (string, string, error) {

	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		//Fonts: []string{"wqy-micorhei.tc"},
	}

	var driver base64Captcha.Driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

func Verify(id string, verifyValue string) bool {
	return store.Verify(id, verifyValue, true)
}
