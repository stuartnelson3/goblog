package app

import (
    "github.com/robfig/revel"
    "github.com/coopernurse/gorp"
    _ "github.com/bmizerany/pq"
    "database/sql"
    "os"
    "log"
)

type Post struct {
    Id int64
    Title string
    Body string
}

func init() {
    db, err := sql.Open("postgres", "user=stuartnelson dbname=goblog sslmode=disable")

    if err != nil { panic(err) }

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
    dbmap.CreateTablesIfNotExists()

    // Filters is the default set of global filters.
    revel.Filters = []revel.Filter{
        revel.PanicFilter,             // Recover from panics and display an error page instead.
        revel.RouterFilter,            // Use the routing table to select the right Action
        revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
        revel.ParamsFilter,            // Parse parameters into Controller.Params.
        revel.SessionFilter,           // Restore and write the session cookie.
        revel.FlashFilter,             // Restore and write the flash cookie.
        revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
        revel.I18nFilter,              // Resolve the requested language
        revel.InterceptorFilter,       // Run interceptors around the action.
        revel.ActionInvoker,           // Invoke the action.
    }
}
