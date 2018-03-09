package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gnenux/pic-bed/models"
)

type HomeController struct {
	beego.Controller
}

type URL struct {
	Signin string
	Signup string
}

func (this *HomeController) Get() {
	user := this.GetSession("user")
	if user != nil {
		this.Data["User"] = user.(models.User)
		imgDates, err := models.GetImageDate(user.(models.User).Id)
		if err != nil {

		} else {
			this.Data["SiderNavs"] = imgDates
		}
		this.TplName = "index.html"
	} else {
		this.Data["URL"] = &URL{
			Signin: "127.0.0.1:8080/v1/user/login",
			Signup: "127.0.0.1:8080/v1/user/signup",
		}
		this.TplName = "layout/signin.html"
	}
}
