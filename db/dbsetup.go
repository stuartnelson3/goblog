package dbsetup

import (
    "github.com/coopernurse/gorp"
    "github.com/bmizerany/pq"
    "database/sql"
    "os"
    "log"
)

func DbSetup(m map[string]interface{}) (dbmap *gorp.DbMap) {
    url := os.Getenv("DATABASE_URL")
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require" // prod
    // connection += " sslmode=disable" // dev
    db, _ := sql.Open("postgres", connection)

    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("query:", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    for k, v := range m {
        dbmap.AddTableWithName(v,k).SetKeys(true, "Id")
    }
    dbmap.CreateTablesIfNotExists()

    return dbmap
}
