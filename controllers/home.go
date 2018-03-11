package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gnenux/pic-bed/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["URL"] = &models.URL
	user := this.GetSession("user")
	if user != nil {
		this.Data["User"] = user.(models.User)
		imgDates, err := models.GetImageDate(user.(models.User).Id)
		if err != nil {
			this.Data["Error"] = err.Error()
		} else {
			this.Data["SiderNavs"] = imgDates
		}

		imgs, err := models.GetAllImage(user.(models.User).Id)
		if err != nil {
			this.Data["Error"] = err.Error()
		} else {
			this.Data["Images"] = imgs
		}

		this.TplName = "index.html"
	} else {
		this.TplName = "layout/signin.html"
	}
}
