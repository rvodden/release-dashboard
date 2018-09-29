package route

import (
    "awesomeProject/app/controller"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "strings"
)

func LoadHTTPS() *http.Handler {
    return routes()
}

func LoadHTTP() *http.Handler {
    return routes()
}





func routes() *http.Handler {
    router := mux.NewRouter()

    controller.NewApiController().Register(router,"/api")
    controller.NewClientController().Register(router, "/ui")

    walk := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        pathTemplate, walk := route.GetPathTemplate()
        if walk == nil {
            log.Println("ROUTE:", pathTemplate)
        }
        pathRegexp, walk := route.GetPathRegexp()
        if walk == nil {
            log.Println("Path regexp:", pathRegexp)
        }
        queriesTemplates, walk := route.GetQueriesTemplates()
        if walk == nil {
            log.Println("Queries templates:", strings.Join(queriesTemplates, ","))
        }
        queriesRegexps, walk := route.GetQueriesRegexp()
        if walk == nil {
            log.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
        }
        methods, walk := route.GetMethods()
        if walk == nil {
            log.Println("Methods:", strings.Join(methods, ","))
        }
        log.Println()
        return nil
    })

    if walk != nil {
        log.Println(walk)
    }

    loggedRouter := handlers.LoggingHandler(os.Stdout, router)

    return &loggedRouter
}
