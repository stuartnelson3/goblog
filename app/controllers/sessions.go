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

func (c Session) Create(user models.User) revel.Result {
    err := user.Verify()
    if err != nil {
        c.Flash.Error("Bad login!")
        return c.Redirect(Session.New)
    }
    c.Session["token"] = os.Getenv("BLOGTOKEN")
    c.Flash.Success("Successful login!")
    return c.Redirect(App.Index)
}

func (c Session) Destroy() revel.Result {
    delete(c.Session, "token")
    c.Flash.Success("Logout successful")
    return c.Redirect(App.Index)
}
