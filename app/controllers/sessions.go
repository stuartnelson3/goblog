package controllers

import (
    "github.com/robfig/revel"
    "blog/app/models"
    "os"
)

type Session struct {
    App
}

func (c Session) New() revel.Result {
    pageHeader := "Login"
    return c.Render(pageHeader)
}

func (c Session) Show() revel.Result {
    if !c.CheckToken() {
        return c.Redirect(Session.Destroy)
    }
    pageHeader := "Session stuff"
    posts := models.Post{}.All()
    return c.Render(pageHeader, posts)
}

func (c Session) Create(user models.User) revel.Result {
    err := user.Verify()
    if err != nil {
        c.Flash.Error("Bad login!")
        return c.Redirect(Session.New)
    }
    c.Session["token"] = os.Getenv("BLOGTOKEN")
    c.Flash.Success("Successful login!")
    return c.Redirect(Session.Show)
}

func (c Session) Destroy() revel.Result {
    delete(c.Session, "token")
    c.Flash.Success("Logout successful")
    return c.Redirect(App.Index)
}
