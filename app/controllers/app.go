package controllers

import (
    "github.com/robfig/revel"
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "log"
    "os"
)

var db, err = sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable")
var dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

type App struct {
    *revel.Controller
}

type Post struct {
    Id int64 `db:"id"`
    Title string `db:"title"`
    Body string `db:"body"`
}

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
    dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
    obj, err := dbmap.Get(Post{}, id)
    // post, err := dbmap.Select("select body from posts where id=$1", id)
    dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
    post := obj.(*Post)
    if err != nil {
        c.Response.Status = 404
        return c.NotFound("Not found")
    }
    return c.Render(post, pageHeader)
}

func (c App) New() revel.Result {
    pageHeader := "New post!!"
    return c.Render(pageHeader)
}
