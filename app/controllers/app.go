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
    var posts []*models.Post
    query := "select * from posts"
    dbmap.Select(&posts, query)
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
