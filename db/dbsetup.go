package dbsetup

import (
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "os"
    "log"
)

func DbSetup(m map[string]interface{}) (dbmap *gorp.DbMap) {
    db, _ := sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable")
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("query:", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    for k, v := range m {
        dbmap.AddTableWithName(v,k).SetKeys(true, "Id")
    }
    dbmap.CreateTablesIfNotExists()

    return dbmap
}
