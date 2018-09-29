package shared

import (
    "github.com/gorilla/sessions"
    "net/http"
)

var (
    Store *sessions.CookieStore
    Name string
)

type Session struct {
    Options   session.Options `json:"Options"`
    Name      string          `json:"Name"`
    SecretKey string          `json:"SecretKey"`
}

func Configure(session Session) {
    Store = sessions.NewCookieStore([]byte(session.SecretKey))
    Store.Options = &session.Options
    Name = session.Name
}

func Instance(request *http.Request) *sessions.Session {
    session, _ := Store.Get(request, Name)
    return session
}
