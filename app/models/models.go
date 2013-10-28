package models

import "blog/db"
var tables = map[string]interface{}{"posts":Post{}}
var dbmap = *dbsetup.DbSetup(tables)

type Post struct {
    Id int64
    Title string
    Body string
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

func (p Post) Create() error {
    err := dbmap.Insert(&p)
    return err
}

func (p Post) Update() error {
    _, err := dbmap.Update(&p)
    return err
}

func (p Post) Destroy() error {
    _, err := dbmap.Delete(&p)
    return err
}
