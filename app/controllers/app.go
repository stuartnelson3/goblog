package controllers

import (
    "github.com/robfig/revel"
    "blog/app/models"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    pageHeader := "Main page!!"
    posts := models.Post{}.All()
    return c.Render(posts, pageHeader)
}

func (c App) Show(slug string) revel.Result {
    pageHeader := "Show page!!"
    post := models.Post{}.FindBy("slug", slug)
    if post == nil {
        c.Response.Status = 404
        return c.NotFound("Not found")
    }
    return c.Render(post, pageHeader)
}

func (c App) New() revel.Result {
    pageHeader := "New post!!"
    return c.Render(pageHeader)
}

func (c App) Create(post models.Post) revel.Result {
    err := post.Create()
    if err != nil {
        c.Flash.Error("Save failed!", err)
        return c.Redirect(App.New)
    }
    c.Flash.Success("Save successful!")
    return c.Redirect(App.Index)
}

func (c App) Destroy(id int) revel.Result {
    post := models.Post{}.Find(id)
    if post == nil {
        c.Flash.Error("Post not found!")
        return c.Redirect(App.Index)
    }
    err := post.Destroy()
    if err != nil {
        c.Flash.Error("Delete failed!")
        return c.Redirect(App.Index)
    }
    c.Flash.Success("Delete successful!")
    return c.Redirect(App.Index)
}

func (c App) Edit() revel.Result {
    return c.Render()
}

func (c App) Update() revel.Result {
    return c.Render()
}
