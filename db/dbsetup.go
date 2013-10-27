package dbsetup

import (
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "blog/app/models"
    "os"
    "log"
)

func DbSetup() (dbmap *gorp.DbMap) {
    db, _ := sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable")
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("query:", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    dbmap.AddTableWithName(models.Post{}, "posts").SetKeys(true, "Id")
    dbmap.CreateTablesIfNotExists()

    return dbmap
}
