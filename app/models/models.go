package models

import (
    "blog/db"
    "strings"
    "github.com/russross/blackfriday"
)
var tables = map[string]interface{}{"posts":Post{}}
var dbmap = *dbsetup.DbSetup(tables)

type Post struct {
    Id int64
    Title string
    Body string
    Slug string
}

func (p Post) All() []*Post {
    var posts []*Post
    dbmap.Select(&posts, "select * from posts")
    return posts
}

func (p Post) Find(id int) *Post {
    obj, err := dbmap.Get(Post{}, id)
    if err != nil {
        panic(err)
    }
    if obj == nil {
        return nil
    }
    return obj.(*Post)
}

func (p Post) FindBy(field string, cond string) *Post {
    query := "select * from posts where "+field+"="+cond+" limit 1"
    obj, err := dbmap.Select(Post{}, query)
    if err != nil {
        panic(err)
    }
    if obj == nil {
        return nil
    }
    return obj[0].(*Post)
}

func (p Post) Create() error {
    p.ParseBody()
    p.CreateSlug()
    err := dbmap.Insert(&p)
    return err
}

func (p Post) Update() error {
    p.ParseBody()
    _, err := dbmap.Update(&p)
    return err
}

func (p Post) Destroy() error {
    _, err := dbmap.Delete(&p)
    return err
}

func (p *Post) CreateSlug() {
    lower := strings.ToLower(p.Title)
    fields := strings.Fields(lower)
    slug := strings.Join(fields, "-")
    p.Slug = slug
}

func (p *Post) ParseBody() {
    body_as_byte_slice := []byte(p.Body)
    p.Body = string(blackfriday.MarkdownCommon(body_as_byte_slice))
}
