package dbsetup

import (
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "blog/app/models"
)

func DbSetup() (dbmap *gorp.DbMap) {
    db, _ := sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable")
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

    dbmap.AddTableWithName(models.Post{}, "posts").SetKeys(true, "Id")
    return dbmap
}
