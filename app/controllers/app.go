package controllers

import (
    "code.google.com/p/go.net/websocket"
    "github.com/robfig/revel"
    "blog/app/models"
    "os"
)

type App struct {
    *revel.Controller
}

func (c App) CheckToken() bool {
    return c.Session["token"] == os.Getenv("BLOGTOKEN")
}

func (c App) Index() revel.Result {
    posts := models.Post{}.All()
    return c.Render(posts)
}

func (c App) Show(slug string) revel.Result {
    post := models.Post{}.FindBy("slug", slug)
    title := post.Title
    if post == nil {
        c.Response.Status = 404
        return c.NotFound("Doesn't Exist")
    }
    return c.Render(post, title)
}

func (c App) New() revel.Result {
    if !c.CheckToken() {
        return c.Redirect(Session.Destroy)
    }
    return c.Render()
}

func (c App) Create(post models.Post) revel.Result {
    if !c.CheckToken() {
        return c.Redirect(Session.Destroy)
    }

    c.Validation.Required(post.Title).Message("A Title is Required")
    c.Validation.Required(post.Body).Message("A Body is Required")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.New)
    }

    err := post.Create()
    if err != nil {
        c.Flash.Error("Save failed!", err)
        return c.Redirect(App.New)
    }
    c.Flash.Success("Save successful!")
    return c.Redirect(App.Index)
}

func (c App) MarkdownPreview(ws *websocket.Conn) revel.Result {
    for {
        var markdown string
        websocket.Message.Receive(ws, &markdown)
        post := models.Post{Body: markdown}
        post.ParseBody()
        websocket.JSON.Send(ws, &post.Body)
    }
    return nil
}

func (c App) Destroy(id int) revel.Result {
    if !c.CheckToken() {
        return c.Redirect(Session.Destroy)
    }
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
    return c.Redirect(Session.Show)
}
