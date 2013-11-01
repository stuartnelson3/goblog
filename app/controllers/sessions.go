package controllers

import (
    "github.com/robfig/revel"
    "blog/app/models"
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
    c.Flash.Success("Successful login!")
    return c.Redirect(App.Index)
}
