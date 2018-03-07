package controllers

import (
	"encoding/json"

	"github.com/gnenux/pic-bed/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Signup
// @Description signup user
// @router /signup [get]
func (u *UserController) Signup() {
	u.TplName = "layout/signup.html"
	u.Render()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {"uid": uid} or {"error":error}
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(&user)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		u.Data["json"] = map[string]string{"uid": uid}
		u.SetSession("user", user)
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
			u.Abort("403")
			return
		}
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {"uid":uid} or {"error":error}
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = map[string]string{"error": err.Error()}
		}
		u.Data["json"] = map[string]string{"uid": uu.Id}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	err := models.DeleteUser(uid)
	if err != nil {
		u.Abort("403")
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	loginUser		body 	models.LoginUser	true		"The info for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var loginUser models.LoginUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginUser)

	if user, err := models.Login(loginUser.Email, loginUser.Password); err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		u.Data["json"] = map[string]string{"success": "true"}
		u.SetSession("user", *user)
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.DelSession("user")
	u.ServeJSON()
}
