package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func init() {

}

type User struct {
	Id       string `orm:"size(25);pk"`
	Username string `json:"name";orm:"size(25)"`
	Email    string `json:"email";orm:"size(25)"`
	Password string `json:"password";orm:"size(25)"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(u *User) (string, error) {
	u.Id = strconv.FormatInt(time.Now().UnixNano(), 10)
	if u.Email == "" {
		return "", errors.New("email is empty")
	}

	uu := &User{Email: u.Email}

	err := O.Read(uu, "Email")
	if err == nil {
		return "", errors.New("Email has been registered")
	}

	_, err = O.Insert(u)
	return u.Id, err
}

func GetUser(uid string) (u *User, err error) {
	u = new(User)
	u.Id = uid
	err = O.Read(u)
	return u, err
}

func GetAllUsers() ([]*User, error) {
	var users []*User
	querySeter := O.QueryTable((*User)(nil))
	_, err := querySeter.All(&users)
	return users, err
}

func UpdateUser(uid string, uu *User) (*User, error) {
	uu.Id = uid
	_, err := O.Update(uu)
	return uu, err
}

func Login(email, password string) (*User, error) {
	u := &User{Email: email}
	fmt.Println("email: " + email)
	err := O.Read(u, "Email")
	if err != nil {
		return nil, errors.New("User not found")
	}

	if u.Password != password {
		return nil, errors.New("password not corrent")
	}
	return u, nil
}

func DeleteUser(uid string) error {
	_, err := O.Delete(&User{Id: uid})
	return err
}
