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
// @Success 201 {int}
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
	ic.Ctx.ResponseWriter.Status = 201
	ic.ServeJSON()
}

// @Title DeleteImage
// @Description delete image
// @Param iid string  true "imageid for delete"
// @Success 204 {int}
// @router /:iid [delete]
func (ic *ImgaeController) Delete() {
	iid := ic.GetString(":iid")
	res := make(map[string]string)

	u := ic.GetSession("user")
	if u == nil {
		ic.Abort("403")
		return
	}

	err := models.DeleteImage(iid, u.(models.User).Id)
	if err != nil {
		res["error"] = err.Error()
	} else {
		ic.Ctx.ResponseWriter.Status = 204
	}

	ic.Data["json"] = res
	ic.ServeJSON()
}

// @Title ViewImage
// @Description view image
// @Param path query string  true "imageid for delete"
// @Success 200 {int}
// @router /view [get]
func (ic *ImgaeController) View() {
	var imgPath string
	ic.Ctx.Input.Bind(&imgPath, "path")
	ic.Data["Path"] = imgPath

	ic.Data["URL"] = &models.URL

	u := ic.GetSession("user")
	if u != nil {
		ic.Data["User"] = u.(models.User)
	}

	ic.TplName = "layout/view.html"
}
