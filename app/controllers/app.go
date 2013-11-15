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
    pageHeader := "Main page!!"
    posts := models.Post{}.All()
    return c.Render(posts, pageHeader)
}

func (c App) Show(slug string) revel.Result {
    pageHeader := "Show page!!"
    post := models.Post{}.FindBy("slug", slug)
    if post == nil {
        c.Response.Status = 404
        return c.NotFound("Doesn't Exist")
    }
    return c.Render(post, pageHeader)
}

func (c App) New() revel.Result {
    if !c.CheckToken() {
        return c.Redirect(Session.Destroy)
    }
    pageHeader := "New post!!"
    return c.Render(pageHeader)
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
    newMarkdown := make(chan string)
    go func() {
        var markdown string
        for {
            err := websocket.Message.Receive(ws, &markdown)
            if err != nil {
                close(newMarkdown)
                return
            }
            post := models.Post{Body: markdown}
            post.ParseBody()
            newMarkdown <- post.Body
        }
    }()

    for {
        markdown, ok := <-newMarkdown
        if !ok {
            return nil
        }
        websocket.JSON.Send(ws, &markdown)
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
