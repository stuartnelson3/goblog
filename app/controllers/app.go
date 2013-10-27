package controllers

import (
    "github.com/robfig/revel"
    "blog/db"
    "blog/app/models"
)

type App struct {
    *revel.Controller
}

var dbmap = *dbsetup.DbSetup()

func (c App) Index() revel.Result {
    pageHeader := "Main page!!"
    posts := make([]string, 5)
    posts[0] = "This is the first post!!!"
    posts[1] = "This is the second post!!!"
    posts[2] = "This is the third post!!!"
    posts[3] = "This is the fourth post!!!"
    posts[4] = "This is the fifth post!!!"
    return c.Render(posts, pageHeader)
}

func (c App) Show(id int) revel.Result {
    pageHeader := "Show page!!"
    obj, _ := dbmap.Get(models.Post{}, id)
    if obj == nil {
        c.Response.Status = 404
        return c.NotFound("Not found")
    }
    post := obj.(*models.Post)
    return c.Render(post, pageHeader)
}

func (c App) New() revel.Result {
    pageHeader := "New post!!"
    return c.Render(pageHeader)
}
