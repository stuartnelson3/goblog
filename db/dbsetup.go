package dbsetup

import (
    "github.com/coopernurse/gorp"
    "github.com/bmizerany/pq"
    "database/sql"
    "os"
    "log"
)

// func CreatePgSpec(key string, heroku_env_var string) string {
//     return key+"="+os.Getenv(heroku_env_var)+" "
// }

func DbSetup(m map[string]interface{}) (dbmap *gorp.DbMap) {
    // user := CreatePgSpec("user", os.Getenv("HEROKU_POSTGRESQL_USER"))
    // pass := CreatePgSpec("password", os.Getenv("HEROKU_POSTGRESQL_PASS"))
    // dbname := CreatePgSpec("dbname", os.Getenv("HEROKU_POSTGRESQL_DBNAME"))
    // host := CreatePgSpec("host", os.Getenv("HEROKU_POSTGRESQL_HOST"))
    url := os.Getenv("HEROKU_POSTGRESQL_VIOLET_URL")
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require"
    db, _ := sql.Open("postgres", connection) // "prod"
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("query:", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    for k, v := range m {
        dbmap.AddTableWithName(v,k).SetKeys(true, "Id")
    }
    dbmap.CreateTablesIfNotExists()

    return dbmap
}
