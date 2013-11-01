package dbsetup

import (
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "os"
    "log"
)

func DbSetup(m map[string]interface{}) (dbmap *gorp.DbMap) {
    // db, _ := sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable") // dev
    db, _ := sql.Open("postgres", "user=u6aq2gkmaiir3hrv password=56afebf98a164079989849adc9d589d1 host=bcvpvcjrov5cbmt6.postgresql.clvrcld.net dbname=bcvpvcjrov5cbmt6 sslmode=disable") // "prod"
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("query:", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    for k, v := range m {
        dbmap.AddTableWithName(v,k).SetKeys(true, "Id")
    }
    dbmap.CreateTablesIfNotExists()

    return dbmap
}
