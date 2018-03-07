package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gnenux/pic-bed/models"
)

type ImgaeController struct {
	beego.Controller
}

// @Title UploadImage
// @Description upload images
// @Param images formData file true "body for user content"
// @Success 200 {int}
// @router / [post]
func (ic *ImgaeController) Post() {
	res := make(map[string]string)
	u := ic.GetSession("user")
	if u == nil {
		res["error"] = "please login"
		ic.ServeJSON()
		return
	}

	hs, err := ic.GetFiles("images")
	if err != nil {
		res["error"] = err.Error()
	} else {
		for _, v := range hs {
			_, err = models.AddOneImage(v, u.(models.User).Id)
			if err != nil {
				res["error"] = err.Error()
				break
			}
		}
	}

	ic.Data["json"] = res
	ic.ServeJSON()
}
