package models

import (
    "encoding/json"
    "path/filepath"
    "io/ioutil"
    "time"
    "fmt"
    "strconv"
    "strings"
    "regexp"
    "github.com/russross/blackfriday"
)

type Post struct {
    Title     string `json:"title"`
    Body      string `json:"body"`
    Slug      string `json:"slug"`
    CreatedAt string `json:"createdAt"`
}

func (p Post) All() []*Post {
    var posts []*Post
    matches, _ := filepath.Glob("app/views/Posts/*.json")
    for i:=0; i<len(matches); i++ {
        var post = &Post{}
        data, _ := ioutil.ReadFile(matches[i])
        json.Unmarshal(data, post)
        posts = append(posts, post)
    }
    return posts
}

func (p Post) FindBy(field string, cond string) *Post {
    var post = &Post{}
    absPath, _ := filepath.Abs("app/views/Posts/" + cond + ".json")
    data, _ := ioutil.ReadFile(absPath)
    json.Unmarshal(data, post)
    return post
}

func (p *Post) Create() error {
    p.ParseBody()
    p.CreateSlug()
    p.CreateTimestamp()
    err := p.SaveJson()
    return err
}

func (p *Post) SaveJson() error {
    absPath, _ := filepath.Abs("app/views/Posts/" + p.Slug + ".json")
    err := ioutil.WriteFile(absPath, p.ToJson(), 0644)
    if err != nil { return err }
    return nil
}

func (p *Post) ToJson() []byte {
    serialized, _ := json.Marshal(p)
    return serialized
}

func (p *Post) CreateSlug() {
    lower := strings.ToLower(p.Title)
    fields := strings.Fields(lower)
    slug := strings.Join(fields, "-")
    re := regexp.MustCompile("[^0-9A-Za-z_-]")
    cleaned_slug := re.ReplaceAllLiteralString(slug, "")
    p.Slug = cleaned_slug
}

func (p *Post) ParseBody() {
    body_as_byte_slice := []byte(p.Body)
    p.Body = string(blackfriday.MarkdownCommon(body_as_byte_slice))
}

func (p *Post) CreateTimestamp() {
    createdAt := time.Now()
    month := fmt.Sprint(createdAt.Month())[0:3]
    day := strconv.Itoa(createdAt.Day())
    year := strconv.Itoa(createdAt.Year())
    p.CreatedAt = month + " " + day + " " + year
}

